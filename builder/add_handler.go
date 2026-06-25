package builder

import (
	"fmt"
	"os"

	"github.com/muhfaris/rocket/builder/hexagonal"
	"github.com/muhfaris/rocket/config"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/spf13/viper"
)

// AddHandler generates a single handler for the given operationId
// in the specified project directory using its rocket.yaml config.
func AddHandler(projectDir, openapiFilePath, operationID, ignoreDataResponse string) error {
	// Load OpenAPI spec
	content, doc, err := libos.LoadOpenapi(openapiFilePath)
	if err != nil {
		return fmt.Errorf("load openapi: %w", err)
	}

	// Read rocket.yaml from the project directory
	// The config is expected at <projectDir>/rocket.yaml
	cfg, err := loadRocketConfig(projectDir)
	if err != nil {
		return fmt.Errorf("load rocket config: %w", err)
	}

	projectName := cfg.App.Project
	packagePath := cfg.App.Package

	if projectName == "" || packagePath == "" {
		return fmt.Errorf("rocket.yaml must have 'app.project' and 'app.package' set")
	}

	// Set arch layout for template resolution
	archLayout := cfg.App.Arch
	if archLayout == "" {
		archLayout = "hexagonal"
	}
	templates.SetArchLayout(archLayout)

	// Build the hexagonal project context
	based := hexagonal.Based{
		Project: hexagonal.BaseProject{
			AppName:     projectName,
			ProjectName: projectName,
			PackagePath: packagePath,
		},
		Package: hexagonal.BasePackage{},
	}

	ignoreDataResponseBool := ignoreDataResponse == "true"
	if ignoreDataResponse == "" {
		ignoreDataResponseBool = false
	}

	project := hexagonal.NewProject(doc, &hexagonal.Config{
		Based:              based,
		ProjectName:        projectName,
		CacheParam:         cfg.App.Cache,
		DBParam:            cfg.App.Database,
		IgnoreDataResponse: ignoreDataResponseBool,
	})

	// Copy the OpenAPI spec into the project's spec/ directory if different
	// (only needed if user is pointing to an updated spec)
	specDir := fmt.Sprintf("%s/spec", projectName)
	if err := os.MkdirAll(specDir, os.ModePerm); err != nil {
		return fmt.Errorf("ensure spec dir: %w", err)
	}
	specFile := fmt.Sprintf("%s/openapi.yaml", specDir)
	if err := libos.CreateFile(specFile, content); err != nil {
		return fmt.Errorf("copy spec: %w", err)
	}

	// Delegate to hexagonal handler generation
	return project.AddHandler(operationID)
}

// loadRocketConfig reads the rocket.yaml from a project directory.
func loadRocketConfig(projectDir string) (config.Config, error) {
	var cfg config.Config

	viperInstance := viper.New()
	viperInstance.SetConfigName("rocket")
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath(projectDir)
	viperInstance.AddConfigPath(".")

	if err := viperInstance.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("read rocket.yaml in %s: %w", projectDir, err)
	}

	if err := viperInstance.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("parse rocket.yaml: %w", err)
	}

	return cfg, nil
}
