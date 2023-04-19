package storage

import (
	"context"
	"transactions/internal/config"
	"transactions/internal/logger"
	"transactions/internal/model"
	"transactions/internal/storage/postgre"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction model.Transaction) (uint, error)
	GetByID(ctx context.Context, ID uint) (*model.Transaction, error)
	Delete(ctx context.Context, transactionID uint) error
}

type IAccountRepository interface {
	Create(ctx context.Context, account model.Account) (uint, error)
	GetByUser(ctx context.Context, userID uint) (*model.Account, error)
	Update(ctx context.Context, account model.Account) error
	Delete(ctx context.Context, accountID uint) error
}
type Storage struct {
	Transaction ITransactionRepository
	Account     IAccountRepository
}

func New(ctx context.Context, cfg *config.Config, logger logger.RequestLogger) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	transRepo := postgre.NewTransactionRepository(pgDB, logger)
	accountRepo := postgre.NewAccountRepository(pgDB, logger)

	var storage Storage
	storage.Transaction = transRepo
	storage.Account = accountRepo

	return &storage, nil
}
