package hexagonal

import (
	"fmt"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/muhfaris/rocket/helper/ui"
	libcase "github.com/muhfaris/rocket/shared/case"
	"github.com/muhfaris/rocket/shared/constanta"
	liboas "github.com/muhfaris/rocket/shared/oas"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/muhfaris/rocket/shared/utils"
)

type Config struct {
	Based              Based
	ProjectName        string
	CacheParam         string
	DBParam            string
	IgnoreDataResponse bool
}

var config *Config

func NewProject(doc *openapi3.T, cfg *Config) *Project {
	// re-assign config
	config = cfg
	projectName := cfg.ProjectName

	return &Project{
		based:     cfg.Based,
		doc:       doc,
		cacheType: cfg.CacheParam,
		dbType:    cfg.DBParam,
		AppRepository: AppRepository{
			dirpath:  fmt.Sprintf("%s/cmd/bootstrap", cfg.ProjectName),
			filepath: fmt.Sprintf("%s/cmd/bootstrap/app_repository.go", cfg.ProjectName),
			template: templates.GetAppRepositoryTemplate(),
		},
		AppService: AppService{
			dirpath:  fmt.Sprintf("%s/cmd/bootstrap", cfg.ProjectName),
			filepath: fmt.Sprintf("%s/cmd/bootstrap/app_service.go", cfg.ProjectName),
			template: templates.GetAppServiceTemplate(),
		},
		Dirs: []string{
			"internal/adapter/inbound/rest/router",
			"internal/adapter/inbound/rest/router/v1/handler",
			"internal/adapter/inbound/rest/router/v1/middleware",
			"internal/adapter/inbound/rest/router/v1/response",
			"internal/core/domain",
			"internal/core/service",
			"internal/core/port/inbound/adapter",
			"internal/core/port/inbound/registry",
			"internal/core/port/inbound/service",
			"internal/core/port/outbound",
			"internal/core/port/outbound/datastore",
			"shared",
		},
		Rest: Rest{
			templateCmd: templates.GetRestTemplate(),
			dirpathCmd:  fmt.Sprintf("%s/cmd", projectName),
			filepathCmd: fmt.Sprintf("%s/cmd/rest.go", projectName),
			entrypoint:  "rest",
		},
		GroupRest: GroupRest{
			template: templates.GetGroupRestTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/inbound/rest/router/group", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/inbound/rest/router/group/v1.go", projectName),
		},
		RestRouter: RestRouter{
			template: templates.GetRestRouterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/inbound/rest/router", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/inbound/rest/router/router.go", projectName),
		},
		RestPortAdapter: RestPortAdapter{
			dirpath:  fmt.Sprintf("%s/internal/core/port/inbound/adapter", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/inbound/adapter/rest.go", projectName),
			template: templates.GetRestAdapterTemplate(),
		},
		RestMiddlewares: RestMiddlewares{
			dirpath: fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/middleware", projectName),
			data: []DataRestMiddleware{
				{
					filepaths: fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/middleware/latency.go", projectName),
					template:  templates.GetRestLatencyMiddlewareTemplate(),
				},
			},
		},
		RestResponse: RestResponse{
			dirpath:  fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/response", projectName),
			template: templates.GetRestResponseTemplate(),
			filepath: "response.go",
		},
		SharedLibrary: SharedLibrary{
			dirpath: fmt.Sprintf("%s/shared", projectName),
			data: []DataSharedLibrary{
				{
					name:     "context",
					filepath: "context.go",
					template: templates.GetSharedContextTemplate(),
				},
			},
		},
		RestPortService: RestPortService{
			dirpath:  fmt.Sprintf("%s/internal/core/port/inbound/service", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/inbound/service/service.go", projectName),
			template: templates.GetRestPortServiceTemplate(),
		},
		Domains: DomainModel{
			dirpath:  fmt.Sprintf("%s/internal/core/domain", projectName),
			template: templates.GetDomainModel(),
		},
		RegistryService: RegistryService{
			dirpath:  fmt.Sprintf("%s/internal/core/port/inbound/registry", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/inbound/registry/registry.go", projectName),
			template: templates.GetRegistryServiceTemplate(),
		},
		Service: Service{
			dirpath:  fmt.Sprintf("%s/internal/core/service", projectName),
			filepath: fmt.Sprintf("%s/internal/core/service/%%s.go", projectName),
			template: templates.GetServiceTemplate(),
		},
		RedisAdapter: RedisAdapter{
			template: templates.GetRedisAdapterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/cache/redis", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/cache/redis/redis.go", projectName),
		},
		RedisCommandAdapter: RedisCommandAdapter{
			template: templates.GetRedisCommandTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/cache/redis", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/cache/redis/command.go", projectName),
		},
		CacheRepository: CacheRepository{
			template: templates.GetRedisRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/cache.go", projectName),
		},
		PSQLAdapter: PSQLAdapter{
			template: templates.GetPSQLAdapterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql/psql.go", projectName),
		},
		PSQLCommandAdapter: PSQLCommandAdapter{
			template: templates.GetPSQLCommandTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql/command.go", projectName),
		},
		PSQLRepository: PSQLRepository{
			template: templates.GetPSQLRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/psql.go", projectName),
		},
		PSQLQueryRepository: PSQLQueryRepository{
			template: templates.GetPSQLQueryRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/psql/repository/%%s.go", projectName),
		},
		MySQLAdapter: MySQLAdapter{
			template: templates.GetMySQLAdapterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql/mysql.go", projectName),
		},
		MySQLCommandAdapter: MySQLCommandAdapter{
			template: templates.GetMySQLCommandTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql/command.go", projectName),
		},
		MySQLRepository: MySQLRepository{
			template: templates.GetMySQLRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/mysql.go", projectName),
		},
		MySQLQueryRepository: MySQLQueryRepository{
			template: templates.GetMySQLQueryRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/mysql/repository/%%s.go", projectName),
		},
		SQLiteAdapter: SQLiteAdapter{
			template: templates.GetSQLiteAdapterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite/sqlite.go", projectName),
		},
		SQLiteCommandAdapter: SQLiteCommandAdapter{
			template: templates.GetSQLiteCommandTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite/command.go", projectName),
		},
		SQLiteRepository: SQLiteRepository{
			template: templates.GetSQLiteRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/sqlite.go", projectName),
		},
		SQLiteQueryRepository: SQLiteQueryRepository{
			template: templates.GetSQLiteQueryRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/sqlite/repository/%%s.go", projectName),
		},
		MongoAdapter: MongoAdapter{
			template: templates.GetMongoDBAdapterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/mongo", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/mongo/mongo.go", projectName),
		},
		MongoCommandAdapter: MongoCommandAdapter{
			template: templates.GetMongoDBCommandTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/outbound/datastore/mongo", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/outbound/datastore/mongo/command.go", projectName),
		},
		MongoRepository: MongoRepository{
			template: templates.GetMongoDBRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/mongo.go", projectName),
		},
		Dockerfile: Dockerfile{
			template: templates.GetDockerfileTemplate(),
			filepath: fmt.Sprintf("%s/Dockerfile", projectName),
		},
		DockerCompose: DockerCompose{
			template: templates.GetDockerComposeTemplate(),
			filepath: fmt.Sprintf("%s/docker-compose.yml", projectName),
		},
		Makefile: Makefile{
			template: templates.GetMakefileTemplate(),
			filepath: fmt.Sprintf("%s/Makefile", projectName),
		},
		ReadmeFile: ReadmeFile{
			template: templates.GetReadmeTemplate(),
			filepath: fmt.Sprintf("%s/README.md", projectName),
		},
		MethodRepository: MethodRepository{
			template: templates.GetMethodRepositoryTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/core/port/outbound/repository", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/outbound/repository/%%s.go", projectName),
		},
		APIError: APIError{
			template: templates.GetAPIErrorTemplate(),
			dirpath:  fmt.Sprintf("%s/shared/apierror", projectName),
			filepath: fmt.Sprintf("%s/shared/apierror/apierror.go", projectName),
		},
	}
}

