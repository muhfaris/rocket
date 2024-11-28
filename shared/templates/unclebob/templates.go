package unclebob

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

	//go:embed config/config.yaml
	ConfigFileTemplate []byte

	//go:embed cmd/rest.tmpl
	RestTemplate []byte

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
)
