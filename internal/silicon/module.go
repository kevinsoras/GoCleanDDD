package silicon

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinsoras/GoCleanDDD/internal/db/postgres"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/infrastructure"
	persons "github.com/kevinsoras/GoCleanDDD/internal/silicon/persons/interfaces"
)

type SiliconModule struct {
	personModule *persons.PersonModule
}

func NewSiliconModule(db *pgxpool.Pool) *SiliconModule {

	// Create unit of work
	uow := postgres.NewPgUnitOfWork(db)
	// Implementación concreta del repositorio
	repoPerson := infrastructure.NewPersonRepository(db)
	// Crear dependencias propias del contexto Silicon
	personModule := persons.NewPersonModule(persons.PersonModuleDeps{
		UnitOfWork: uow,
		Repository: repoPerson,
	})

	return &SiliconModule{
		personModule: personModule,
	}
}

func (s *SiliconModule) RegisterRoutes(router *http.ServeMux) {
	s.personModule.RegisterRoutesPersons(router)
}
