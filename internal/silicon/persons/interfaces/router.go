package persons

import (
	"net/http"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
)

// LoadRoutesPersons carga las rutas y crea sus propias dependencias
func LoadRoutesPersons(router *http.ServeMux) {
	// Crea el validador aquí (no necesita exponerse a main)
	validator := dto.NewCustomValidator()
	handlerV1 := NewHandlerV1(validator) // Inyecta el validador creado localmente

	router.HandleFunc("POST /api/v1/persons", handlerV1.CreatePerson)
}
