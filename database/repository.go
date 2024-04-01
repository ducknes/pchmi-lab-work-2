package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"lab-work-2/config"
)

type LogisticRepository struct {
	db *sqlx.DB
}

func NewLogisticRepository(ctx context.Context, cfg config.Config) (*LogisticRepository, error) {
	db, err := sqlx.ConnectContext(ctx, cfg.DatabaseDriver, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных. ошибка: %w", err)
	}

	return &LogisticRepository{
		db: db,
	}, nil
}
