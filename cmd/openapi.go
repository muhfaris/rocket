package cmd

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/muhfaris/rocket/builder"
	"github.com/spf13/cobra"
)

var openapiCmd = &cobra.Command{
	Use:  "openapi",
	RunE: openapiRunE,
}

func openapiRunE(cmd *cobra.Command, args []string) error {
	var (
		ctx             = context.Background()
		loader          = &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
		pathOpenapiSpec = "spec/openapi.yaml"
	)

	doc, err := loader.LoadFromFile(pathOpenapiSpec)
	if err != nil {
		return err
	}

	fmt.Println(doc.Info.Title)

	m := builder.New("github.com/muhfaris/rocket-test", "rocket-test")
	return m.Generate()
}
