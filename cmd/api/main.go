package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load(".env")
	e := echo.New()
	if err != nil {
		panic(err)
	}

	mustEnv()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))
	err = e.Start(fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")))
	if err != nil {
		panic(err)
	}

}

func mustEnv() {
	if os.Getenv("BACKEND_PORT") == "" {
		panic("BACKEND_PORT must be set")
	}
}
