package templates

import (
	_ "embed"
)

var (
	//go:embed main.tmpl
	_mainTemplate []byte

	//go:embed gitignore.tmpl
	_gitIgnoreTemplate []byte

	//go:embed cmd/root.tmpl
	_cmdTemplate []byte

	//go:embed config/config.tmpl
	_configTemplate []byte

	//go:embed config/config.yaml
	_configFileTemplate []byte

	//go:embed cmd/rest.tmpl
	_restTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/router.tmpl
	_restRouterTemplate []byte

	//go:embed internal/core/port/inbound/adapter/rest.tmpl
	_restAdapterTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/middlewares/latency.tmpl
	_restLatencyMiddlewareTemplate []byte

	//go:embed shared/context/context.tmpl
	_sharedContextTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/handlers/handler.tmpl
	_restHandlerTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/handlers/init.tmpl
	_restInitHandlerTemplate []byte

	//go:embed internal/adapter/inbound/rest/routers/v1/response/response.tmpl
	_restResponseTemplate []byte

	//go:embed internal/core/port/inbound/service/service.tmpl
	_restPortServiceTemplate []byte

	//go:embed internal/core/port/inbound/registry/registry.tmpl
	_registryServiceTemplate []byte

	//go:embed internal/core/domain/domain.tmpl
	_domainModel []byte
)

func GetMainTemplate() []byte {
	return _mainTemplate
}

func GetGitIgnore() []byte {
	return _gitIgnoreTemplate
}

func GetCMDTemplate() []byte {
	return _cmdTemplate
}

func GetConfigTemplate() []byte {
	return _configTemplate
}

func GetRestTemplate() []byte {
	return _restTemplate
}

func GetRestRouterTemplate() []byte {
	return _restRouterTemplate
}

func GetRestAdapterTemplate() []byte {
	return _restAdapterTemplate
}

func GetRestLatencyMiddlewareTemplate() []byte {
	return _restLatencyMiddlewareTemplate
}

func GetSharedContextTemplate() []byte {
	return _sharedContextTemplate
}

func GetRestHandlerTemplate() []byte {
	return _restHandlerTemplate
}

func GetRestInitHandlerTemplate() []byte {
	return _restInitHandlerTemplate
}

func GetRestResponseTemplate() []byte {
	return _restResponseTemplate
}

func GetRestPortServiceTemplate() []byte {
	return _restPortServiceTemplate
}

func GetDomainModel() []byte {
	return _domainModel
}

func GetRegistryServiceTemplate() []byte {
	return _registryServiceTemplate
}

func GetConfigFileTemplate() []byte {
	return _configFileTemplate
}
