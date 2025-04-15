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

func GetGroupRestTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.GroupRestTemplate
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

func GetRestPresenterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.RestPresenterTemplate
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

func GetMySQLAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MySQLAdapterTemplate
	default:
		return nil
	}
}

func GetMySQLCommandTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MySQLCommandTemplate
	default:
		return nil
	}
}

func GetMySQLRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MySQLRepositoryTemplate
	default:
		return nil
	}
}

func GetSQLiteAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.SQLiteAdapterTemplate
	default:
		return nil
	}
}

func GetSQLiteCommandTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.SQLiteCommandTemplate
	default:
		return nil
	}
}

func GetSQLiteRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.SQLiteRepositoryTemplate
	default:
		return nil
	}
}

func GetMongoDBRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MongoDBRepositoryTemplate
	default:
		return nil
	}
}

func GetMongoDBAdapterTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MongoDBAdapterTemplate
	default:
		return nil
	}
}

func GetMongoDBCommandTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MongoDBCommandTemplate
	default:
		return nil
	}
}

func GetDockerfileTemplate() []byte {
	return hexagonal.DockerfileTemplate
}

func GetDockerComposeTemplate() []byte {
	return hexagonal.DockerCompose
}

func GetMakefileTemplate() []byte {
	return hexagonal.MakefileTemplate
}

func GetReadmeTemplate() []byte {
	return hexagonal.ReadmeTemplate
}

func GetPSQLQueryRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.PSQLQueryRepositoryTemplate
	default:
		return nil
	}
}

func GetMethodRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MethodRepositoryTemplate
	default:
		return nil
	}
}

func GetMySQLQueryRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.MySQLQueryRepositoryTemplate
	default:
		return nil
	}
}

func GetSQLiteQueryRepositoryTemplate() []byte {
	switch archType {
	case hexagonalType:
		return hexagonal.SQLiteQueryRepositoryTemplate
	default:
		return nil
	}
}
