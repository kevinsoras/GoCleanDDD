package mapper

import (
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/factory"
)

func ToPersonInput(input dto.CreatePersonInput) factory.PersonInput {
	switch input.Type {
	case "NATURAL":
		return factory.NaturalPersonInput{
			ID:               input.ID,
			DocumentType:     input.TypeDocument,
			DocumentNumber:   input.DocumentNumber,
			FirstName:        input.FirstName,
			LastNamePaternal: input.LastNamePaternal,
			LastNameMaternal: input.LastNameMaternal,
		}
	case "JURIDICA":
		return factory.JuridicPersonInput{
			ID:             input.ID,
			RUC:            *input.RUC,
			LegalName:      input.LegalName,
			CommercialName: *input.CommercialName,
			Address:        input.Address,
			District:       input.District,
			Province:       input.Province,
			Department:     input.Department,
			Country:        input.Country,
			HasConstituted: *input.HasConstituted,
		}
	default:
		return nil
	}
}
