package liboas

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// createAnnotation generates a single Swaggo annotation string for a given path and operation
func CreateSwaggerAnnotation(path string, method string, operation *openapi3.Operation) (string, error) {
	var annotation string

	annotation += fmt.Sprintf("// %s %s handler\n", method, path)
	if operation.Summary != "" {
		annotation += fmt.Sprintf("// @Summary %s\n", operation.Summary)
	}
	if operation.Description != "" {
		annotation += fmt.Sprintf("// @Description %s\n", operation.Description)
	}
	if len(operation.Tags) > 0 {
		annotation += fmt.Sprintf("// @Tags %s\n", operation.Tags)
	}

	// Parameters
	for _, paramRef := range operation.Parameters {
		param := paramRef.Value
		annotation += fmt.Sprintf("// @Param %s %s %s %v \"%s\"\n",
			param.Name,
			param.In,
			param.Schema.Value.Type,
			param.Required,
			param.Description)
	}

	// Request Body
	if operation.RequestBody != nil {
		for contentType, mediaType := range operation.RequestBody.Value.Content {
			annotation += fmt.Sprintf("// @Accept %s\n", contentType)
			fragment := mediaType.Schema.RefPath().Fragment
			if fragment != "" && strings.Contains(fragment, "/components/schemas/") {
				refName := strings.ReplaceAll(fragment, "/components/schemas/", "")
				annotation += fmt.Sprintf("// @Param body body presenter.%s true %q\n", refName, "Request body")
			}
		}
	}

	// Responses
	if operation.Responses != nil {
		for statusCode, response := range operation.Responses.Map() {
			for _, mediaType := range response.Value.Content {
				var description string
				if response.Value != nil {
					if response.Value.Description != nil {
						description = *response.Value.Description
					}
				}

				responseModel := mediaType.Schema.Ref
				if responseModel != "" && strings.Contains(responseModel, "#/components/schemas/") {
					responseModel = strings.ReplaceAll(responseModel, "#/components/schemas/", "")
				}

				annotation += fmt.Sprintf("// @Success %s {object} %s \"%s\"\n",
					statusCode,
					responseModel, // Reference to the schema
					description)
			}
		}
	}

	// Router annotation
	annotation += fmt.Sprintf("// @Router %s [%s]\n", path, method)
	annotation += "\n"

	return annotation, nil
}

// generateOASDescription creates Swaggo annotations for the API metadata (title, version, etc.)
func OASDescriptionSwagger(doc *openapi3.T) (string, error) {
	var annotation string
	if doc.Info.Title != "" {
		annotation += fmt.Sprintf("// @title %s\n", doc.Info.Title)
	}
	if doc.Info.Description != "" {
		annotation += fmt.Sprintf("// @description %s\n", doc.Info.Description)
	}
	if doc.Info.Version != "" {
		annotation += fmt.Sprintf("// @version %s\n", doc.Info.Version)
	}
	if doc.Info.Contact != nil {
		if doc.Info.Contact.Name != "" {
			annotation += fmt.Sprintf("// @contact.name %s\n", doc.Info.Contact.Name)
		}
		if doc.Info.Contact.Email != "" {
			annotation += fmt.Sprintf("// @contact.email %s\n", doc.Info.Contact.Email)
		}
	}

	if len(doc.Servers) > 0 {
		annotation += fmt.Sprintf("// @host %s\n", doc.Servers[0].URL)
	}

	// Security schemes
	if doc.Components != nil {
		if doc.Components.SecuritySchemes != nil {
			for name, securitySchemeRef := range doc.Components.SecuritySchemes {
				securityScheme := securitySchemeRef.Value
				if securityScheme != nil {
					switch securityScheme.Type {
					case "apiKey":
						annotation += fmt.Sprintf("// @securityDefinitions.apikey %s\n", name)
						annotation += fmt.Sprintf("// @in %s\n", securityScheme.In)
						annotation += fmt.Sprintf("// @name %s\n", securityScheme.Name)
					case "http":
						annotation += fmt.Sprintf("// @securityDefinitions.basic %s\n", name)
					case "oauth2":
						annotation += fmt.Sprintf("// @securityDefinitions.oauth2.application %s\n", name)
						if securityScheme.Flows != nil {
							if flow := securityScheme.Flows.ClientCredentials; flow != nil {
								annotation += fmt.Sprintf("// @tokenUrl %s\n", flow.TokenURL)
							}
						}
					case "openIdConnect":
						annotation += fmt.Sprintf("// @securityDefinitions.openId %s\n", name)
						annotation += fmt.Sprintf("// @openIdConnectUrl %s\n", securityScheme.OpenIdConnectUrl)
					}
				}
			}
		}
	}

	return annotation, nil
}

// reference: https://swagger.io/docs/specification/v3_0/data-models/data-types/?sbsearch=Types
func DataTypeToGo(dataType string) string {
	switch dataType {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "null", "nullable", "array":
		return "any"
	default:
		return "any"
	}
}
