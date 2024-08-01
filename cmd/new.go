package cmd

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/muhfaris/rocket/builder"
	"github.com/muhfaris/rocket/shared/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openapiCmd = &cobra.Command{
	Use:  "new",
	RunE: openapiRunE,
}

func openapiRunE(cmd *cobra.Command, args []string) error {
	var (
		ctx             = context.Background()
		loader          = &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
		pathOpenapiSpec = "spec/openapi.yaml"
	)

	packageNameParam := viper.Get("package")
	projectNameParam := viper.Get("project")

	if packageNameParam == "" || projectNameParam == "" {
		return fmt.Errorf("package and project name must be set")
	}

	packageName, ok := packageNameParam.(string)
	if !ok {
		return fmt.Errorf("package must be string")
	}

	projectName, ok := projectNameParam.(string)
	if !ok {
		return fmt.Errorf("project name must be string")
	}

	if has := utils.ContainsSpaceOrSpecialChar(projectName); has {
		return fmt.Errorf("project name can't contain space or special character")
	}

	doc, err := loader.LoadFromFile(pathOpenapiSpec)
	if err != nil {
		return err
	}

	m := builder.New(doc, packageName, projectName)
	return m.Generate()
}
