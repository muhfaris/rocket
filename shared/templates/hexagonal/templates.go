package hexagonal

import (
	_ "embed"
)

var (
	//go:embed main.tmpl
	MainTemplate []byte

	//go:embed gitignore.tmpl
	GitIgnoreTemplate []byte

	//go:embed cmd/root.tmpl
	CMDTemplate []byte

	//go:embed config/config.tmpl
	ConfigTemplate []byte

	//go:embed config/config-file.tmpl
	ConfigFileTemplate []byte

	//go:embed cmd/rest.tmpl
	RestTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/group/v1.tmpl
	GroupRestTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/router.tmpl
	RestRouterTemplate []byte

	//go:embed internal/core/port/inbound/adapter/rest.tmpl
	RestAdapterTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/middlewares/latency.tmpl
	RestLatencyMiddlewareTemplate []byte

	//go:embed shared/context/context.tmpl
	SharedContextTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/handlers/handler.tmpl
	RestHandlerTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/handlers/init.tmpl
	RestInitHandlerTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/presenter/presenter.tmpl
	RestPresenterTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/response/response.tmpl
	RestResponseTemplate []byte

	//go:embed internal/core/port/inbound/service/service.tmpl
	RestPortServiceTemplate []byte

	//go:embed internal/core/service/service.tmpl
	RestServiceTemplate []byte

	//go:embed internal/core/port/inbound/registry/registry.tmpl
	RegistryServiceTemplate []byte

	//go:embed internal/core/domain/domain.tmpl
	DomainModel []byte

	//go:embed cmd/bootstrap/app.tmpl
	AppTemplate []byte

	//go:embed internal/adapter/outbound/cache/redis/redis.tmpl
	RedisAdapterTemplate []byte

	//go:embed internal/adapter/outbound/cache/redis/command.tmpl
	RedisCommandTemplate []byte

	//go:embed internal/core/port/outbound/repository/cache.tmpl
	RedisRepositoryTemplate []byte

	//go:embed internal/adapter/outbound/datastore/psql/psql.tmpl
	PSQLAdapterTemplate []byte

	//go:embed internal/adapter/outbound/datastore/psql/command.tmpl
	PSQLCommandTemplate []byte

	//go:embed internal/core/port/outbound/repository/psql.tmpl
	PSQLRepositoryTemplate []byte

	//go:embed internal/adapter/outbound/datastore/mysql/mysql.tmpl
	MySQLAdapterTemplate []byte

	//go:embed internal/adapter/outbound/datastore/mysql/command.tmpl
	MySQLCommandTemplate []byte

	//go:embed internal/core/port/outbound/repository/mysql.tmpl
	MySQLRepositoryTemplate []byte

	//go:embed internal/adapter/outbound/datastore/sqlite/sqlite.tmpl
	SQLiteAdapterTemplate []byte

	//go:embed internal/adapter/outbound/datastore/sqlite/command.tmpl
	SQLiteCommandTemplate []byte

	//go:embed internal/core/port/outbound/repository/sqlite.tmpl
	SQLiteRepositoryTemplate []byte

	//go:embed internal/adapter/outbound/datastore/mongodb/mongodb.tmpl
	MongoDBAdapterTemplate []byte

	//go:embed internal/adapter/outbound/datastore/mongodb/command.tmpl
	MongoDBCommandTemplate []byte

	//go:embed internal/core/port/outbound/repository/mongodb.tmpl
	MongoDBRepositoryTemplate []byte

	//go:embed Dockerfile.tmpl
	DockerfileTemplate []byte

	//go:embed docker-compose.tmpl
	DockerCompose []byte

	//go:embed Makefile.tmpl
	MakefileTemplate []byte

	//go:embed README.tmpl
	ReadmeTemplate []byte
)
