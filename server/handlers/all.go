package handlers

import (
	"github.com/alephshahor/Mirlo/server/services"
	"github.com/labstack/echo"
)

type Handlers struct {
	userHandler *userHandler
}

func InitializeHandlers(e *echo.Echo, services services.IServices) *Handlers {
	return &Handlers{
		userHandler: InitializeUserHandler(e, services),
	}
}
