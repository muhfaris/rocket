package hexagonal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/muhfaris/rocket/helper/ui"
	"github.com/muhfaris/rocket/shared/generate"
	libcase "github.com/muhfaris/rocket/shared/case"
	liboas "github.com/muhfaris/rocket/shared/oas"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/muhfaris/rocket/shared/utils"
)



// findOperationById locates a path+method+operation by operationId in the OpenAPI doc.
func (p *Project) findOperationById(operationID string) (path string, method string, operation *openapi3.Operation, err error) {
	for pPath, pathItem := range p.doc.Paths.Map() {
		for m, op := range pathItem.Operations() {
			if op.OperationID == operationID || strings.HasPrefix(op.OperationID, operationID+"::") {
				return pPath, m, op, nil
			}
		}
	}
	return "", "", nil, fmt.Errorf("operationId %q not found in OpenAPI spec", operationID)
}


// AddHandler generates handler, presenter, domain, service port, service impl,
// and router entries for a single operationId in an existing generated project.
func (p *Project) AddHandler(operationID string) error {
	path, method, operation, err := p.findOperationById(operationID)
	if err != nil {
		return err
	}

	if len(operation.Tags) == 0 {
		return fmt.Errorf("operation %q must have at least one tag", operationID)
	}

	fmt.Printf(" %s Adding handler for %s %s\n", ui.LineOnProgress, method, path)

	// --- Prepare handler data ---
	opsID, serviceName := getOperationIDInfo(operation)
	_, handlerStructName := libcase.Format(opsID)

	handlerData := &HandlerData{
		PackagePath: p.based.Project.PackagePath,
		HandlerName: opsID,
		Structs:     make([]Struct, 0),
		HasParams:   false,
		HasQuery:    false,
		HasBody:     false,
		ParamsData:  ParamsData{},
		QueryData:   QueryData{},
		BodyData:    BodyData{},
		DomainModel: Struct{StructName: handlerStructName},
	}

	err = handlerData.Generate(method, operation)
	if err != nil {
		return fmt.Errorf("generate handler data: %w", err)
	}

	// service handler
	svcMethodFunc := strings.ReplaceAll(opsID, "Handler", "")
	handlerService := PortServiceMethods{
		MethodName:  svcMethodFunc,
		ReturnTypes: []PortServiceReturnType{{Type: "error"}},
	}

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

	handlerService.Params = append(handlerService.Params, PortServiceMethodParams{
		Name: "payload",
		Type: fmt.Sprintf("domain.%s", handlerData.DomainModel.StructName),
	})

	handlerData.HasService = true
	handlerData.Service = handlerService
	handlerData.ServiceName = serviceName
	if handlerData.ServiceName == "" {
		handlerData.ServiceName = "AppSvc"
	}

	// annotation
	annotation, err := liboas.CreateSwaggerAnnotation(path, method, operation)
	if err != nil {
		return err
	}
	handlerData.Annotation = annotation

	// --- Determine directories from the existing project ---
	projectName := p.based.Project.ProjectName
	handlerDir := fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/handler", projectName)
	presenterDir := fmt.Sprintf("%s/internal/adapter/inbound/rest/router/v1/presenter", projectName)
	domainDir := fmt.Sprintf("%s/internal/core/domain", projectName)
	serviceDir := fmt.Sprintf("%s/internal/core/service", projectName)
	portServiceDir := fmt.Sprintf("%s/internal/core/port/inbound/service", projectName)
	groupRouterFile := fmt.Sprintf("%s/internal/adapter/inbound/rest/router/group/v1.go", projectName)

	// --- 1. Generate handler file ---
	fmt.Printf("  %s handler file\n", ui.LineOnProgress)
	err = p.createHandlerFile(handlerDir, handlerData)
	if err != nil {
		return fmt.Errorf("create handler file: %w", err)
	}

	// --- 2. Generate presenter file ---
	fmt.Printf("  %s presenter file\n", ui.LineOnProgress)
	err = p.createPresenterFile(presenterDir, handlerData)
	if err != nil {
		return fmt.Errorf("create presenter file: %w", err)
	}

	// --- 3. Domain model (append or create) ---
	fmt.Printf("  %s domain model\n", ui.LineOnProgress)
	tag := strings.ToLower(operation.Tags[0])
	domainFilename := fmt.Sprintf("%s.go", libcase.ToSnakeCase(tag))
	domainFilepath := fmt.Sprintf("%s/%s", domainDir, domainFilename)

	err = p.appendDomainModels(domainFilepath, handlerData)
	if err != nil {
		return fmt.Errorf("domain model: %w", err)
	}

	// --- 4. Port service (interface) ---
	fmt.Printf("  %s port service interface\n", ui.LineOnProgress)
	svcLower := strings.ToLower(handlerData.ServiceName)
	portServiceFilepath := fmt.Sprintf("%s/%s.go", portServiceDir, svcLower)

	err = p.appendPortServiceMethod(portServiceFilepath, handlerData.ServiceName, handlerService)
	if err != nil {
		return fmt.Errorf("port service: %w", err)
	}

	// --- 5. Service implementation ---
	fmt.Printf("  %s service implementation\n", ui.LineOnProgress)
	serviceFilepath := fmt.Sprintf("%s/%s.go", serviceDir, svcLower)

	err = p.appendServiceMethod(serviceFilepath, handlerData.ServiceName, handlerService)
	if err != nil {
		return fmt.Errorf("service impl: %w", err)
	}

	// --- 6. Group router ---
	fmt.Printf("  %s router registration\n", ui.LineOnProgress)
	err = p.appendRouterRegistration(groupRouterFile, path, method, opsID, operation)
	if err != nil {
		return fmt.Errorf("router: %w", err)
	}

	fmt.Printf(" %s Done. Handler '%s' added to project '%s'.\n", ui.LineLast, opsID, projectName)
	return nil
}


