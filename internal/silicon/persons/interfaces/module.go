package interfaces

import (
	"net/http"

	"github.com/kevinsoras/GoCleanDDD/internal/db/postgres"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain"
)

type PersonModuleDeps struct {
	UnitOfWork postgres.UnitOfWork
	Repository domain.PersonRepository
}

type PersonModule struct {
	handler *HandlerV1
}

func NewPersonModule(deps PersonModuleDeps) *PersonModule {
	validator := dto.NewCustomValidator()
	handler := NewHandlerV1(validator, deps.Repository)

	return &PersonModule{handler: handler}
}

func (m *PersonModule) RegisterRoutesPersons(router *http.ServeMux) {
	router.HandleFunc("POST /api/v1/persons", m.handler.CreatePerson)
}
