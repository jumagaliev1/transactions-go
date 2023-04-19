package service

import (
	"errors"
	"transactions/internal/config"
	"transactions/internal/logger"
	"transactions/internal/storage"
)

type Service struct {
	Transaction ITransactionService
	Account     IAccountService
}

func New(repo *storage.Storage, cfg config.Config, logger logger.RequestLogger) (*Service, error) {
	if repo == nil {
		return nil, errors.New("No storage")
	}
	accountService := NewAccountService(repo, logger)
	transService := NewTransactionService(repo, logger, accountService)
	return &Service{
		Transaction: transService,
		Account:     accountService,
	}, nil
}
