package usecase

import (
	"log"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/mapper"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/factory"
)

type CreatePersonUseCase struct {
	repository domain.PersonRepository
}

func NewCreatePersonUseCase(repo domain.PersonRepository) *CreatePersonUseCase {
	return &CreatePersonUseCase{repository: repo}
}

func (uc *CreatePersonUseCase) Execute(input dto.CreatePersonsInput) error {
	// Mapear DTO a input del dominio
	for _, person := range input.Shareholders {
		personInput := mapper.ToPersonInput(person)

		// Crear persona usando factory
		person, detail, err := factory.CreatePerson(personInput.GetType(), personInput)
		if err != nil {
			return err
		}
		log.Printf("Person: %+v\n", person)
		log.Printf("Detail: %+v\n", detail)
	}

	return nil
}
