package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/utils"
)

type CreatePersonsInput struct {
	Shareholders []CreatePersonInput `json:"shareholders" validate:"required,dive"`
}

type CreatePersonInput struct {
	Type string `json:"type" validate:"required,oneof=NATURAL JURIDICA"`

	// NATURAL
	TypeDocument     string `json:"typeDocument" validate:"omitempty,min=2,max=20"` // Considera oneof si hay valores fijos
	DocumentNumber   string `json:"documentNumber" validate:"omitempty,min=8,max=15"`
	FirstName        string `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastNamePaternal string `json:"lastNamePaternal" validate:"omitempty,min=2,max=50"`
	LastNameMaternal string `json:"lastNameMaternal" validate:"omitempty,min=2,max=50"`

	// JURIDICA
	RUC            *string `json:"ruc" validate:"omitempty,len=11"`
	LegalName      string  `json:"legalName" validate:"omitempty,min=3,max=100"`
	CommercialName *string `json:"commercialName" validate:"omitempty,min=3,max=100"`
	Address        string  `json:"address" validate:"omitempty,min=3,max=150"`
	District       string  `json:"district" validate:"omitempty,min=3,max=50"`
	Province       string  `json:"province" validate:"omitempty,min=3,max=50"`
	Department     string  `json:"department" validate:"omitempty,min=3,max=50"`
	Country        string  `json:"country" validate:"omitempty,min=3,max=50"`
	HasConstituted *bool   `json:"hasConstituted" validate:"omitempty"`
}

func CreatePersonInputStructLevelValidation(sl validator.StructLevel) {
	p := sl.Current().Interface().(CreatePersonInput)

	switch p.Type {
	case "NATURAL":
		utils.ReportIfZeroValue(sl, p.TypeDocument, "typeDocument", "TypeDocument", "required")
		utils.ReportIfZeroValue(sl, p.DocumentNumber, "documentNumber", "DocumentNumber", "required")
		utils.ReportIfZeroValue(sl, p.FirstName, "firstName", "FirstName", "required")
		utils.ReportIfZeroValue(sl, p.LastNamePaternal, "lastNamePaternal", "LastNamePaternal", "required")
		utils.ReportIfZeroValue(sl, p.LastNameMaternal, "lastNameMaternal", "LastNameMaternal", "required")

	case "JURIDICA":
		utils.ReportIfZeroValue(sl, p.LegalName, "legalName", "LegalName", "required")
	}
}
