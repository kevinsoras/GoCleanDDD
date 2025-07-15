// internal/silicon/shared/interfaces/validator.go
package interfaces

type Validator interface {
	ValidateStruct(interface{}) error
}