func (p *Project) GenerateDirectories() error {
	// slog.Info("└── Creating based project directories")
	fmt.Printf("%s%s\n", ui.LineLast, "Creating based project directories")
	for _, dir := range p.Dirs {
		dirpath := fmt.Sprintf("%s/%s", p.based.Project.ProjectName, dir)
		_, err := os.Stat(dirpath)
		if os.IsExist(err) {
			continue
		}

		err = os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Generate cmd rest
	err := p.GenerateRest()
	if err != nil {
		return err
	}

	// Generate rest handlers
	err = p.GenerateRestHandlers()
	if err != nil {
		return err
	}

	// Generate Group rest router
	err = p.GenerateGroupRestRouter()
	if err != nil {
		return err
	}

	// Generate rest router
	err = p.GenerateRestRouter()
	if err != nil {
		return err
	}

	// Generate rest adapter
	err = p.GenerateRestPortAdapter()
	if err != nil {
		return err
	}

	err = p.GenerateRegistryService()
	if err != nil {
		return err
	}

	// Generate domain
	err = p.GenerateDomainModel()
	if err != nil {
		return err
	}

	// Generate service adapter
	err = p.GenerateRestPortService()
	if err != nil {
		return err
	}

	// Generate service
	err = p.GenerateRestService()
	if err != nil {
		return err
	}

	// Generate rest middlewares
	err = p.GenerateRestMiddlewares()
	if err != nil {
		return err
	}

	// Generate shared library
	err = p.GenerateSharedLibrary()
	if err != nil {
		return err
	}

	// Generate rest response
	err = p.GenerateRestResponse()
	if err != nil {
		return err
	}

	// Generate App Repository
	err = p.GenerateAppRepository()
	if err != nil {
		return err
	}

	// Generate App Service
	err = p.GenerateAppService()
	if err != nil {
		return err
	}

	// Generate redis adapter
	err = p.GenerateRedisAdapter()
	if err != nil {
		return err
	}

	// Generate cache repository
	err = p.GenerateCacheRepository()
	if err != nil {
		return err
	}

	// Generate psql adapter
	err = p.GeneratePSQLAdapter()
	if err != nil {
		return err
	}

	// Generate psql repository
	err = p.GeneratePSQLRepository()
	if err != nil {
		return err
	}

	err = p.GeneratePSQLQueryRepository()
	if err != nil {
		return err
	}

	// Generate mysql adapter
	err = p.GenerateMySQLAdapter()
	if err != nil {
		return err
	}

	// Generate mysql repository
	err = p.GenerateMySQLRepository()
	if err != nil {
		return err
	}

	err = p.GenerateMySQLQueryRepository()
	if err != nil {
		return err
	}

	// Generate sqlite adapter
	err = p.GenerateSQLiteAdapter()
	if err != nil {
		return err
	}

	// Generate sqlite repository
	err = p.GenerateSQLiteRepository()
	if err != nil {
		return err
	}

	err = p.GenerateSQLiteQueryRepository()
	if err != nil {
		return err
	}

	// Generate mongo adapter
	err = p.GenerateMongoAdapter()
	if err != nil {
		return err
	}

	// Generate mongo repository
	err = p.GenerateMongoRepository()
	if err != nil {
		return err
	}

	err = p.GenerateMethodRepository()
	if err != nil {
		return err
	}

	// Generated api error
	err = p.GenerateAPIError()
	if err != nil {
		return err
	}

	// Generate dockerfile
	err = p.GenerateDockerfile()
	if err != nil {
		return err
	}

	// Generate docker compose
	err = p.GenerateDockerCompose()
	if err != nil {
		return err
	}

	// Generate makefile
	err = p.GenerateMakefile()
	if err != nil {
		return err
	}

	// Generate readme
	err = p.GenerateReadmeFile()
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRest() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.Rest.dirpathCmd)
	_, err := os.Stat(p.Rest.dirpathCmd)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.Rest.dirpathCmd, os.ModePerm)
		if err != nil {
			return err
		}
	}

	restData := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"Entrypoint":  p.Rest.entrypoint,
	}

	raw, err := libos.ExecuteTemplate(p.Rest.templateCmd, restData)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.Rest.filepathCmd, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateGroupRestRouter() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.GroupRest.dirpath)
	_, err := os.Stat(p.GroupRest.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.GroupRest.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	routerData := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"AppName":     p.based.Project.AppName,
		"Groups":      p.RoutesGroup,
	}

	raw, err := libos.ExecuteTemplate(p.GroupRest.template, routerData)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.GroupRest.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestRouter() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RestRouter.dirpath)
	_, err := os.Stat(p.RestRouter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RestRouter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	routerData := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"AppName":     p.based.Project.AppName,
		"Groups":      p.RoutesGroup,
	}

	raw, err := libos.ExecuteTemplate(p.RestRouter.template, routerData)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RestRouter.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestPortAdapter() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RestPortAdapter.dirpath)
	_, err := os.Stat(p.RestPortAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RestPortAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	raw, err := libos.ExecuteTemplate(p.RestPortAdapter.template, nil)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RestPortAdapter.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestPortService() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RestPortService.dirpath)
	_, err := os.Stat(p.RestPortService.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RestPortService.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"ServiceName": p.RestPortService.Data.ServiceName,
		"Methods":     p.RestPortService.Data.Methods,
	}
	raw, err := libos.ExecuteTemplate(p.RestPortService.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RestPortService.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestMiddlewares() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RestMiddlewares.dirpath)
	_, err := os.Stat(p.RestMiddlewares.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RestMiddlewares.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dataMiddleware := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	for _, middleware := range p.RestMiddlewares.data {
		raw, err := libos.ExecuteTemplate(middleware.template, dataMiddleware)
		if err != nil {
			return err
		}

		err = libos.CreateFile(middleware.filepaths, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateSharedLibrary() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.SharedLibrary.dirpath)
	_, err := os.Stat(p.SharedLibrary.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.SharedLibrary.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, data := range p.SharedLibrary.data {
		path := fmt.Sprintf("%s/%s", p.SharedLibrary.dirpath, data.name)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}

		raw, err := libos.ExecuteTemplate(data.template, nil)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf("%s/%s/%s", p.SharedLibrary.dirpath, data.name, data.filepath)
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", filepath, err)
		}
	}

	return nil
}

