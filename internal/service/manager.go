package service

import (
	"context"
	"transactions/internal/model"
)

type ITransactionService interface {
	Create(ctx context.Context, transaction model.Transaction) (uint, error)
	Cancel(ctx context.Context, transactionID uint) error
}

type IAccountService interface {
	Create(ctx context.Context, account model.Account) (uint, error)
	GetByUser(ctx context.Context, userID uint) (*model.Account, error)
	Update(ctx context.Context, account model.Account) error
	Delete(ctx context.Context, accountID uint) error
}
