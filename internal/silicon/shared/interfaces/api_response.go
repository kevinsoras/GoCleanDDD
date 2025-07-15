package interfaces

// ApiResponse define la estructura estándar de todas las respuestas HTTP
type ApiResponse struct {
	Status  int               `json:"status"`           // Código HTTP
	Message string            `json:"message"`          // Mensaje descriptivo
	Data    interface{}       `json:"data,omitempty"`   // Datos de respuesta (opcional)
	Errors  map[string]string `json:"errors,omitempty"` // Errores de validación (opcional)
}