func (p *Project) GenerateRestResponse() error {
	fmt.Printf(" %s%s\n", ui.LineLast, p.RestResponse.dirpath)
	_, err := os.Stat(p.RestResponse.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RestResponse.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(p.RestResponse.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RestResponse.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dataResponse := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.RestResponse.template, dataResponse)
	if err != nil {
		return err
	}

	filepath := fmt.Sprintf("%s/%s", p.RestResponse.dirpath, p.RestResponse.filepath)
	err = libos.CreateFile(filepath, raw)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filepath, err)
	}

	return nil
}

func (p *Project) GenerateRestHandlers() error {
	var (
		childsRouter       []ChildRouterGroup
		handlerDir         = fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/handler", p.based.Project.ProjectName)
		presenterDir       = fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/presenter", p.based.Project.ProjectName)
		routesGroupMap     = make(map[string]RouterGroup)
		domainMap          = make(map[string]DataDomainModel)
		serviceRegistryMap = make(map[string]bool)
		serviceHandler     DataRestPortService
		servicesMap        = make(map[string]ServiceParams)
	)

	fmt.Printf(" %s%s\n", ui.LineOnProgress, handlerDir)

	for path, pathItem := range p.doc.Paths.Map() {
		var (
			groupRoute     = "routeGroup"
			groupRoutePath = "/"
			serviceName    = "AppSvc"
		)

		for method, operation := range pathItem.Operations() {
			if len(operation.Tags) == 0 {
				return fmt.Errorf("operation must have tags at path '%s'", path)
			}

			filenameDomain := fmt.Sprintf("%s.go", libcase.ToSnakeCase(operation.Tags[0]))
			domainModel := DataDomainModel{filename: filenameDomain}

			operationsID := strings.Split(operation.OperationID, "::")
			if len(operationsID) == 0 {
				return fmt.Errorf("operation id can't be empty at path '%s'", path)
			}

			operationID := operationsID[0]

			if len(operationsID) == 2 {
				serviceName = operationsID[1]
			}

			childRouter := ChildRouterGroup{
				Method:  libcase.ToTitleCase(method), // method fiber
				Path:    utils.ConvertBracesToColon(path),
				Handler: operationsID[0], // handler name
			}

			xRouteGroupAny := operation.Extensions["x-route-group"]
			xRouteGroup, ok := xRouteGroupAny.(string)
			if ok && xRouteGroup != "" {
				xRouteGroups := strings.Split(xRouteGroup, "::")
				if len(xRouteGroups) != 2 {
					return fmt.Errorf("invalid x-route-group format'%s' at path %s", xRouteGroup, path)
				}

				groupRoute = xRouteGroups[0]
				groupRoutePath = xRouteGroups[1]
			}

			existRoute, exist := routesGroupMap[groupRoute]
			if !exist {
				routesGroupMap[groupRoute] = RouterGroup{
					GroupName: groupRoute,
					GroupPath: groupRoutePath,
					Routes:    []ChildRouterGroup{childRouter},
				}
			} else {
				existRoute.Routes = append(existRoute.Routes, childRouter)
				routesGroupMap[groupRoute] = existRoute
			}

			childsRouter = append(childsRouter, childRouter)

			handlerData := &HandlerData{
				PackagePath: p.based.Project.PackagePath,
				HandlerName: operationsID[0],
				Structs:     make([]Struct, 0),
				HasParams:   false,
				HasQuery:    false,
				HasBody:     false,
				ParamsData:  ParamsData{},
				QueryData:   QueryData{},
				BodyData:    BodyData{},
				DomainModel: Struct{},
			}

			err := handlerData.Generate(method, operation)
			if err != nil {
				return err
			}

			// create presenter file
			err = p.createPresenterFile(presenterDir, handlerData)
			if err != nil {
				return err
			}

			// Prepare service handler
			svcMethodFunc := strings.ReplaceAll(operationID, "Handler", "") // Handler name, remnove Handler for service name

			handlerService := PortServiceMethods{
				MethodName:  svcMethodFunc,
				ReturnTypes: []PortServiceReturnType{{Type: "error"}},
			}

			// override service handler
			if handlerData.HasStructsResponse {
				handlerService = PortServiceMethods{
					MethodName: svcMethodFunc,
					ReturnTypes: []PortServiceReturnType{
						{Type: fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName)},
						{Type: "error"},
					},
				}
				handlerData.ServiceHasReturn = true
			}

			// if handlerData.HasQuery {
			// 	structType := fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName)
			// 	if strings.Contains(handlerData.QueryStructName, "map") {
			// 		structType = handlerData.QueryStructName
			// 	}
			// 	handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
			// 		Name: handlerData.QueryName,
			// 		Type: structType,
			// 	})
			// }
			//
			// if handlerData.HasParams {
			// 	structType := fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName)
			// 	if strings.Contains(handlerData.ParamsStructName, "map") {
			// 		structType = handlerData.ParamsStructName
			// 	}
			//
			// 	handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
			// 		Name: handlerData.ParamsName,
			// 		Type: structType,
			// 	})
			// }
			//
			// if handlerData.HasBody {
			// 	structType := fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName)
			// 	if strings.Contains(handlerData.BodyStructName, "map") {
			// 		structType = handlerData.BodyStructName
			// 	}
			//
			// 	handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
			// 		Name: handlerData.BodyName,
			// 		Type: structType,
			// 	})
			// }

			handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
				Name: "payload",
				Type: fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName),
			})

			// service handler
			serviceHandler.Methods = append(serviceHandler.Methods, handlerService)
			serviceHandler.ServiceName = serviceName

			// service service
			if _, exist = servicesMap[serviceName]; !exist {
				servicesMap[serviceName] = ServiceParams{
					PackagePath: p.based.Project.PackagePath,
					ServiceName: serviceName,
					Methods:     serviceHandler.Methods,
				}
			} else {
				service := servicesMap[serviceName]
				service.Methods = append(service.Methods, handlerService)
				servicesMap[serviceName] = service
			}

			// handler call service
			handlerData.HasService = true
			handlerData.Service = handlerService
			handlerData.ServiceName = serviceName

			// assigning domain model
			existDomain, exist := domainMap[domainModel.filename]
			if exist {
				existDomain.Structs = append(existDomain.Structs, handlerData.DomainModel)
				domainMap[domainModel.filename] = existDomain
			} else {
				domainModel.Structs = append(domainModel.Structs, handlerData.DomainModel)
				domainMap[domainModel.filename] = domainModel
			}

			annotation, err := liboas.CreateSwaggerAnnotation(path, method, operation)
			if err != nil {
				return err
			}

			handlerData.Annotation = annotation
			// create handler file
			err = p.createHandlerFile(handlerDir, handlerData)
			if err != nil {
				return err
			}

			// add to service registry
			serviceRegistryMap[serviceName] = true

		} // end look operations
	} // end look paths

	var routesGroup []RouterGroup
	for _, routeGroup := range routesGroupMap {
		routesGroup = append(routesGroup, routeGroup)
	}

	p.RoutesGroup = routesGroup
	p.RestPortService.Data = serviceHandler

	var domainsModel []DataDomainModel
	for _, dm := range domainMap {
		domainsModel = append(domainsModel, dm)
	}

	var servicesRegistry []string
	for k := range serviceRegistryMap {
		servicesRegistry = append(servicesRegistry, k)
	}

	p.RegistryService.Services = servicesRegistry
	p.Domains.Data = domainsModel

	var services []ServiceParams
	for _, service := range servicesMap {
		services = append(services, service)
	}

	p.Service.Services = services
	return nil
}

