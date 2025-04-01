package models

// ErrorResponse representa la estructura de respuesta para errores de la API en la documentaicón
type ErrorResponse struct {
	Error string `json:"error" example:"mensaje descriptivo del error"`
}

// SuccessResponse representa una respuesta exitosa con mensaje  en la documentaicón
type SuccessResponse struct {
	Message string `json:"message" example:"operación realizada con éxito"`
}
