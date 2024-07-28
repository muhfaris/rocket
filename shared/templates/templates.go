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
