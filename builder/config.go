package builder

import (
	"fmt"
	"os"

	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

type Config struct {
	template       []byte
	configTemplate []byte
	dirpath        string
	filepath       string
	ConfigName     string
	ConfigType     string
	IsCache        bool
	IsRedis        bool
	IsDatabase     bool
	IsPSQL         bool
}

func NewConfig(configName, configType, projectName, cacheType, dbType string) *Config {
	return &Config{
		template:       templates.GetConfigTemplate(),
		configTemplate: templates.GetConfigFileTemplate(),
		dirpath:        fmt.Sprintf("%s/config", projectName),
		filepath:       fmt.Sprintf("%s/config/%s.go", projectName, configName),
		ConfigName:     configName,
		ConfigType:     configType,
		IsCache:        cacheType != "",
		IsRedis:        cacheType == "redis",
		IsDatabase:     dbType != "",
		IsPSQL:         dbType == "postgres",
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

	data := map[string]any{
		"IsRedis": c.IsRedis,
		"IsPSQL":  c.IsPSQL,
	}

	rawConfig, err := libos.ExecuteTemplate(c.configTemplate, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(fmt.Sprintf("%s/config.yaml", c.dirpath), rawConfig)
	if err != nil {
		return err
	}

	return nil
}
