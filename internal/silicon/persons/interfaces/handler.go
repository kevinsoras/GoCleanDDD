package interfaces

import (
	"net/http"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/dto"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/application/usecase"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/interfaces"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/utils"
)

type HandlerV1 struct {
	validator  interfaces.Validator
	repository domain.PersonRepository
}

func NewHandlerV1(v interfaces.Validator, r domain.PersonRepository) *HandlerV1 {
	return &HandlerV1{validator: v, repository: r}
}

func (h *HandlerV1) CreatePerson(w http.ResponseWriter, r *http.Request) {
	input, ok := utils.ProcessRequest[dto.CreatePersonsInput](w, r, h.validator)
	if !ok {
		return
	}
	useCase := usecase.NewCreatePersonUseCase(h.repository)

	err := useCase.Execute(*input)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, "Error creating person", nil, nil)
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
