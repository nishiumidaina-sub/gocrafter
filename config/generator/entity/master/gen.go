//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

type StructField struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
}

type StructInfo struct {
	Name    string                 `yaml:"name"`
	Package string                 `yaml:"package"`
	Fields  map[string]StructField `yaml:"structure"`
	Primary []string               `yaml:"primary"`
	Index   []string               `yaml:"index"`
}

const templateCode = `
package {{.Package}}

import (
	"time"
)

type {{.Name}}s []{{.Name}}

type {{.Name}} struct {
{{range $field := sortByNumber .Fields}}
	{{$field.Name}} {{$field.TypeWithPointer}} ` + "`json:\"{{$field.Json}}\"{{if eq $field.Name \"CreatedAt\"}} gorm:\"autoCreateTime\"{{else if eq $field.Name \"UpdatedAt\"}} gorm:\"autoUpdateTime\"{{end}}`" + `
{{end}}
}
`

func generateEntity(yamlFile string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFile)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_entiry.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	if err := generateTemplate(structInfo, outputFile); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Entity in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateTemplate(structInfo *StructInfo, outputFile *os.File) error {
	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sortByNumber": sortByNumber,
	}).Parse(templateCode)
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Name    string
		Package string
		Fields  map[string]StructField
	}{
		Name:    structInfo.Name,
		Package: structInfo.Package,
		Fields:  structInfo.Fields,
	}); err != nil {
		return err
	}

	return nil
}

func sortByNumber(fields map[string]StructField) []struct {
	Name            string
	FieldInfo       StructField
	TypeWithPointer string
	Json            string
} {
	var sortedFields []struct {
		Name            string
		FieldInfo       StructField
		TypeWithPointer string
		Json            string
	}

	for name, fieldInfo := range fields {
		sortedFields = append(sortedFields, struct {
			Name            string
			FieldInfo       StructField
			TypeWithPointer string
			Json            string
		}{
			Name:            name,
			FieldInfo:       fieldInfo,
			TypeWithPointer: getTypeWithPointer(fieldInfo),
			Json:            fieldInfo.Name,
		})
	}

	sort.SliceStable(sortedFields, func(i, j int) bool {
		return fields[sortedFields[i].Name].Number < fields[sortedFields[j].Name].Number
	})

	return sortedFields
}

func getStructInfo(yamlFilePath string) (*StructInfo, error) {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	if err := yaml.Unmarshal(yamlData, &structInfo); err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	return &structInfo, nil
}

func getTypeWithPointer(fieldInfo StructField) string {
	if fieldInfo.Nullable {
		return "*" + fieldInfo.Type
	}

	return fieldInfo.Type
}

func main() {
	outputBaseDir := "../../../../domain/entity/master"
	yamlFiles, err := filepath.Glob("../../../../docs/entity/master/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateEntity(yamlFile, outputBaseDir)
		if err != nil {
			log.Printf("Error generating entity from YAML file %s: %v", yamlFile, err)
		}
	}
}
