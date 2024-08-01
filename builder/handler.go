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
	ParamsData
	QueryData
	BodyData
}

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	FielName  string
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
		FielName:  fieldName,
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
				h.HasParams = true
				h.ParamsData.ParamsName = parameter.Value.Name
				_, h.ParamsData.ParamsStructName = libcase.Format(parameter.Value.Name)

				if xParameterName == nil {
					slog.Error("x-parameters-name is should be set", "operation", operation.OperationID)
				}

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
			}

			if parameter.Value.In == "query" {
				h.HasQuery = true
				h.QueryData.QueryName = parameter.Value.Name
				_, h.QueryData.QueryStructName = libcase.Format(parameter.Value.Name)

				if xParameterName == nil {
					slog.Error("x-parameters-name is should be set", "operation", operation.OperationID)
				}

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
			}
		}

		h.Structs = append(h.Structs, s)
		return nil

	case "POST":
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
				h.HasBody = true
				h.BodyData.BodyName = "bodyRequest"
				h.BodyData.BodyStructName = "map[string]any"
				return nil
			}

			fmt.Println(contextType)
			for key, property := range content.Schema.Value.Properties {
				fmt.Println(key, property)

			}
		}

		raw, _ := operation.RequestBody.MarshalJSON()

		fmt.Println(string(raw))
		h.HasBody = true
		return nil

	default:
		return nil
	}
}
