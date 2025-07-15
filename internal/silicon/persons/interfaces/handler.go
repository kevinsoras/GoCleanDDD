package persons

import (
	"net/http"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/interfaces"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/utils"
)

type HandlerV1 struct {
	validator interfaces.Validator
}

func NewHandlerV1(v interfaces.Validator) *HandlerV1 {
	return &HandlerV1{validator: v}
}

func (h *HandlerV1) CreatePerson(w http.ResponseWriter, r *http.Request) {
	input, ok := utils.ProcessRequest[dto.CreatePersonInput](w, r, h.validator)
	if !ok {
		return
	}

	utils.WriteResponse(
		w,
		http.StatusCreated,
		"Person created successfully",
		input,
		nil,
	)
}
