package factory

import (
	"errors"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/entity"
)

// Interface común para todos los tipos de persona
type PersonInput interface {
	GetID() string
	GetType() entity.PersonType
}

// Tipo para funciones creadoras
type createFunc func(id string, input PersonInput) (entity.PersonDetail, error)

var creators = map[entity.PersonType]createFunc{
	entity.PersonNatural:  createNaturalPerson,
	entity.PersonJuridica: createJuridicPerson,
}

func CreatePerson(personType entity.PersonType, input PersonInput) (*entity.Person, entity.PersonDetail, error) {
	// Crear la entidad raíz Person
	person := entity.NewCreatePerson(personType)

	// Buscar el factory específico (Natural o Jurídica)
	fn, ok := creators[personType]
	if !ok {
		return nil, nil, errors.New("person type not supported")
	}

	detail, err := fn(person.ID, input)
	if err != nil {
		return nil, nil, err
	}

	return person, detail, nil
}
