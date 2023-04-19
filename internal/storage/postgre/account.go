package postgre

import (
	"context"
	"gorm.io/gorm"
	"transactions/internal/logger"
	"transactions/internal/model"
)

type AccountRepository struct {
	DB     *gorm.DB
	logger logger.RequestLogger
}

func NewAccountRepository(DB *gorm.DB, logger logger.RequestLogger) *AccountRepository {
	return &AccountRepository{DB: DB, logger: logger}
}

func (r *AccountRepository) Create(ctx context.Context, account model.Account) (uint, error) {
	if err := r.DB.WithContext(ctx).Create(&account).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return 0, err
	}

	return account.ID, nil
}

func (r *AccountRepository) GetByUser(ctx context.Context, userID uint) (*model.Account, error) {
	var account model.Account

	if err := r.DB.WithContext(ctx).Find(&account, "user_id = ?", userID).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) Update(ctx context.Context, account model.Account) error {
	if err := r.DB.WithContext(ctx).Save(account).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return err
	}

	return nil
}

func (r *AccountRepository) Delete(ctx context.Context, accountID uint) error {
	if err := r.DB.WithContext(ctx).Delete(model.Account{}, accountID).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return err
	}

	return nil
}
