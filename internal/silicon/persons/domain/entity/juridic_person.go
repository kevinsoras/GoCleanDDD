package entity

import (
	"errors"
	"time"
)

type JuridicPerson struct {
	ID             string
	RUC            string
	LegalName      string
	CommercialName string
	Address        string
	District       string
	Province       string
	Department     string
	Country        string
	HasConstituted bool
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewCreateJuridicPerson(
	id string,
	ruc string,
	legalName string,
	commercialName string,
	address string,
	district string,
	province string,
	department string,
	country string,
	hasConstituted bool,
) *JuridicPerson {
	now := time.Now()

	return &JuridicPerson{
		ID:             id,
		RUC:            ruc,
		LegalName:      legalName,
		CommercialName: commercialName,
		Address:        address,
		District:       district,
		Province:       province,
		Department:     department,
		Country:        country,
		HasConstituted: hasConstituted,
		Status:         true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func (n JuridicPerson) GetID() string       { return n.ID }
func (n JuridicPerson) GetType() PersonType { return PersonJuridica }

func (p *JuridicPerson) Validate() error {
	switch {
	case len(p.RUC) != 11:
		return errors.New("RUC debe tener 11 dígitos")
	case p.LegalName == "":
		return errors.New("razón social es requerida")
	}
	return nil
}
