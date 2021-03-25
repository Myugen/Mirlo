package main

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/handlers"
	"github.com/alephshahor/Mirlo/server/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"os"

	"github.com/labstack/echo"
)

func main() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	handlers.InitializeHandlers(e, services.Services())
	e.Logger.Fatal(e.Start(":8080"))
}
