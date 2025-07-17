package utils

import (
	"github.com/samborkent/uuidv7"
)

// UUIDGenerator interface para permitir mocking en tests
type UUIDGenerator interface {
	Generate() string
}

type uuidV7Generator struct{}

// NewUUIDGenerator crea una nueva instancia del generador de UUID v7
func NewUUIDGenerator() UUIDGenerator {
	return &uuidV7Generator{}
}

// Generate genera un nuevo UUID v7 como string
func (g *uuidV7Generator) Generate() string {
	return uuidv7.New().String()
}

// Package-level helper functions
var generator UUIDGenerator = NewUUIDGenerator()

// GenerateUUID genera un UUID v7 usando el generador por defecto
func GenerateUUID() string {
	return generator.Generate()
}

// SetGenerator permite inyectar un generador alternativo (útil para testing)
func SetGenerator(g UUIDGenerator) {
	generator = g
}