// appendDomainModels appends new structs to an existing domain file or creates one.
func (p *Project) appendDomainModels(domainFilepath string, handlerData *HandlerData) error {
	domainDir := fmt.Sprintf("%s/internal/core/domain", p.based.Project.ProjectName)
	if err := os.MkdirAll(domainDir, os.ModePerm); err != nil {
		return fmt.Errorf("ensure domain dir: %w", err)
	}

	if _, err := os.Stat(domainFilepath); os.IsNotExist(err) {
		// Create new domain file — filename is just the base name
		_, filename := filepath.Split(domainFilepath)
		dataStruct := DataDomainModel{
			filename: filename,
			Structs:  []Struct{handlerData.DomainModel},
		}

		raw, err := libos.ExecuteTemplate(templates.GetDomainModel(), dataStruct)
		if err != nil {
			return err
		}
		return libos.CreateFile(domainFilepath, raw)
	}

	// Append new structs to existing domain file
	// Build the struct definition text using the same format as the template
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\ntype %s struct {\n", handlerData.DomainModel.StructName))
	for _, f := range handlerData.DomainModel.Fields {
		sb.WriteString(fmt.Sprintf("\t%s %s %s\n", f.FieldName, f.FieldType, f.Tag))
	}
	sb.WriteString("}\n")

	return generate.AppendToFile(domainFilepath, sb.String())
}


// appendPortServiceMethod adds a new method to an existing port service interface.
func (p *Project) appendPortServiceMethod(filepath, serviceName string, method PortServiceMethods) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// Create new port service file
		data := map[string]any{
			"PackagePath": p.based.Project.PackagePath,
			"ServiceName": serviceName,
			"Methods":     []PortServiceMethods{method},
		}

		raw, err := libos.ExecuteTemplate(templates.GetRestPortServiceTemplate(), data)
		if err != nil {
			return err
		}
		return libos.CreateFile(filepath, raw)
	}

	// Build method signature
	var sb strings.Builder
	returnTypes := buildReturnTypeString(method.ReturnTypes)
	params := buildParamString(method.Params)

	sb.WriteString(fmt.Sprintf("\n\t%s(ctx context.Context%s) %s",
		method.MethodName, params, returnTypes))

	// Insert before interface closing brace
	decl := fmt.Sprintf("type %s interface {", serviceName)
	return generate.InsertBeforePackageCloser(filepath, decl, sb.String())
}


// appendServiceMethod adds a new method implementation to an existing service file.
func (p *Project) appendServiceMethod(filepath, serviceName string, method PortServiceMethods) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// Create new service file
		data := map[string]any{
			"PackagePath": p.based.Project.PackagePath,
			"ServiceName": serviceName,
			"Methods":     []PortServiceMethods{method},
		}

		raw, err := libos.ExecuteTemplate(templates.GetServiceTemplate(), data)
		if err != nil {
			return err
		}
		return libos.CreateFile(filepath, raw)
	}

	// Build method implementation stub
	var sb strings.Builder
	returnTypes := buildReturnTypeString(method.ReturnTypes)
	params := buildParamString(method.Params)

	// Generate zero-value return
	returnVals := buildReturnValues(method.ReturnTypes)

	sb.WriteString(fmt.Sprintf("\nfunc (s *%s) %s(ctx context.Context%s) %s {",
		serviceName, method.MethodName, params, returnTypes))
	sb.WriteString(fmt.Sprintf("\n\treturn %s", returnVals))
	sb.WriteString("\n}\n")

	return generate.AppendToFile(filepath, sb.String())
}


