package cmdproject

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/muhfaris/rocket/builder"
	"github.com/muhfaris/rocket/config"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
	"github.com/muhfaris/rocket/shared/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var cfg config.Config

var OpenapiCMD = &cobra.Command{
	Use:     "new",
	Short:   "Create new project",
	Example: "new --package github.com/muhfaris/myproject --project myproject --openapi myopenapi.yaml",
	PreRun: func(cmd *cobra.Command, args []string) {
		// if a config file is found, read it in.
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		err = viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatal(err)
		}
	},
	RunE: openapiRunE,
}

func init() {
	cobra.OnInitialize(initconfig)
}

func initconfig() {
	cfg := viper.GetString("config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		return
	}

	// search config in home directory with name "config" (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("rocket")
}

func openapiRunE(cmd *cobra.Command, args []string) (err error) {
	var (
		openapiFileParam = viper.Get("openapi")
		packageNameParam = viper.Get("app.package")
		projectNameParam = viper.Get("app.project")
		cacheParam       = viper.GetString("app.cache")
		dbParam          = viper.GetString("app.database")
	)

	// Header
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Configuration Overview")
	fmt.Println(strings.Repeat("=", 50))

	// Display config
	for key, value := range viper.GetStringMap("app") {
		fmt.Printf("%-20s : %v\n", key, value)
	}

	// Footer
	fmt.Println(strings.Repeat("=", 50))

	defer func() {
		if err != nil {
			fmt.Println(err)
		}

		raw, err := yaml.Marshal(cfg)
		if err != nil {
			return
		}

		path := fmt.Sprintf("%s/rocket.yaml", projectNameParam)
		err = libos.CreateFile(path, raw)
		if err != nil {
			return
		}
	}()

	if packageNameParam == "" || projectNameParam == "" {
		return fmt.Errorf("package and project name must be set")
	}

	packageName, ok := packageNameParam.(string)
	if !ok {
		return fmt.Errorf("package must be string")
	}

	err = utils.ValidateImportPath(packageName)
	if err != nil {
		return err
	}

	projectName, ok := projectNameParam.(string)
	if !ok {
		return fmt.Errorf("project name must be string")
	}

	projectName = utils.SanitizeString(projectName)

	openapiFilePath, ok := openapiFileParam.(string)
	if !ok {
		return fmt.Errorf("openapi file must be string")
	}

	if openapiFilePath == "" {
		return fmt.Errorf("openapi file must be set")
	}

	archLayout, ok := viper.Get("app.arch").(string)
	if !ok {
		return fmt.Errorf("arch must be string")
	}
	templates.SetArchLayout(archLayout)

	if has := utils.ContainsSpaceOrSpecialChar(projectName); has {
		return fmt.Errorf("project name can't contain space or special character")
	}

	content, doc, err := libos.LoadOpenapi(openapiFilePath)
	if err != nil {
		return err
	}

	ignoreDataResponseTxt := cfg.App.IgnoreDataResponse
	if ignoreDataResponseTxt == "" {
		ignoreDataResponseTxt = "true"
	}

	ignoreDataResponse, _ := strconv.ParseBool(ignoreDataResponseTxt)

	m := builder.New(content, doc,
		builder.ConfigBuilder{
			ProjectName:        projectName,
			PackagePath:        packageName,
			Arch:               archLayout,
			CacheParam:         cacheParam,
			DBParam:            dbParam,
			IgnoreDataResponse: ignoreDataResponse,
		})
	return m.Generate()
}
