package hexagonal

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	libcase "github.com/muhfaris/rocket/shared/case"
	liboas "github.com/muhfaris/rocket/shared/oas"
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
	Annotation  string
	ParamsData
	QueryData
	BodyData
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
		FieldType: liboas.DataTypeToGo(fieldTypes[0]),
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
			operationID, _ = getOperationIDInfo(operation)
			_, structName  = libcase.Format(operationID)
			s              = Struct{StructName: structName}
			xParameterName = operation.Extensions["x-parameters-name"]
			hasStruct      bool
		)

		for _, parameter := range operation.Parameters {
			if parameter.Value.In == "path" {
				// addSuffix struct name
				s.StructName = fmt.Sprintf("%s%s", s.StructName, "Params")

				// if x-parameters-name is set override the struct name
				if xParameterName != nil {
					tempXParameterName, ok := xParameterName.(string)
					if ok {
						_, tempXParameterName = libcase.Format(tempXParameterName)
						s.StructName = tempXParameterName
					}
				}

				h.HasParams = true
				h.ParamsData.ParamsName = parameter.Value.Name
				h.ParamsData.ParamsStructName = s.StructName

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
				hasStruct = true
			}

			if parameter.Value.In == "query" {
				// add Suffix struct name
				s.StructName = fmt.Sprintf("%s%s", s.StructName, "Query")

				// if x-parameters-name is set override the struct name
				if xParameterName != nil {
					tempXParameterName, ok := xParameterName.(string)
					if ok {
						_, tempXParameterName = libcase.Format(tempXParameterName)
						s.StructName = tempXParameterName
					}
				}

				h.HasQuery = true
				h.QueryData.QueryName = parameter.Value.Name
				h.QueryData.QueryStructName = s.StructName

				field, err := h.parametersToField(parameter)
				if err != nil {
					return err
				}

				s.Fields = append(s.Fields, field)
				hasStruct = true
			}
		}

		if hasStruct {
			h.Structs = append(h.Structs, s)
			h.HasStructs = hasStruct
		}

		return nil

	case "POST", "PATCH", "PUT", "DELETE":
		var (
			operationID, _     = getOperationIDInfo(operation)
			_, structName      = libcase.Format(operationID)
			defaultStruct      = Struct{StructName: structName}
			tempXParameterName string
		)

		// Grouping all path parameters with same struct
		sParams := Struct{StructName: fmt.Sprintf("%s%s", structName, "Params")}
		for _, parameter := range operation.Parameters {
			xParameterName := parameter.Extensions["x-parameters-name"]
			if parameter.Value.In != "path" {
				continue
			}

			// if x-parameters-name is set override the struct name
			if xParameterName != nil && tempXParameterName == "" {
				xParameterNameStr, ok := xParameterName.(string)
				if ok {
					_, tempXParameterName = libcase.Format(xParameterNameStr)
					sParams.StructName = tempXParameterName
					tempXParameterName = xParameterNameStr
				}
			}

			field, err := h.parametersToField(parameter)
			if err != nil {
				return err
			}

			sParams.Fields = append(sParams.Fields, field)
		}

		if len(sParams.Fields) > 0 {
			h.Structs = append(h.Structs, sParams)
			h.HasStructs = true
			h.HasParams = true
			h.ParamsData.ParamsStructName = sParams.StructName
			h.ParamsData.ParamsName = libcase.ToLowerFirst(sParams.StructName)
			tempXParameterName = "" // reset
		}

		sQuery := Struct{StructName: fmt.Sprintf("%s%s", structName, "Query")}
		for _, parameter := range operation.Parameters {
			xParameterName := parameter.Extensions["x-parameters-name"]
			if parameter.Value.In != "query" {
				continue
			}

			// if x-parameters-name is set override the struct name
			if xParameterName != nil {
				xParameterNameStr, ok := xParameterName.(string)
				if ok {
					_, xParameterNameStr = libcase.Format(xParameterNameStr)
					sQuery.StructName = tempXParameterName
					tempXParameterName = xParameterNameStr
				}
			}

			field, err := h.parametersToField(parameter)
			if err != nil {
				return err
			}

			sQuery.Fields = append(sQuery.Fields, field)
		}

		if len(sQuery.Fields) > 0 {
			h.Structs = append(h.Structs, sQuery)
			h.HasStructs = true
			h.HasQuery = true
			h.QueryData.QueryName = libcase.ToLowerFirst(sQuery.StructName)
			h.QueryData.QueryStructName = sQuery.StructName
			tempXParameterName = ""
		} else {
		}

		if operation.RequestBody == nil {
			h.HasStructs = len(h.Structs) > 0
			return nil
		}

		requestBody := *operation.RequestBody
		if requestBody.Value == nil {
			return nil
		}

		// override struct name body body
		sBody := defaultStruct
		for contextType, content := range requestBody.Value.Content {
			xPropertiesName := content.Schema.Value.Extensions["x-properties-name"]

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
				if xPropertiesName != nil {
					tempXPropertiesName, ok := xPropertiesName.(string)
					if ok {
						_, tempPropertiesName := libcase.Format(tempXPropertiesName)
						sBody.StructName = tempPropertiesName
					}
				}

				for key, property := range content.Schema.Value.Properties {
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
						FieldType: liboas.DataTypeToGo(fieldTypes[0]),
						Tag:       fmt.Sprintf("`json:\"%s\"`", key),
					}

					sBody.Fields = append(sBody.Fields, field)
				}
				h.Structs = append(h.Structs, sBody)
			}
		}

		h.HasStructs = len(h.Structs) > 0
		return nil

	default:
		return nil
	}
}

func getOperationIDInfo(operation *openapi3.Operation) (operationID, serviceName string) {
	tempOperationID := strings.Split(operation.OperationID, "::")
	if len(tempOperationID) == 0 {
		return tempOperationID[0], ""
	}
	return tempOperationID[0], tempOperationID[1]
}
