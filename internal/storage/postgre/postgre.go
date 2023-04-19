package postgre

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"transactions/internal/config"
	"transactions/internal/model"
)

func Dial(ctx context.Context, cfg config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.URI()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Transaction{}, &model.Account{})
	return db, nil
}
