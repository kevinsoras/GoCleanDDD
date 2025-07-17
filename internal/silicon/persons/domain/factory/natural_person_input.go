package factory

import (
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/entity"
)

type NaturalPersonInput struct {
	ID               *string
	DocumentType     string
	DocumentNumber   string
	FirstName        string
	LastNamePaternal string
	LastNameMaternal string
}

// Implementación para NaturalPerson
func createNaturalPerson(id string, input PersonInput) (entity.PersonDetail, error) {
	n := input.(NaturalPersonInput)

	return entity.NewNaturalPerson(
		id,
		entity.NaturalPersonType(n.DocumentType),
		n.DocumentNumber,
		n.FirstName,
		n.LastNamePaternal,
		n.LastNameMaternal,
	), nil
}

func (n NaturalPersonInput) GetID() string              { return *n.ID }
func (n NaturalPersonInput) GetType() entity.PersonType { return entity.PersonNatural }
