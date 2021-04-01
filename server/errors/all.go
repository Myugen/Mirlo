package errors

import (
	"net/http"

	"github.com/alephshahor/Mirlo/server/enums/languages"
	"github.com/alephshahor/Mirlo/server/types"
)

type APIError struct {
	HTTPStatusCode int
	Message        map[types.LanguageKey]string
}

var (
	ErrEmailAlreadyRegistered = APIError{
		HTTPStatusCode: http.StatusConflict,
		Message: map[types.LanguageKey]string{
			languages.Spanish: "Ya existe una cuenta registrada con este email",
			languages.English: "An account registered with this email already exists",
		},
	}
	ErrUserNameAlreadyRegistered = APIError{
		HTTPStatusCode: http.StatusConflict,
		Message: map[types.LanguageKey]string{
			languages.Spanish: "Ya existe una cuenta registrada con este nombre de usuario",
			languages.English: "An account registered with this user name already exists",
		},
	}
	ErrInvalidLoginCredentials = APIError{
		HTTPStatusCode: http.StatusBadRequest,
		Message: map[types.LanguageKey]string{
			languages.Spanish: "Las credenciales proporcionadas son inválidas",
			languages.English: "The provided credentials are invalid",
		},
	}
	ErrUserNotFound = APIError{
		HTTPStatusCode: http.StatusNotFound,
		Message: map[types.LanguageKey]string{
			languages.Spanish: "El usuario no ha sido encontrado",
			languages.English: "User not found",
		},
	}
)

// Implementing this method so the Error interface is satisfied
func (e APIError) Error() string {
	return e.Message[languages.Spanish]
}
