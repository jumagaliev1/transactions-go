package service

import (
	"context"
	"transactions/internal/logger"
	"transactions/internal/model"
	"transactions/internal/storage"
)

type AccountService struct {
	repo   *storage.Storage
	logger logger.RequestLogger
}

func NewAccountService(repo *storage.Storage, logger logger.RequestLogger) *AccountService {
	return &AccountService{repo: repo, logger: logger}
}

func (s *AccountService) Create(ctx context.Context, account model.Account) (uint, error) {
	return s.repo.Account.Create(ctx, account)
}

func (s *AccountService) GetByUser(ctx context.Context, userID uint) (*model.Account, error) {
	return s.repo.Account.GetByUser(ctx, userID)
}

func (s *AccountService) Update(ctx context.Context, account model.Account) error {
	return s.repo.Account.Update(ctx, account)
}

func (s *AccountService) Delete(ctx context.Context, accountID uint) error {
	return s.repo.Account.Delete(ctx, accountID)
}
