package domain

import (
	"context"

	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/entity"
)

type PersonRepository interface {
	Save(ctx context.Context, person *entity.Person) error
	FindByID(ctx context.Context, id string) (*entity.Person, error)
}
