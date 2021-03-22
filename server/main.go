package main

import (
	"github.com/alephshahor/Mirlo/server/handlers"
	"github.com/alephshahor/Mirlo/server/services"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	handlers.InitializeHandlers(e, services.Services())
	e.Logger.Fatal(e.Start(":8080"))
}
