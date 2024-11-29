package templates

import "github.com/muhfaris/rocket/shared/templates/hexagonal"

var archType string

const (
	cleanCodeType = "cleancode"
	hexagonalType = "hexagonal"
)

func SetArchLayout(templateType string) {
	archType = templateType
}

func GetMainTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MainTemplate

	default:
		return nil
	}
}

func GetGitIgnore() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.GitIgnoreTemplate
	default:
		return nil
	}
}

func GetCMDTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.CMDTemplate
	default:
		return nil
	}
}

func GetConfigTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.ConfigTemplate
	default:
		return nil
	}
}

func GetRestTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestTemplate
	default:
		return nil
	}
}

func GetRestRouterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestRouterTemplate
	default:
		return nil
	}
}

func GetRestAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestAdapterTemplate
	default:
		return nil
	}
}

func GetRestLatencyMiddlewareTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestLatencyMiddlewareTemplate
	default:
		return nil
	}
}

func GetSharedContextTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.SharedContextTemplate
	default:
		return nil
	}
}

func GetRestHandlerTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestHandlerTemplate
	default:
		return nil
	}
}

func GetRestInitHandlerTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestInitHandlerTemplate
	default:
		return nil
	}
}

func GetRestResponseTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestResponseTemplate
	default:
		return nil
	}
}

func GetRestPortServiceTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestPortServiceTemplate
	default:
		return nil
	}
}

func GetDomainModel() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.DomainModel
	default:
		return nil
	}
}

func GetRegistryServiceTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RegistryServiceTemplate
	default:
		return nil
	}
}

func GetConfigFileTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.ConfigFileTemplate
	default:
		return nil
	}
}

func GetServiceTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestServiceTemplate
	default:
		return nil
	}
}

func GetAppTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.AppTemplate
	default:
		return nil
	}
}

func GetRedisAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RedisAdapterTemplate
	default:
		return nil
	}
}

func GetRedisCommandTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RedisCommandTemplate
	default:
		return nil
	}
}

func GetRedisRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RedisRepositoryTemplate
	default:
		return nil
	}
}

func GetPSQLAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.PSQLAdapterTemplate
	default:
		return nil
	}
}

func GetPSQLCommandTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.PSQLCommandTemplate
	default:
		return nil
	}
}

func GetPSQLRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.PSQLRepositoryTemplate
	default:
		return nil
	}
}
