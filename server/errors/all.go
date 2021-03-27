package errors

import (
	"github.com/alephshahor/Mirlo/server/enums/languages"
	"github.com/alephshahor/Mirlo/server/types"
	"net/http"
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
)

// Implementing this method so the Error interface is satisfied
func (e APIError) Error() string {
	return e.Message[languages.Spanish]
}
