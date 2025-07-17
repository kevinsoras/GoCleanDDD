package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/domain/entity"
)

type PersonPg struct {
	db *pgxpool.Pool
}

func NewPersonRepository(db *pgxpool.Pool) domain.PersonRepository {
	return &PersonPg{db: db}
}

func (r *PersonPg) Save(ctx context.Context, person *entity.Person) error {
	return nil
}

func (r *PersonPg) FindByID(ctx context.Context, id string) (*entity.Person, error) {
	return nil, nil
}
