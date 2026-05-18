package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/vaniax17/go-urlShortener/internal/repository/postgres"
	"github.com/vaniax17/go-urlShortener/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load(".env")
	e := echo.New()
	if err != nil {
		panic(err)
	}
	file, err := logger.Init(false)
	if err != nil {
		panic(err)
	}
	l := logger.Get()
	l.Info("logger initialized")

	l.Info("starting server")

	mustEnv(l)

	defer func() {
		err := logger.Get().Sync()
		if err != nil {
			return
		}
		err = file.Close()
		if err != nil {
			return
		}
	}()

	db := postgres.New(buildDSN())

	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))
	err = e.Start(fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")))
	if err != nil {
		panic(err)
	}

}

func mustEnv(l *zap.Logger) {
	if os.Getenv("BACKEND_PORT") == "" {
		l.Error("BACKEND_PORT is not set", zap.Error(errors.New("BACKEND_PORT is not set")))
		panic("BACKEND_PORT must be set")
	}

	if os.Getenv("DB_USER") == "" {
		l.Error("DB_USER is not set", zap.Error(errors.New("DB_USER is not set")))
		panic("DB_USER must be set")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		l.Error("DB_PASSWORD is not set", zap.Error(errors.New("DB_PASSWORD is not set")))
		panic("DB_PASSWORD must be set")
	}

	if os.Getenv("DB_NAME") == "" {
		l.Info("DB_NAME is not set", zap.Error(errors.New("DB_NAME is not set")))
		panic("DB_NAME must be set")
	}

	if os.Getenv("DB_HOST") == "" {
		l.Error("DB_HOST is not set", zap.Error(errors.New("DB_HOST is not set")))
		panic("DB_HOST must be set")
	}

	if os.Getenv("DB_PORT") == "" {
		l.Error("DB_PORT is not set", zap.Error(errors.New("DB_PORT is not set")))
		panic("DB_OPORT must be set")
	}

}

func buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}
