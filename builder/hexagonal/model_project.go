package hexagonal

import "github.com/getkin/kin-openapi/openapi3"

type Project struct {
	based                 Based
	doc                   *openapi3.T
	cacheType             string
	dbType                string
	AppRepository         AppRepository
	AppService            AppService
	Dirs                  []string
	Rest                  Rest
	GroupRest             GroupRest
	RestRouter            RestRouter
	RestPortAdapter       RestPortAdapter
	RestMiddlewares       RestMiddlewares
	SharedLibrary         SharedLibrary
	RestResponse          RestResponse
	RoutesGroup           []RouterGroup
	RestPortService       RestPortService
	Domains               DomainModel
	RegistryService       RegistryService
	Service               Service
	RedisAdapter          RedisAdapter
	RedisCommandAdapter   RedisCommandAdapter
	CacheRepository       CacheRepository
	PSQLAdapter           PSQLAdapter
	PSQLCommandAdapter    PSQLCommandAdapter
	PSQLRepository        PSQLRepository
	PSQLQueryRepository   PSQLQueryRepository
	MySQLAdapter          MySQLAdapter
	MySQLCommandAdapter   MySQLCommandAdapter
	MySQLRepository       MySQLRepository
	MySQLQueryRepository  MySQLQueryRepository
	SQLiteAdapter         SQLiteAdapter
	SQLiteCommandAdapter  SQLiteCommandAdapter
	SQLiteRepository      SQLiteRepository
	SQLiteQueryRepository SQLiteQueryRepository
	MongoAdapter          MongoAdapter
	MongoCommandAdapter   MongoCommandAdapter
	MongoRepository       MongoRepository
	Dockerfile            Dockerfile
	DockerCompose         DockerCompose
	Makefile              Makefile
	ReadmeFile            ReadmeFile
	MethodRepository      MethodRepository
	APIError              APIError
}

type Based struct {
	Package BasePackage
	Project BaseProject
}

type BasePackage struct {
	PackageName string
	PackagePath string
}

type BaseProject struct {
	AppName     string
	ProjectName string
	PackagePath string
}

type AppRepository struct {
	dirpath  string
	filepath string
	template []byte
}

type AppService struct {
	dirpath  string
	filepath string
	template []byte
}

type Service struct {
	dirpath  string
	filepath string
	template []byte
	Services []ServiceParams
}

type ServiceParams struct {
	PackagePath string
	ServiceName string
	Methods     []PortServiceMethods
}

type RegistryService struct {
	dirpath  string
	filepath string
	template []byte
	Services []string
}

type DomainModel struct {
	dirpath  string
	template []byte
	Data     []DataDomainModel
}

type DataDomainModel struct {
	filename string
	Structs  []Struct
}

type RestPortService struct {
	dirpath  string
	filepath string
	template []byte
	Data     DataRestPortService
}

type DataRestPortService struct {
	ServiceName string
	Methods     []PortServiceMethods
}

type PortServiceMethods struct {
	MethodName  string
	Params      []PortServiceMethodParams
	ReturnTypes []PortServiceReturnType
}

type PortServiceMethodParams struct {
	Name string
	Type string
}

type PortServiceReturnType struct {
	Type string
}

type RestResponse struct {
	dirpath  string
	template []byte
	filepath string
}

type Rest struct {
	templateCmd []byte
	dirpathCmd  string
	filepathCmd string
	entrypoint  string
}

type GroupRest struct {
	template []byte
	dirpath  string
	filepath string
}

type RestRouter struct {
	template []byte
	dirpath  string
	filepath string
}

type RestPortAdapter struct {
	dirpath  string
	filepath string
	template []byte
}

type RestMiddlewares struct {
	dirpath string
	data    []DataRestMiddleware
}

type DataRestMiddleware struct {
	filepaths string
	template  []byte
}

type SharedLibrary struct {
	dirpath string
	data    []DataSharedLibrary
}

type DataSharedLibrary struct {
	name     string
	filepath string
	template []byte
}

type RedisAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type RedisCommandAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type CacheRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type PSQLAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type PSQLCommandAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type PSQLRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type PSQLQueryRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type MySQLAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type MySQLCommandAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type MySQLRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type MySQLQueryRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type SQLiteAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type SQLiteCommandAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type SQLiteRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type SQLiteQueryRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type MongoAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type MongoCommandAdapter struct {
	template []byte
	dirpath  string
	filepath string
}

type MongoRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type Dockerfile struct {
	filepath string
	template []byte
}

type DockerCompose struct {
	filepath string
	template []byte
}

type Makefile struct {
	filepath string
	template []byte
}

type ReadmeFile struct {
	filepath string
	template []byte
}

type MethodRepository struct {
	template []byte
	dirpath  string
	filepath string
}

type APIError struct {
	template []byte
	dirpath  string
	filepath string
}
