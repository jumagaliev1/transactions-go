package postgre

import (
	"context"
	"gorm.io/gorm"
	"transactions/internal/logger"
	"transactions/internal/model"
)

type TransactionRepository struct {
	DB     *gorm.DB
	logger logger.RequestLogger
}

func NewTransactionRepository(DB *gorm.DB, logger logger.RequestLogger) *TransactionRepository {
	return &TransactionRepository{DB: DB, logger: logger}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction model.Transaction) (uint, error) {
	if err := r.DB.WithContext(ctx).Create(&transaction).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return 0, err
	}

	return transaction.ID, nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, ID uint) (*model.Transaction, error) {
	var transaction model.Transaction

	if err := r.DB.WithContext(ctx).Find(&transaction, ID).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Delete(ctx context.Context, transactionID uint) error {
	if err := r.DB.WithContext(ctx).Delete(&model.Transaction{}, transactionID).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return err
	}

	return nil
}
