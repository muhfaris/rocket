package cmd

import (
	"fmt"

	"github.com/muhfaris/rocket/builder"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/muhfaris/rocket/shared/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openapiCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create new project",
	Example: "new --package github.com/muhfaris/myproject --project myproject --openapi myopenapi.yaml",
	RunE:    openapiRunE,
}

func openapiRunE(cmd *cobra.Command, args []string) error {
	var (
		packageNameParam = viper.Get("package")
		projectNameParam = viper.Get("project")
		openapiFileParam = viper.Get("openapi")
	)

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

	openapiFilePath, ok := openapiFileParam.(string)
	if !ok {
		return fmt.Errorf("openapi file must be string")
	}

	if openapiFilePath == "" {
		return fmt.Errorf("openapi file must be set")
	}

	archLayout := viper.Get("arch").(string)
	templates.SetArchLayout(archLayout)

	if has := utils.ContainsSpaceOrSpecialChar(projectName); has {
		return fmt.Errorf("project name can't contain space or special character")
	}

	content, doc, err := libos.LoadOpenapi(openapiFilePath)
	if err != nil {
		return err
	}

	m := builder.New(content, doc, packageName, projectName)
	return m.Generate()
}
