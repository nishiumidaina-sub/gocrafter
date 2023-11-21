// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	account2 "github.com/game-core/gocrafter/api/presentation/controller/account"
	"github.com/game-core/gocrafter/api/presentation/middleware/account"
	"github.com/game-core/gocrafter/config/database"
	account4 "github.com/game-core/gocrafter/domain/service/account"
	"github.com/game-core/gocrafter/infra/dao/user"
	account3 "github.com/game-core/gocrafter/infra/dao/user/account"
)

// Injectors from wire.go:

func InitializeAccountMiddleware() account.AccountMiddleware {
	accountMiddleware := account.NewAccountMiddleware()
	return accountMiddleware
}

func InitializeAccountController() account2.AccountController {
	sqlHandler := database.NewDB()
	transactionRepository := user.NewTransactionDao(sqlHandler)
	accountRepository := account3.NewAccountDao(sqlHandler)
	accountService := account4.NewAccountService(transactionRepository, accountRepository)
	accountController := account2.NewAccountController(accountService)
	return accountController
}
