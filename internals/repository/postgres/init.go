package postgres

import "github.com/BigNutJaa/user-service/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}