func (p *Project) GenerateDomainModel() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.Domains.dirpath)
	_, err := os.Stat(p.Domains.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.Domains.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, dataStruct := range p.Domains.Data {
		raw, err := libos.ExecuteTemplate(p.Domains.template, dataStruct)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf("%s/%s", p.Domains.dirpath, dataStruct.filename)
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

// createHandlerFile is create 2 file
// 1. handler.go
// 2. <handler namne>.go
func (p *Project) createHandlerFile(handlerDir string, handlerData *HandlerData) error {
	initData := map[string]any{
		"PackagePath": handlerData.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(templates.GetRestInitHandlerTemplate(), initData)
	if err != nil {
		return err
	}

	initFilepath := fmt.Sprintf("%s/handler.go", handlerDir)
	err = libos.CreateFile(initFilepath, raw)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", initFilepath, err)
	}

	raw, err = libos.ExecuteTemplate(templates.GetRestHandlerTemplate(), handlerData)
	if err != nil {
		return err
	}

	var (
		filename = fmt.Sprintf("%s.go", libcase.ToSnakeCase(handlerData.HandlerName))
		filepath = fmt.Sprintf("%s/%s", handlerDir, filename)
	)

	_, err = os.Stat(handlerDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(handlerDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = libos.CreateFile(filepath, raw)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filepath, err)
	}

	return nil
}

func (p *Project) createPresenterFile(presenterDir string, handlerData *HandlerData) error {
	var (
		filename = fmt.Sprintf("%s.go", libcase.ToSnakeCase(handlerData.HandlerName))
		filepath = fmt.Sprintf("%s/%s", presenterDir, filename)
	)

	// prepare struct response
	if len(handlerData.StructsResponse) > 0 {
		fieldsDomainMap := make(map[string]Field)
		for _, fieldDomain := range handlerData.DomainModel.Fields {
			fieldsDomainMap[fieldDomain.FieldName] = fieldDomain
		}

		fieldsStructResponseMap := make(map[string]string)
		for _, structResponse := range handlerData.StructsResponse {
			for _, fieldSR := range structResponse.Fields {
				fieldDomain, ok := fieldsDomainMap[fieldSR.FieldName]
				if ok {
					fieldsStructResponseMap[fieldSR.FieldName] = fieldDomain.FieldName
				}
			}
		}

		if len(fieldsStructResponseMap) > 0 {
			handlerData.HasStructsResponse = true
			handlerData.MappingFieldsStructResponse = fieldsStructResponseMap
		}
	}

	// end

	_, err := os.Stat(presenterDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(presenterDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	raw, err := libos.ExecuteTemplate(templates.GetRestPresenterTemplate(), handlerData)
	if err != nil {
		return err
	}

	err = libos.CreateFile(filepath, raw)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filepath, err)
	}

	return nil
}

func (p *Project) GenerateRegistryService() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RegistryService.dirpath)
	_, err := os.Stat(p.RegistryService.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RegistryService.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var repositories []string
	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		repositories = append(repositories, fmt.Sprintf("%sRepository", repositoryName))
	}

	data := map[string]any{
		"PackagePath":  p.based.Project.PackagePath,
		"Services":     p.RegistryService.Services,
		"IsCache":      p.cacheType != "",
		"IsRedis":      p.cacheType == "redis",
		"Repositories": repositories,
	}
	raw, err := libos.ExecuteTemplate(p.RegistryService.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RegistryService.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestService() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.Service.dirpath)

	_, err := os.Stat(p.Service.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.Service.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, service := range p.Service.Services {
		data := map[string]any{
			"PackagePath": p.based.Project.PackagePath,
			"ServiceName": service.ServiceName,
			"Methods":     p.RestPortService.Data.Methods,
		}

		raw, err := libos.ExecuteTemplate(p.Service.template, data)
		if err != nil {
			return err
		}

		svcLowercase := strings.ToLower(service.ServiceName)
		filePath := fmt.Sprintf(p.Service.filepath, svcLowercase)
		err = libos.CreateFile(filePath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateAppRepository() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.AppRepository.dirpath)
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.AppRepository.filepath)
	_, err := os.Stat(p.AppRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.AppRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var repositories []string
	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		repositories = append(repositories, fmt.Sprintf("%sRepository", repositoryName))
	}

	data := map[string]any{
		"PackagePath":  p.based.Project.PackagePath,
		"IsRedis":      p.cacheType == constanta.CacheRedis,
		"IsPSQL":       p.dbType == constanta.DBPostgres,
		"IsMySQL":      p.dbType == constanta.DBMySQL,
		"IsSQLite":     p.dbType == constanta.DBSQLite,
		"IsMongo":      p.dbType == constanta.DBMongo,
		"Repositories": repositories,
	}

	raw, err := libos.ExecuteTemplate(p.AppRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.AppRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateAppService() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.AppService.dirpath)
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.AppService.filepath)
	_, err := os.Stat(p.AppService.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.AppService.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var services []string
	for _, service := range p.Service.Services {
		services = append(services, service.ServiceName)
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"Services":    services,
	}

	raw, err := libos.ExecuteTemplate(p.AppService.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.AppService.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRedisAdapter() error {
	if p.cacheType != constanta.CacheRedis {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RedisAdapter.dirpath)
	_, err := os.Stat(p.RedisAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RedisAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	rawRedisAdapter, err := libos.ExecuteTemplate(p.RedisAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RedisAdapter.filepath, rawRedisAdapter)
	if err != nil {
		return err
	}

	rawRedisCommand, err := libos.ExecuteTemplate(p.RedisCommandAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RedisCommandAdapter.filepath, rawRedisCommand)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) GenerateCacheRepository() error {
	if p.cacheType == "" {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.CacheRepository.dirpath)
	_, err := os.Stat(p.CacheRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.CacheRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.CacheRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.CacheRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GeneratePSQLAdapter() error {
	if p.dbType != constanta.DBPostgres {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.PSQLAdapter.dirpath)
	_, err := os.Stat(p.PSQLAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.PSQLAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	rawPSQLAdapter, err := libos.ExecuteTemplate(p.PSQLAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.PSQLAdapter.filepath, rawPSQLAdapter)
	if err != nil {
		return err
	}

	rawPSQLCommand, err := libos.ExecuteTemplate(p.PSQLCommandAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.PSQLCommandAdapter.filepath, rawPSQLCommand)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) GeneratePSQLRepository() error {
	if p.dbType != constanta.DBPostgres {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.PSQLRepository.dirpath)
	_, err := os.Stat(p.PSQLRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.PSQLRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.PSQLRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.PSQLRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GeneratePSQLQueryRepository() error {
	if p.dbType != constanta.DBPostgres {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.PSQLQueryRepository.dirpath)
	_, err := os.Stat(p.PSQLQueryRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.PSQLQueryRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		data := map[string]any{
			"PackagePath":    p.based.Project.PackagePath,
			"RepositoryName": fmt.Sprintf("%sRepository", repositoryName),
		}

		raw, err := libos.ExecuteTemplate(p.PSQLQueryRepository.template, data)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf(p.PSQLQueryRepository.filepath, strings.ToLower(repositoryName))
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateMySQLAdapter() error {
	if p.dbType != constanta.DBMySQL {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MySQLAdapter.dirpath)
	_, err := os.Stat(p.MySQLAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MySQLAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	rawMySQLAdapter, err := libos.ExecuteTemplate(p.MySQLAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MySQLAdapter.filepath, rawMySQLAdapter)
	if err != nil {
		return err
	}

	rawMySQLCommand, err := libos.ExecuteTemplate(p.MySQLCommandAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MySQLCommandAdapter.filepath, rawMySQLCommand)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) GenerateMySQLRepository() error {
	if p.dbType != constanta.DBMySQL {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MySQLRepository.dirpath)
	_, err := os.Stat(p.MySQLRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MySQLRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.MySQLRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MySQLRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateMySQLQueryRepository() error {
	if p.dbType != constanta.DBMySQL {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MySQLQueryRepository.dirpath)
	_, err := os.Stat(p.MySQLQueryRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MySQLQueryRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		data := map[string]any{
			"PackagePath":    p.based.Project.PackagePath,
			"RepositoryName": fmt.Sprintf("%sRepository", repositoryName),
		}

		raw, err := libos.ExecuteTemplate(p.MySQLQueryRepository.template, data)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf(p.MySQLQueryRepository.filepath, strings.ToLower(repositoryName))
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateSQLiteAdapter() error {
	if p.dbType != constanta.DBSQLite {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.SQLiteAdapter.dirpath)
	_, err := os.Stat(p.SQLiteAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.SQLiteAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	rawSQLiteAdapter, err := libos.ExecuteTemplate(p.SQLiteAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.SQLiteAdapter.filepath, rawSQLiteAdapter)
	if err != nil {
		return err
	}

	rawSQLiteCommand, err := libos.ExecuteTemplate(p.SQLiteCommandAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.SQLiteCommandAdapter.filepath, rawSQLiteCommand)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) GenerateSQLiteRepository() error {
	if p.dbType != constanta.DBSQLite {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.SQLiteRepository.dirpath)
	_, err := os.Stat(p.SQLiteRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.SQLiteRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.SQLiteRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.SQLiteRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateSQLiteQueryRepository() error {
	if p.dbType != constanta.DBSQLite {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.SQLiteQueryRepository.dirpath)
	_, err := os.Stat(p.SQLiteQueryRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.SQLiteQueryRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		data := map[string]any{
			"PackagePath":    p.based.Project.PackagePath,
			"RepositoryName": fmt.Sprintf("%sRepository", repositoryName),
		}

		raw, err := libos.ExecuteTemplate(p.SQLiteQueryRepository.template, data)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf(p.SQLiteQueryRepository.filepath, strings.ToLower(repositoryName))
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateMongoAdapter() error {
	if p.dbType != constanta.DBMongo {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MongoAdapter.dirpath)
	_, err := os.Stat(p.MongoAdapter.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MongoAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	rawPSQLAdapter, err := libos.ExecuteTemplate(p.MongoAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MongoAdapter.filepath, rawPSQLAdapter)
	if err != nil {
		return err
	}

	rawMongoCommand, err := libos.ExecuteTemplate(p.MongoCommandAdapter.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MongoCommandAdapter.filepath, rawMongoCommand)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) GenerateMongoRepository() error {
	if p.dbType != constanta.DBMongo {
		return nil
	}

	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MongoRepository.dirpath)
	_, err := os.Stat(p.MongoRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MongoRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
	}

	raw, err := libos.ExecuteTemplate(p.MongoRepository.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.MongoRepository.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateDockerfile() error {
	raw, err := libos.ExecuteTemplate(p.Dockerfile.template, nil)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.Dockerfile.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateDockerCompose() error {
	if p.dbType == "" && p.cacheType == "" {
		return nil
	}

	data := map[string]bool{
		"IsRedis":   p.cacheType == constanta.CacheRedis,
		"IsMySQL":   p.dbType == constanta.DBMySQL,
		"IsPSQL":    p.dbType == constanta.DBPostgres,
		"IsMongoDB": p.dbType == constanta.DBMongo,
	}

	raw, err := libos.ExecuteTemplate(p.DockerCompose.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.DockerCompose.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateMakefile() error {
	data := map[string]any{
		"AppName":  p.based.Project.AppName,
		"IsRedis":  p.cacheType == constanta.CacheRedis,
		"IsPSQL":   p.dbType == constanta.DBPostgres,
		"IsMySQL":  p.dbType == constanta.DBMySQL,
		"IsSQLite": p.dbType == constanta.DBSQLite,
		"IsMongo":  p.dbType == constanta.DBMongo,
	}
	raw, err := libos.ExecuteTemplate(p.Makefile.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.Makefile.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

// Readme
func (p *Project) GenerateReadmeFile() error {
	data := map[string]string{
		"ProjectName": p.based.Project.ProjectName,
		"PackagePath": p.based.Project.PackagePath,
		"AppName":     p.based.Project.AppName,
	}

	raw, err := libos.ExecuteTemplate(p.ReadmeFile.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.ReadmeFile.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateMethodRepository() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.MethodRepository.dirpath)
	_, err := os.Stat(p.MethodRepository.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.MethodRepository.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, service := range p.Service.Services {
		repositoryName := service.ServiceName

		for _, suffix := range []string{"svc", "Svc", "SVC"} {
			if strings.Contains(repositoryName, suffix) {
				repositoryName = strings.Replace(repositoryName, suffix, "", 1)
				break
			}
		}

		data := map[string]any{
			"PackagePath":    p.based.Project.PackagePath,
			"RepositoryName": fmt.Sprintf("%sRepository", repositoryName),
		}

		raw, err := libos.ExecuteTemplate(p.MethodRepository.template, data)
		if err != nil {
			return err
		}

		filepath := fmt.Sprintf(p.MethodRepository.filepath, strings.ToLower(repositoryName))
		err = libos.CreateFile(filepath, raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) GenerateAPIError() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.APIError.dirpath)
	_, err := os.Stat(p.APIError.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.APIError.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	raw, err := libos.ExecuteTemplate(p.APIError.template, nil)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.APIError.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}
