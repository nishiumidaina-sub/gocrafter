//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"errors"
	"github.com/game-core/gocrafter/config/token"
	"log"

	"github.com/game-core/gocrafter/config/key"

	request "github.com/game-core/gocrafter/auth/presentation/request/account"
	response "github.com/game-core/gocrafter/auth/presentation/response/account"
	accountEntity "github.com/game-core/gocrafter/domain/entity/auth/account"
	authRepository "github.com/game-core/gocrafter/domain/repository/auth"
	accountRepository "github.com/game-core/gocrafter/domain/repository/auth/account"
)

type AccountService interface {
	RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error)
	LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error)
}

type accountService struct {
	transactionRepository authRepository.TransactionRepository
	accountRepository     accountRepository.AccountRepository
}

func NewAccountService(
	transactionRepository authRepository.TransactionRepository,
	accountRepository accountRepository.AccountRepository,
) AccountService {
	return &accountService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

// RegisterAccount アカウントを登録する
func (s *accountService) RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error) {
	// transaction
	tx, err := s.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := s.transactionRepository.CommitOrRollback(tx, err); err != nil {
			log.Panicln(err)
		}
	}()

	hashedPassword, err := key.GenerateHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	account := &accountEntity.Account{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	ar, err := s.accountRepository.Create(account, tx)
	if err != nil {
		return nil, err
	}

	return response.ToRegisterAccount(200, *response.ToAccount(ar.ID, ar.Name, ar.Email, req.Password, "")), nil
}

// LoginAccount アカウントをログインする
func (s *accountService) LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error) {
	ar, err := s.accountRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !key.CheckPassword(req.Password, ar.Password) {
		return nil, errors.New("faild to key.CheckPassword")
	}

	token, err := token.GenerateAuthTokenByEmail(ar.Email, ar.Name)
	if err != nil {
		return nil, errors.New("faild to token.GenerateAuthToken")
	}

	return response.ToLoginAccount(200, *response.ToAccount(ar.ID, ar.Name, ar.Email, req.Password, token)), nil
}