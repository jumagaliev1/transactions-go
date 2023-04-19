package service

import (
	"context"
	"errors"
	"transactions/internal/logger"
	"transactions/internal/model"
	"transactions/internal/storage"
)

type TransactionService struct {
	repo    *storage.Storage
	logger  logger.RequestLogger
	account IAccountService
}

func NewTransactionService(repo *storage.Storage, logger logger.RequestLogger, accountService IAccountService) *TransactionService {
	return &TransactionService{repo: repo, logger: logger, account: accountService}
}

func (s *TransactionService) Create(ctx context.Context, transaction model.Transaction) (uint, error) {
	account, err := s.account.GetByUser(ctx, transaction.UserID)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return 0, err
	}

	if !checkBalance(account.Balance, transaction.Price) {
		s.logger.Logger(ctx).Error("not enough money")
		return 0, errors.New("not enough money")
	}

	account.Balance -= transaction.Price

	err = s.account.Update(ctx, *account)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return 0, err
	}

	return s.repo.Transaction.Create(ctx, transaction)
}

func (s *TransactionService) Cancel(ctx context.Context, transactionID uint) error {
	transaction, err := s.repo.Transaction.GetByID(ctx, transactionID)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	account, err := s.account.GetByUser(ctx, transaction.UserID)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	account.Balance += transaction.Price

	err = s.account.Update(ctx, *account)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	return s.repo.Transaction.Delete(ctx, transactionID)
}

func checkBalance(userBalance, price int) bool {
	if userBalance >= price {
		return true
	}

	return false
}
