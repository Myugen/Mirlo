package handlers

import (
	"net/http"

	"github.com/alephshahor/Mirlo/server/mappers"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/responses"
	"github.com/alephshahor/Mirlo/server/services"
	"github.com/labstack/echo"
)

type IUserHandlerServices interface {
	SignUp() services.ISignUpService
}

type userHandler struct {
	services IUserHandlerServices
}

func InitializeUserHandler(e *echo.Echo, services services.IServices) *userHandler {
	var h = &userHandler{
		services: services,
	}
	h.registerRoutes(e)
	return h
}

func (h *userHandler) registerRoutes(e *echo.Echo) {
	e.POST("/users", h.Register)
}

func (h *userHandler) Register(c echo.Context) error {
	var err error

	var newUserReq requests.NewUser
	if err = c.Bind(&newUserReq); err != nil {
		return err
	}

	var newUser models.User
	if newUser, err = h.services.SignUp().Register(newUserReq); err != nil {
		return err
	}

	var newUserRes responses.NewUser
	newUserRes = mappers.UserModelToNewUserResponse(newUser)

	return c.JSON(http.StatusCreated, newUserRes)
}
