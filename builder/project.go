package builder

import (
	"fmt"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	libcase "github.com/muhfaris/rocket/shared/case"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/muhfaris/rocket/shared/utils"
)

type Project struct {
	doc             *openapi3.T
	Dirs            []string
	Rest            Rest
	RestRouter      RestRouter
	RespPortAdapter RestPortAdapter
	RestMiddlewares RestMiddlewares
	SharedLibrary   SharedLibrary
	RestResponse    RestResponse
	RoutesGroup     []RouterGroup
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

func NewProject(doc *openapi3.T, projectName string) *Project {
	return &Project{
		doc: doc,
		Dirs: []string{
			"internal/adapter/inbound/rest/routers",
			"internal/adapter/inbound/rest/routers/v1/handlers",
			"internal/adapter/inbound/rest/routers/v1/middlewares",
			"internal/adapter/inbound/rest/routers/v1/response",
			"internal/core/domain",
			"internal/core/service",
			"internal/core/port/inbound/adapter",
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
		RespPortAdapter: RestPortAdapter{
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
	}
}

func (p *Project) GenerateDirectories() error {
	// slog.Info("└── Creating based project directories")
	fmt.Printf("%s%s\n", lineLast, "Creating based project directories")
	for _, dir := range p.Dirs {
		dirpath := fmt.Sprintf("%s/%s", _baseproject.ProjectName, dir)
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

	return nil
}

func (p *Project) GenerateRest() error {
	fmt.Printf(" %s%s\n", lineOnProgress, p.Rest.dirpathCmd)
	_, err := os.Stat(p.Rest.dirpathCmd)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.Rest.dirpathCmd, os.ModePerm)
		if err != nil {
			return err
		}
	}

	restData := map[string]any{
		"PackagePath": _baseproject.PackagePath,
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
	fmt.Printf(" %s%s\n", lineOnProgress, p.RestRouter.dirpath)
	_, err := os.Stat(p.RestRouter.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RestRouter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	routerData := map[string]any{
		"PackagePath": _baseproject.PackagePath,
		"AppName":     _baseproject.AppName,
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
	fmt.Printf(" %s%s\n", lineOnProgress, p.RespPortAdapter.dirpath)
	_, err := os.Stat(p.RespPortAdapter.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RespPortAdapter.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	raw, err := libos.ExecuteTemplate(p.RespPortAdapter.template, nil)
	if err != nil {
		return err
	}

	err = libos.CreateFile(p.RespPortAdapter.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateRestMiddlewares() error {
	fmt.Printf(" %s%s\n", lineOnProgress, p.RestMiddlewares.dirpath)
	_, err := os.Stat(p.RestMiddlewares.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RestMiddlewares.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dataMiddleware := map[string]any{
		"PackagePath": _baseproject.PackagePath,
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
	fmt.Printf(" %s%s\n", lineOnProgress, p.SharedLibrary.dirpath)
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
			err := os.MkdirAll(path, os.ModePerm)
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
	fmt.Printf(" %s%s\n", lineLast, p.RestResponse.dirpath)
	_, err := os.Stat(p.RestResponse.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RestResponse.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(p.RestResponse.dirpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p.RestResponse.dirpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dataResponse := map[string]any{
		"PackagePath": _baseproject.PackagePath,
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

type RouterGroup struct {
	GroupName string
	GroupPath string
	Routes    []ChildRouterGroup
}

type ChildRouterGroup struct {
	Method  string
	Path    string
	Handler string
}

func (p *Project) GenerateRestHandlers() error {
	var (
		childsRouter   []ChildRouterGroup
		handlerDir     = fmt.Sprintf("%s/internal/adapter/inbound/rest/routers/v1/handlers", _baseproject.ProjectName)
		routesGroupMap = make(map[string]RouterGroup)
	)

	fmt.Printf(" %s%s\n", lineOnProgress, handlerDir)

	for path, pathItem := range p.doc.Paths.Map() {
		var (
			groupRoute     = "routeGroup"
			groupRoutePath = "/"
		)

		for method, operation := range pathItem.Operations() {
			childRouter := ChildRouterGroup{
				Method:  libcase.ToTitleCase(method), // method fiber
				Path:    utils.ConvertBracesToColon(path),
				Handler: operation.OperationID, // handler name
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
					Routes:    []ChildRouterGroup{childRouter}}
			} else {
				existRoute.Routes = append(existRoute.Routes, childRouter)
				routesGroupMap[groupRoute] = existRoute
			}

			childsRouter = append(childsRouter, childRouter)

			handlerData := &HandlerData{
				PackagePath: _baseproject.PackagePath,
				HandlerName: operation.OperationID,
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

			raw, err := libos.ExecuteTemplate(templates.GetRestHandlerTemplate(), handlerData)
			if err != nil {
				return err
			}

			var (
				filename = fmt.Sprintf("%s.go", libcase.ToSnakeCase(handlerData.HandlerName))
				filepath = fmt.Sprintf("%s/%s", handlerDir, filename)
			)

			_, err = os.Stat(handlerDir)
			if os.IsNotExist(err) {
				err := os.MkdirAll(handlerDir, os.ModePerm)
				if err != nil {
					return err
				}
			}

			err = libos.CreateFile(filepath, raw)
			if err != nil {
				return fmt.Errorf("error creating file %s: %w", filepath, err)
			}
		}
	}

	var routesGroup []RouterGroup
	for _, routeGroup := range routesGroupMap {
		routesGroup = append(routesGroup, routeGroup)
	}
	p.RoutesGroup = routesGroup
	return nil
}