// appendRouterRegistration adds a route to the group router file.
// If the group function already exists, it adds a route line inside it.
// Otherwise it creates a new group function and wires it into V1().
func (p *Project) appendRouterRegistration(groupRouterFile, path, method, opsID string, operation *openapi3.Operation) error {
	// Determine group info from x-route-group or default
	groupRoute := "routeGroup"
	groupRoutePath := "/"

	xRouteGroupAny := operation.Extensions["x-route-group"]
	xRouteGroup, ok := xRouteGroupAny.(string)
	if ok && xRouteGroup != "" {
		xRouteGroups := strings.Split(xRouteGroup, "::")
		if len(xRouteGroups) == 2 {
			groupRoute = xRouteGroups[0]
			groupRoutePath = xRouteGroups[1]
		}
	}

	routerPath := utils.ConvertBracesToColon(path)
	methodTitle := libcase.ToTitleCase(method)

	// Check if group function already exists
	groupFuncHeader := fmt.Sprintf("func %s(r *fiber.App, h *handlerv1.Handler) {", groupRoute)

	if _, err := os.Stat(groupRouterFile); os.IsNotExist(err) {
		// Create new group router file with just this route
		group := RouterGroup{
			GroupName: groupRoute,
			GroupPath: groupRoutePath,
			Routes: []ChildRouterGroup{
				{Method: methodTitle, Path: routerPath, Handler: opsID},
			},
		}

		data := map[string]any{
			"PackagePath": p.based.Project.PackagePath,
			"AppName":     p.based.Project.AppName,
			"Groups":      []RouterGroup{group},
		}

		raw, err := libos.ExecuteTemplate(templates.GetGroupRestTemplate(), data)
		if err != nil {
			return err
		}
		return libos.CreateFile(groupRouterFile, raw)
	}

	// Read existing file
	raw, err := os.ReadFile(groupRouterFile)
	if err != nil {
		return fmt.Errorf("read group router: %w", err)
	}

	txt := string(raw)

	if strings.Contains(txt, groupFuncHeader) {
		// Group function exists — add route line inside it
		routeLine := fmt.Sprintf("\n\t%s.%s(\"%s\", h.%s()).Name(\"%s\")",
			groupRoute, methodTitle, routerPath, opsID, opsID)

		return generate.InsertBeforePackageCloser(groupRouterFile, groupFuncHeader, routeLine)
	}

	// Group function doesn't exist — append new group function + V1 call
	// First add the V1 call
	v1Call := fmt.Sprintf("\n\t%s(r, h)", groupRoute)
	v1Header := "func V1(r *fiber.App, h *handlerv1.Handler) {"
	err = generate.InsertBeforePackageCloser(groupRouterFile, v1Header, v1Call)
	if err != nil {
		return fmt.Errorf("add V1 call: %w", err)
	}

	// Then append the new group function
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nfunc %s(r *fiber.App, h *handlerv1.Handler) {\n", groupRoute))
	sb.WriteString(fmt.Sprintf("\t%s := r.Group(\"%s\")\n", groupRoute, groupRoutePath))
	sb.WriteString(fmt.Sprintf("\t%s.%s(\"%s\", h.%s()).Name(\"%s\")\n",
		groupRoute, methodTitle, routerPath, opsID, opsID))
	sb.WriteString("}\n")

	return generate.AppendToFile(groupRouterFile, sb.String())
}


// --- helper functions ---

func buildReturnTypeString(types []PortServiceReturnType) string {
	if len(types) == 0 {
		return ""
	}
	parts := make([]string, len(types))
	for i, t := range types {
		parts[i] = t.Type
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

func buildParamString(params []PortServiceMethodParams) string {
	if len(params) == 0 {
		return ""
	}
	parts := make([]string, len(params))
	for i, p := range params {
		parts[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
	}
	return ", " + strings.Join(parts, ", ")
}

func buildReturnValues(types []PortServiceReturnType) string {
	if len(types) == 0 {
		return ""
	}
	vals := make([]string, len(types))
	for i, t := range types {
		if t.Type == "error" {
			vals[i] = "nil"
		} else {
			vals[i] = t.Type + "{}"
		}
	}
	return strings.Join(vals, ", ")
}
