package postgres

import (
	"github.com/vaniax17/go-urlShortener/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func New(dsn string) *DB {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	l := logger.Get()
	l.Info("database connection established")

	return &DB{DB: db}
}

func (d *DB) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
