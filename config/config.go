package config

type Config struct {
	Openapi string    `mapstructure:"openapi"`
	App     AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Package  string `mapstructure:"package"`
	Project  string `mapstructure:"project"`
	Arch     string `mapstructure:"arch"`
	Docket   bool   `mapstructure:"docket"`
	Database string `mapstructure:"database"`
	Cache    string `mapstructure:"cache"`
}
