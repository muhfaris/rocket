package liboas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	libcase "github.com/muhfaris/rocket/shared/case"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ResponseStruct struct {
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Fields   []PropertyStruct `json:"fields,omitempty"`
	Children *ResponseStruct  `json:"children,omitempty"`
}

type PropertyStruct struct {
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Enum     []string         `json:"enum,omitempty"`
	Format   string           `json:"format,omitempty"`
	Tag      string           `json:"tag,omitempty"`
	Children *ResponseStruct  `json:"children,omitempty"`
	Fields   []PropertyStruct `json:"fields,omitempty"`
}

func ParseSchema(parentStruct, contentType string, schema *openapi3.Schema, ignoreDataResponse bool) (ResponseStruct, error) {
	if schema == nil {
		return ResponseStruct{}, nil
	}

	sn, ok := schema.Extensions["x-struct-response"]
	if ok {
		snString, ok := sn.(string)
		if !ok && len(snString) == 0 {
			return ResponseStruct{}, errors.New("response should has x-struct-response as struct name")
		}

		_, df := libcase.Format(snString)
		parentStruct = strings.ToTitle(df[:1]) + df[1:]
	}

	var (
		schemaType = schema.Type.Slice()[0]
		res        = ResponseStruct{Name: parentStruct, Type: schemaType}
	)

	if parentStruct == "" {
		return ResponseStruct{}, errors.New("response should has x-struct-response as struct name")
	}

	if schemaType == "object" {
		for name, prop := range schema.Properties {
			var (
				nameLower = strings.ToLower(name)
				propType  = prop.Value.Type.Slice()[0]
			)
			_, name = libcase.Format(name)
			name = strings.ToTitle(name[:1]) + name[1:]

			enumFunc := func() []string {
				var res []string
				for _, v := range prop.Value.Enum {
					r, ok := v.(string)
					if ok {
						res = append(res, r)
					}
				}
				return res
			}

			/**
			 Check if the properties is data
			 properties:
				data:
			 		type: object
			*/
			sn, ok := prop.Value.Extensions["x-struct-response"]
			if ok {
				snString, ok := sn.(string)
				if !ok && len(snString) == 0 {
					return ResponseStruct{}, fmt.Errorf("response %s should has x-struct-response as struct name", name)
				}

				name = strings.ToTitle(snString[:1]) + snString[1:]
			}

			field := PropertyStruct{
				Name:   name,
				Type:   prop.Value.Type.Slice()[0],
				Enum:   enumFunc(),
				Format: prop.Value.Format,
				Tag:    GetTag(contentType, nameLower),
			}

			/**
			Type of properties is object
			properties:
				user:
					type: object
					properties:
						id :
							type: string
			*/
			if propType == "object" {
				nestedStruct, err := ParseSchema(name, contentType, prop.Value, ignoreDataResponse)
				if err != nil {
					return ResponseStruct{}, err
				}

				resFieldNameData := strings.ToLower(nestedStruct.Name)
				if ignoreDataResponse && resFieldNameData == "data" {
					raw, _ := json.Marshal(nestedStruct)
					fmt.Println(string(raw))
					return nestedStruct, nil
				}

				return ResponseStruct{
					Name: res.Name,
					Type: res.Type,
					Fields: []PropertyStruct{
						{
							Name:   nestedStruct.Name,
							Type:   nestedStruct.Type,
							Fields: nestedStruct.Fields,
							Tag:    GetTag(contentType, nameLower),
						},
					},
				}, nil
			}

			/**
			Type of properties is array
			properties:
				user:
					type: array
					properties:
						id :
							type: string
			*/
			if propType == "array" && prop.Value.Items != nil {
				structsName := strings.Split(prop.Value.Items.RefString(), "/")
				structName := structsName[len(structsName)-1]
				structName = strings.ToTitle(structName[:1]) + structName[1:]

				arrayStruct, err := ParseSchema(structName, contentType, prop.Value.Items.Value, ignoreDataResponse)
				if err != nil {
					return ResponseStruct{}, err
				}

				resFieldNameData := strings.ToLower(name)
				if ignoreDataResponse && resFieldNameData == "data" {
					return ResponseStruct{
						Name: res.Name,
						Type: propType,
						Children: &ResponseStruct{
							Name:   arrayStruct.Name,
							Type:   arrayStruct.Type,
							Fields: arrayStruct.Fields,
						},
					}, nil
				}

				return ResponseStruct{
					Name: res.Name,
					Type: res.Type,
					Fields: []PropertyStruct{
						{
							Name: name,
							Type: propType,
							Tag:  GetTag(contentType, nameLower),
							Children: &ResponseStruct{
								Name:   arrayStruct.Name,
								Type:   arrayStruct.Type,
								Fields: arrayStruct.Fields,
							},
						},
					},
				}, nil
			}

			res.Fields = append(res.Fields, field)
		}

		return res, nil

	} else if schemaType == "array" && schema.Items != nil {
		sn, ok := schema.Extensions["x-struct-response"]
		if ok {
			snString, ok := sn.(string)
			if !ok {
				return ResponseStruct{}, fmt.Errorf("array items should has x-struct-response")
				// return ResponseStruct{}, nil
			}

			parentStruct = strings.ToTitle(snString)
		}

		arrayStruct, err := ParseSchema(parentStruct, contentType, schema.Items.Value, ignoreDataResponse)
		if err != nil {
			return ResponseStruct{}, err
		}

		res.Fields = append(res.Fields, PropertyStruct{
			Name:     parentStruct,
			Type:     "array",
			Children: &arrayStruct,
		})
		return res, nil
	}
	return res, nil
}

func ToCamelCase(s string) string {
	if strings.Contains(s, "_") {
		ss := strings.Split(s, "_")
		for i, v := range ss {
			ss[i] = cases.Title(language.English).String(v)
		}
		return strings.Join(ss, "")
	}

	return cases.Title(language.English).String(s)
}

func GetTag(contentType, data string) string {
	switch contentType {
	case "application/json":
		return fmt.Sprintf("`json:\"%s\"`", data)
	case "application/xml":
		return "xml"
	default:
		return data
	}
}
