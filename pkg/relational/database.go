package relational

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)

type Database struct {
	context context.Context
	client  *pgxpool.Pool
	mapper  *gorm.DB
}
