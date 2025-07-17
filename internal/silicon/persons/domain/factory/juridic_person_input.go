package factory

import (
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/entity"
)

type JuridicPersonInput struct {
	ID             *string
	RUC            string
	LegalName      string
	CommercialName string
	Address        string
	District       string
	Province       string
	Department     string
	Country        string
	HasConstituted bool
}

func createJuridicPerson(id string, input PersonInput) (entity.PersonDetail, error) {
	j := input.(JuridicPersonInput)

	return entity.NewCreateJuridicPerson(
		id,
		j.RUC,
		j.LegalName,
		j.CommercialName,
		j.Address,
		j.District,
		j.Province,
		j.Department,
		j.Country,
		j.HasConstituted,
	), nil
}
func (j JuridicPersonInput) GetID() string              { return *j.ID }
func (j JuridicPersonInput) GetType() entity.PersonType { return entity.PersonJuridica }
