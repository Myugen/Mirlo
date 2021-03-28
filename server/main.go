package main

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/handlers"
	"github.com/alephshahor/Mirlo/server/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	var err error
	var e *echo.Echo

	if err = loadEnv(); err != nil {
		panic(err)
	}

	e = setupEcho()

	handlers.InitializeHandlers(e, services.Services())

	e.Logger.Fatal(e.Start(":8080"))
}

func loadEnv() error {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		return fmt.Errorf("Fatal error loading .env file: %s \n", err)
	}
	return nil
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(ErrorHandler)
	return e
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
