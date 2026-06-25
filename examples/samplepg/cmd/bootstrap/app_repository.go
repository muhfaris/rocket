package bootstrap

import (
	"github.com/muhfaris/rocket/examples/samplepg/config"
	psqladapter "github.com/muhfaris/rocket/examples/samplepg/internal/adapter/outbound/datastore/psql"
	pgsqlrepository "github.com/muhfaris/rocket/examples/samplepg/internal/adapter/outbound/datastore/psql/repository"
	portregistry "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/registry"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/port/outbound/repository"
)

func InitializeRepository() portregistry.Repository {
	return &AppRepository{

		psqlRepository: psqladapter.New(psqladapter.PSQLConfig{
			Host:     config.App.Datastore.PSQL.Host,
			Port:     config.App.Datastore.PSQL.Port,
			Username: config.App.Datastore.PSQL.Username,
			Password: config.App.Datastore.PSQL.Password,
			DB:       config.App.Datastore.PSQL.DB,
		}),
	}
}

// AppRepository struct would need to be updated to include new repository types
type AppRepository struct {
	psqlRepository repository.PSQLRepository
}

func (a *AppRepository) PSQLRepository() repository.PSQLRepository {
	return a.psqlRepository
}
func (a *AppRepository) GetReportRepository() repository.ReportRepository {
	return pgsqlrepository.NewReportRepository(a.psqlRepository)
}
