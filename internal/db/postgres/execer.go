package postgres

import (
	"net/http"

	persons "github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/interfaces"
)

func LoadRoutesSilicon(router *http.ServeMux) {
	// Load persons routes
	persons.LoadRoutesPersons(router)
}
