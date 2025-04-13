package config

type Config struct {
	Openapi string    `mapstructure:"openapi"`
	App     AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Package            string `mapstructure:"package"`
	Project            string `mapstructure:"project"`
	Arch               string `mapstructure:"arch"`
	Docket             bool   `mapstructure:"docker"`
	Database           string `mapstructure:"database"`
	Cache              string `mapstructure:"cache"`
	IgnoreDataResponse string `mapstructure:"ignore_data_response"`
}
