package models

// ErrorResponse representa la estructura de respuesta para errores de la API
type ErrorResponse struct {
	Error string `json:"error" example:"mensaje descriptivo del error"`
}

// SuccessResponse representa una respuesta exitosa con mensaje
type SuccessResponse struct {
	Message string `json:"message" example:"operación realizada con éxito"`
}
