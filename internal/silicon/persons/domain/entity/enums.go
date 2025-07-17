package entity

type PersonType string

const (
	PersonNatural  PersonType = "NATURAL"
	PersonJuridica PersonType = "JURIDICA"
)

type NaturalPersonType string

const (
	DNI       NaturalPersonType = "DNI"
	Pasaporte NaturalPersonType = "PASAPORTE"
	CarnetExt NaturalPersonType = "CARNET DE EXTRANJERIA"
)
