package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinsoras/GoCleanDDD/internal/db/postgres"
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
	_, err := postgres.Exec(ctx, r.db,
		"INSERT INTO persons (id, first_name, last_name, document, type) VALUES ($1, $2, $3, $4, $5)",
		person.ID, person.FirstName, person.LastName, person.Document, person.Type,
	)
	return err
}

func (r *PersonPg) FindByID(ctx context.Context, id string) (*entity.Person, error) {
	query := `
		SELECT id, first_name, last_name, document, type
		FROM persons
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var person entity.Person
	err := row.Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.Document,
		&person.Type,
	)

	if err != nil {
		return nil, fmt.Errorf("person not found: %w", err)
	}

	return &person, nil
}
