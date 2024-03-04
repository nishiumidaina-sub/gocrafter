package router

import (
	"flag"
	"log"

	"github.com/game-core/gocrafter/pkg/logger"
)

func Router() {
	// Batch list
	batchList := map[string]func() error{}

	commandPtr := flag.String("command", "", "Specify the command")
	flag.Parse()
	command := *commandPtr
	selectedCommand := command

	// Run
	if batch, ok := batchList[selectedCommand]; ok {
		logFile := logger.DebugLog("./pkg/logger/batch/game/", "admin_debug.")
		log.SetOutput(logFile)
		if err := batch(); err != nil {
			log.Printf("failed to batch(): %s", err)
		}
	}
}
