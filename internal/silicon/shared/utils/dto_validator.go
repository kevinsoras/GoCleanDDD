package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/interfaces"
)

// ApiResponse estructura estándar para respuestas

// WriteResponse escribe respuestas HTTP estandarizadas
func WriteResponse(w http.ResponseWriter, status int, message string, data interface{}, errors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(interfaces.ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}

// WriteValidationErrors maneja errores de validación con soporte para arrays
func WriteValidationErrors(w http.ResponseWriter, err error) {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErrs {
			fieldPath := getFieldPath(e.StructNamespace(), e.Field())
			errors[fieldPath] = getValidationMessage(e.Tag(), e.Param())
		}

		WriteResponse(
			w,
			http.StatusBadRequest,
			"Validation failed",
			nil,
			errors,
		)
		return
	}
	WriteResponse(w, http.StatusBadRequest, "Invalid request data", nil, nil)
}

// ProcessRequest procesa y valida requests automáticamente
func ProcessRequest[T any](w http.ResponseWriter, r *http.Request, validator interfaces.Validator) (*T, bool) {
	var input T

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid JSON format", nil, nil)
		return nil, false
	}

	if validator != nil {
		if err := validator.ValidateStruct(input); err != nil {
			WriteValidationErrors(w, err)
			return nil, false
		}
	}

	return &input, true
}

// Helpers internos --------------------------------------------------

// getFieldPath convierte "CreatePersonsInput.Shareholders[1].RUC" → "shareholders[1].ruc"
func getFieldPath(namespace, field string) string {
	parts := strings.Split(namespace, ".")
	var cleanParts []string

	for i, part := range parts {
		if i == 0 {
			continue // Omite el nombre del struct raíz
		}

		// Maneja arrays (ej: Shareholders[0] → shareholders[0])
		if idx := strings.Index(part, "["); idx > 0 {
			fieldName := strings.ToLower(part[:idx])
			cleanParts = append(cleanParts, fieldName+part[idx:])
		} else {
			cleanParts = append(cleanParts, strings.ToLower(part))
		}
	}

	return strings.Join(cleanParts, ".")
}

// getValidationMessage devuelve mensajes legibles
func getValidationMessage(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "len":
		return "Must be exactly " + param + " characters"
	case "min":
		return "Minimum length is " + param
	case "max":
		return "Maximum length is " + param
	case "oneof":
		return "Must be one of: " + strings.Replace(param, " ", ", ", -1)
	default:
		return "Validation failed on '" + tag + "'"
	}
}
