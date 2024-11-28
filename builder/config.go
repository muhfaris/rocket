package builder

import (
	"fmt"
	"os"

	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

type Config struct {
	template   []byte
	configFile []byte
	dirpath    string
	filepath   string
	ConfigName string
	ConfigType string
}

func NewConfig(configName, configType, projectName string) *Config {
	return &Config{
		template:   templates.GetConfigTemplate(),
		configFile: templates.GetConfigFileTemplate(),
		dirpath:    fmt.Sprintf("%s/config", projectName),
		filepath:   fmt.Sprintf("%s/config/%s.go", projectName, configName),
		ConfigName: configName,
		ConfigType: configType,
	}
}

func (c *Config) Generate() error {
	fmt.Printf("%s%s\n", lineOnProgress, "config")
	_, err := os.Stat(c.dirpath)
	if os.IsExist(err) {
		return fmt.Errorf("directory config %s already exists", c.dirpath)
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(c.dirpath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory config %s: %w", c.dirpath, err)
		}
	}

	raw, err := libos.ExecuteTemplate(c.template, c)
	if err != nil {
		return err
	}

	err = libos.CreateFile(c.filepath, raw)
	if err != nil {
		return err
	}

	err = libos.CreateFile(fmt.Sprintf("%s/config.yaml", c.dirpath), c.configFile)
	if err != nil {
		return err
	}

	return nil
}
