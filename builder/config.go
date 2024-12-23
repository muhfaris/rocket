package builder

import (
	"fmt"
	"os"
	"strings"

	"github.com/muhfaris/rocket/shared/constanta"
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
	HasDatabase    bool
	IsPSQL         bool
	IsMySQL        bool
	IsSQLite       bool
	IsMongoDB      bool
	AppName        string
}

func NewConfig(configName, configType, projectName, cacheType, dbType string) *Config {
	return &Config{
		AppName:        strings.ToLower(projectName),
		template:       templates.GetConfigTemplate(),
		configTemplate: templates.GetConfigFileTemplate(),
		dirpath:        fmt.Sprintf("%s/config", projectName),
		filepath:       fmt.Sprintf("%s/config/%s.go", projectName, configName),
		ConfigName:     configName,
		ConfigType:     configType,
		IsCache:        cacheType != "",
		IsRedis:        cacheType == constanta.CacheRedis,
		HasDatabase:    dbType != "",
		IsPSQL:         dbType == constanta.DBPostgres,
		IsMySQL:        dbType == constanta.DBMySQL,
		IsSQLite:       dbType == constanta.DBSQLite,
		IsMongoDB:      dbType == constanta.DBMongo,
	}
}

func (c *Config) Generate() error {
	fmt.Printf("%s%s\n", LineOnProgress, "config")
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
		"IsRedis":     c.IsRedis,
		"IsPSQL":      c.IsPSQL,
		"HasDatabase": c.HasDatabase,
		"IsMySQL":     c.IsMySQL,
		"IsSQLite":    c.IsSQLite,
		"IsMongoDB":   c.IsMongoDB,
		"AppName":     c.AppName,
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
