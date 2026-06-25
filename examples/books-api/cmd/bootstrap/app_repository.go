package bootstrap

import (
	"github.com/muhfaris/rocket/examples/books-api/config"
	redisadapter "github.com/muhfaris/rocket/examples/books-api/internal/adapter/outbound/cache/redis"
	psqladapter "github.com/muhfaris/rocket/examples/books-api/internal/adapter/outbound/datastore/psql"
	pgsqlrepository "github.com/muhfaris/rocket/examples/books-api/internal/adapter/outbound/datastore/psql/repository"
	portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/port/outbound/repository"
)

func InitializeRepository() portregistry.Repository {
	return &AppRepository{

		cacheRepository: redisadapter.New(
			redisadapter.RedisOptions{
				Addr:     config.App.Cache.Redis.Addr,
				Username: config.App.Cache.Redis.Username,
				Password: config.App.Cache.Redis.Password,
				DB:       config.App.Cache.Redis.DB,
			},
		),

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
	cacheRepository repository.CacheRepository
	psqlRepository  repository.PSQLRepository
}

func (a *AppRepository) CacheRepository() repository.CacheRepository {
	return a.cacheRepository
}

func (a *AppRepository) PSQLRepository() repository.PSQLRepository {
	return a.psqlRepository
}
func (a *AppRepository) GetBookRepository() repository.BookRepository {
	return pgsqlrepository.NewBookRepository(a.psqlRepository)
}
