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
	_tomlTemplate []byte

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
	return _tomlTemplate
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
