package apierror

import (
	"errors"
	"net/http"
)

// Type es el tipo de error
type Type string

// Definición de tipos de errores
const (
	NotFound       Type = "NOT_FOUND"
	ValidationFail Type = "VALIDATION_FAILED"
	ResourceExists Type = "RESOURCE_EXISTS"
	BadRequest     Type = "BAD_REQUEST"
	Internal       Type = "INTERNAL_ERROR"
	Unauthorized   Type = "UNAUTHORIZED"
	Forbidden      Type = "FORBIDDEN"
)

// Error es la estructura para errores personalizados
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error devuelve el mensaje de error
func (e Error) Error() string {
	return e.Message
}

// Status devuelve el código HTTP apropiado según el tipo de error
func (e Error) Status() int {
	switch e.Type {
	case NotFound:
		return http.StatusNotFound
	case ValidationFail, BadRequest:
		return http.StatusBadRequest
	case ResourceExists:
		return http.StatusConflict
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// NewError crea un nuevo error personalizado
func NewError(errType Type, message string) Error {
	return Error{
		Type:    errType,
		Message: message,
	}
}

// AsError convierte un error a Error personalizado si es posible
func AsError(err error) (Error, bool) {
	var apiErr Error
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return Error{}, false
}
