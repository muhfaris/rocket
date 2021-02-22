package builder

import (
	cmdBuilder "github.com/muhfaris/rocket/builder/cmd"
	configBuilder "github.com/muhfaris/rocket/builder/configs"
	"github.com/pkg/errors"
)

// Builder is wrap data
type Builder struct {
	Name     string // project name
	Database string // postgresql, mysql, mongodb
	Cache    string // redis
	Queue    string
	NetHTTP  string // mux, gin, echo
	Config   string // toml, yaml, env
}

// NewBuilder is initialize the builder
func NewBuilder(name, db, cache, queue, config string) *Builder {
	return &Builder{
		Name:     name,
		Database: db,
		Cache:    cache,
		Queue:    queue,
		Config:   config,
	}
}

func (builder *Builder) Generate() error {
	if err := builder.validate(); err != nil {
		return err
	}

	cmdBuilder := cmdBuilder.NewCMDBuilder(builder.Name, builder.Database, builder.Cache, builder.Queue)
	err := cmdBuilder.Generate()
	if err != nil {
		return err
	}

	configBuilder := configBuilder.NewConfigBuilder(builder.Name)
	err = configBuilder.Generate()
	if err != nil {
		return err
	}

	return nil
}

func (builder *Builder) validate() error {
	if findDatabase(builder.Database) {
		return errors.Errorf("database %s not available", builder.Database)
	}

	if findCache(builder.Cache) {
		return errors.Errorf("cache %s not available", builder.Cache)
	}

	if findConfig(builder.Config) {
		return errors.Errorf("config %s not available", builder.Config)
	}

	return nil
}
