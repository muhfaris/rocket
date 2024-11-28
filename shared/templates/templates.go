package templates

import "github.com/muhfaris/rocket/shared/templates/unclebob"

var archType string

func SetArchLayout(templateType string) {
	archType = templateType
}

func GetMainTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.MainTemplate

	default:
		return nil
	}
}

func GetGitIgnore() []byte {
	switch archType {
	case "unclebob":
		return unclebob.GitIgnoreTemplate
	default:
		return nil
	}
}

func GetCMDTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.CMDTemplate
	default:
		return nil
	}
}

func GetConfigTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.ConfigTemplate
	default:
		return nil
	}
}

func GetRestTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestTemplate
	default:
		return nil
	}
}

func GetRestRouterTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestRouterTemplate
	default:
		return nil
	}
}

func GetRestAdapterTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestAdapterTemplate
	default:
		return nil
	}
}

func GetRestLatencyMiddlewareTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestLatencyMiddlewareTemplate
	default:
		return nil
	}
}

func GetSharedContextTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.SharedContextTemplate
	default:
		return nil
	}
}

func GetRestHandlerTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestHandlerTemplate
	default:
		return nil
	}
}

func GetRestInitHandlerTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestInitHandlerTemplate
	default:
		return nil
	}
}

func GetRestResponseTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestResponseTemplate
	default:
		return nil
	}
}

func GetRestPortServiceTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestPortServiceTemplate
	default:
		return nil
	}
}

func GetDomainModel() []byte {
	switch archType {
	case "unclebob":
		return unclebob.DomainModel
	default:
		return nil
	}
}

func GetRegistryServiceTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RegistryServiceTemplate
	default:
		return nil
	}
}

func GetConfigFileTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.ConfigFileTemplate
	default:
		return nil
	}
}

func GetServiceTemplate() []byte {
	switch archType {
	case "unclebob":
		return unclebob.RestServiceTemplate
	default:
		return nil
	}
}
