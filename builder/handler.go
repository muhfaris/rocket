package builder

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	libcase "github.com/muhfaris/rocket/shared/case"
)

type HandlerData struct {
	PackagePath string
	HandlerName string
	Structs     []Struct
	HasParams   bool
	HasQuery    bool
	HasBody     bool
	HasStructs  bool
	HasService  bool
	Service     PortServiceMethods
	ServiceName string
	ParamsData
	QueryData
	BodyData
}

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	FieldName string
	FieldType string
	Tag       string
}

type ParamsData struct {
	ParamsName       string
	ParamsStructName string
}

type QueryData struct {
	QueryName       string
	QueryStructName string
}

type BodyData struct {
	BodyName       string
	BodyStructName string
}

func (h *HandlerData) parametersToField(parameter *openapi3.ParameterRef) (Field, error) {
	// FieldType is parameter type
	fieldTypePtr := parameter.Value.Schema.Value.Type
	if fieldTypePtr == nil {
		return Field{}, nil
	}

	var (
		fieldTypes   = fieldTypePtr.Slice()
		_, fieldName = libcase.Format(parameter.Value.Name)
	)

	field := Field{
		FieldName: fieldName,
		FieldType: fieldTypes[0],
		Tag:       fmt.Sprintf("`params:\"%s\"`", parameter.Value.Name),
	}

	if parameter.Value.In == "query" {
		field.Tag = fmt.Sprintf("`query:\"%s\"`", parameter.Value.Name)
	}

	return field, nil
}

func (h *HandlerData) Generate(method string, operation *openapi3.Operation) error {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		var (
			xParameterName   = operation.Extensions["x-parameters-name"]
			parameterName, _ = xParameterName.(string)
			_, structName    = libcase.Format(parameterName)
			s                = Struct{StructName: structName}
		)

		for _, parameter := range operation.Parameters {
			if parameter.Value.In == "path" {
				if xParameterName == nil {
					err := fmt.Errorf("x-parameters-name is should be set, operation %s", operation.OperationID)
					slog.Error(err.Error())
					return err
				}

				h.HasParams = true
				h.ParamsData.ParamsName = parameter.Value.Name
				h.ParamsData.ParamsStructName = structName

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
			}

			if parameter.Value.In == "query" {
				if xParameterName == nil {
					slog.Error("x-parameters-name is should be set", "operation", operation.OperationID)
				}

				h.HasQuery = true
				h.QueryData.QueryName = parameter.Value.Name
				h.QueryData.QueryStructName = structName

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
			}
		}

		h.Structs = append(h.Structs, s)
		h.HasStructs = len(h.Structs) > 0
		return nil

	case "POST", "PATCH", "PUT", "DELETE":
		if operation.RequestBody == nil {
			return nil
		}

		requestBody := *operation.RequestBody
		if requestBody.Value == nil {
			return nil
		}

		for contextType, content := range requestBody.Value.Content {
			// if properties is empty use map
			if len(content.Schema.Value.Properties) == 0 {
				if !h.HasBody {
					h.HasBody = true
				}
				h.BodyData.BodyName = "bodyRequest"
				h.BodyData.BodyStructName = "map[string]any"
				return nil
			}

			if contextType == "application/json" {
				var (
					xPropertiesName   = content.Schema.Value.Extensions["x-properties-name"]
					propertiesName, _ = xPropertiesName.(string)
					_, structName     = libcase.Format(propertiesName)
					s                 = Struct{StructName: structName}
				)

				for key, property := range content.Schema.Value.Properties {
					if xPropertiesName == nil {
						err := fmt.Errorf("x-properties-name is should be set on properties, operation %s", operation.OperationID)
						slog.Error(err.Error())
						return err
					}

					if !h.HasBody {
						h.HasBody = true
						h.BodyData.BodyName = "bodyRequest"
						h.BodyData.BodyStructName = structName
					}

					var (
						fieldTypes   = property.Value.Type.Slice()
						_, fieldName = libcase.Format(key)
					)

					field := Field{
						FieldName: fieldName,
						FieldType: fieldTypes[0],
						Tag:       fmt.Sprintf("`json:\"%s\"`", key),
					}

					s.Fields = append(s.Fields, field)
				}
				h.Structs = append(h.Structs, s)
			}
		}

		h.HasStructs = len(h.Structs) > 0
		return nil

	default:
		return nil
	}
}
