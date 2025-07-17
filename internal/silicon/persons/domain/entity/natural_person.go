package entity

import (
	"time"
)

type NaturalPerson struct {
	ID               string
	DocumentType     NaturalPersonType
	DocumentNumber   string
	FirstName        string
	LastNamePaternal string
	LastNameMaternal string
	Status           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewNaturalPerson(
	id string,
	docType NaturalPersonType,
	docNumber string,
	firstName string,
	paternalName string,
	maternalName string,
) *NaturalPerson {
	return &NaturalPerson{
		ID:               id,
		DocumentType:     docType,
		DocumentNumber:   docNumber,
		FirstName:        firstName,
		LastNamePaternal: paternalName,
		LastNameMaternal: maternalName,
		Status:           true,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}
func (n NaturalPerson) GetID() string       { return n.ID }
func (n NaturalPerson) GetType() PersonType { return PersonNatural }
func (p *NaturalPerson) Validate() error {
	return nil
}
