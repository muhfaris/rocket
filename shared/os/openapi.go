package libos

import (
	"context"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func LoadOpenapi(path string) ([]byte, *openapi3.T, error) {
	var loader = &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	return content, doc, nil
}
