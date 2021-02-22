package configs

import (
	"fmt"
	"os"

	"github.com/muhfaris/rocket/builder/configs/templates"
	osHelper "github.com/muhfaris/rocket/helper/os"
	"github.com/naoina/toml"
)

type ConfigBuilder struct {
	Project     string
	Pkg         string
	CurrentPath string
}

// NewConfigBuilder is initialize for config package
func NewConfigBuilder(project string) *ConfigBuilder {
	pwd, err := os.Getwd()
	if err != nil {
		return &ConfigBuilder{}
	}

	return &ConfigBuilder{
		Project:     project,
		Pkg:         "configs",
		CurrentPath: pwd,
	}
}

// Generate is generate config file
func (c *ConfigBuilder) Generate() error {
	folder := fmt.Sprintf("%s/%s", c.Project, c.Pkg)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0755)
	}

	filePath := fmt.Sprintf("%s/%s/config.toml", c.CurrentPath, folder)
	var file *os.File
	f, ok := osHelper.FileExists(filePath)
	if !ok {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		file = f

	} else {
		file = f
	}

	var config = make(map[string]interface{})
	config["app"] = templates.App{}

	if err := toml.NewEncoder(file).Encode(&config); err != nil {
		return err
	}

	return nil
}
