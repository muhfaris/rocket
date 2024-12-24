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

func NewProject(doc *openapi3.T, based Based, projectName, cacheParam, dbParam string) *Project {
	return &Project{
		based:     based,
		doc:       doc,
		cacheType: cacheParam,
		dbType:    dbParam,
		App: App{
			dirpath:  fmt.Sprintf("%s/internal/app", projectName),
			filepath: fmt.Sprintf("%s/internal/app/app.go", projectName),
			template: templates.GetAppTemplate(),
		},
		Dirs: []string{
			"internal/adapter/inbound/rest/routers",
			"internal/adapter/inbound/rest/routers/v1/handlers",
			"internal/adapter/inbound/rest/routers/v1/middlewares",
			"internal/adapter/inbound/rest/routers/v1/response",
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
		RestRouter: RestRouter{
			template: templates.GetRestRouterTemplate(),
			dirpath:  fmt.Sprintf("%s/internal/adapter/inbound/rest/routers", projectName),
			filepath: fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/router.go", projectName),
		},
		RestPortAdapter: RestPortAdapter{
			dirpath:  fmt.Sprintf("%s/internal/core/port/inbound/adapter", projectName),
			filepath: fmt.Sprintf("%s/internal/core/port/inbound/adapter/rest.go", projectName),
			template: templates.GetRestAdapterTemplate(),
		},
		RestMiddlewares: RestMiddlewares{
			dirpath: fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/v1/middlewares", projectName),
			data: []DataRestMiddleware{
				{
					filepaths: fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/v1/middlewares/latency.go", projectName),
					template:  templates.GetRestLatencyMiddlewareTemplate(),
				},
			},
		},
		RestResponse: RestResponse{
			dirpath:  fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/v1/response", projectName),
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

	// Generate App
	err = p.GenerateApp()
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
		handlerDir         = fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/v1/handlers", p.based.Project.ProjectName)
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
			}

			err := handlerData.Generate(method, operation)
			if err != nil {
				return err
			}

			// Prepare service handler
			svcMethodFunc := strings.ReplaceAll(operationID, "Handler", "") // Handler name, remnove Handler for service name
			handlerService := PortServiceMethods{
				MethodName: svcMethodFunc,
				ReturnTypes: []PortServiceReturnType{
					{Type: "error"},
				},
			}

			if handlerData.HasQuery {
				structType := fmt.Sprintf("domain.%s", handlerData.QueryData.QueryStructName)
				if strings.Contains(handlerData.QueryData.QueryStructName, "map") {
					structType = handlerData.QueryData.QueryStructName
				}
				handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
					Name: handlerData.QueryData.QueryName,
					Type: structType,
				})
			}

			if handlerData.HasParams {
				structType := fmt.Sprintf("domain.%s", handlerData.ParamsData.ParamsStructName)
				if strings.Contains(handlerData.ParamsData.ParamsStructName, "map") {
					structType = handlerData.ParamsData.ParamsStructName
				}

				handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
					Name: handlerData.ParamsData.ParamsName,
					Type: structType,
				})
			}

			if handlerData.HasBody {
				structType := fmt.Sprintf("domain.%s", handlerData.BodyData.BodyStructName)
				if strings.Contains(handlerData.BodyData.BodyStructName, "map") {
					structType = handlerData.BodyData.BodyStructName
				}

				handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
					Name: handlerData.BodyData.BodyName,
					Type: structType,
				})
			}

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

			// assinging domain model
			existDomain, exist := domainMap[domainModel.filename]
			if exist {
				existDomain.Structs = append(existDomain.Structs, handlerData.Structs...)
				domainMap[domainModel.filename] = existDomain
			} else {
				domainModel.Structs = handlerData.Structs
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

func (p *Project) GenerateRegistryService() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.RegistryService.dirpath)
	_, err := os.Stat(p.RegistryService.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.RegistryService.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"Services":    p.RegistryService.Services,
		"IsCache":     p.cacheType != "",
		"IsRedis":     p.cacheType == "redis",
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

func (p *Project) GenerateApp() error {
	fmt.Printf(" %s%s\n", ui.LineOnProgress, p.App.dirpath)
	_, err := os.Stat(p.App.dirpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p.App.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data := map[string]any{
		"PackagePath": p.based.Project.PackagePath,
		"IsRedis":     p.cacheType == constanta.CacheRedis,
		"IsPSQL":      p.dbType == constanta.DBPostgres,
		"IsMySQL":     p.dbType == constanta.DBMySQL,
		"IsSQLite":    p.dbType == constanta.DBSQLite,
		"IsMongo":     p.dbType == constanta.DBMongo,
	}

	raw, err := libos.ExecuteTemplate(p.App.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.App.filepath, raw)
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

	err = libos.CreateFile(p.PSQLAdapter.filepath, rawSQLiteAdapter)
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
