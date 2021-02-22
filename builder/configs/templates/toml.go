package templates

// App is wrap app config
type App struct {
	Name string
	Env  string
	Port int
	Auth AppAuthentication
}

// AppAuthentication is wrap authentication app
type AppAuthentication struct {
	Secret string
}

// PostgreSQL is wrap postgresql database
type PostgreSQL struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	SSLMode  string
}

// Redis is wrap redis cache
type Redis struct {
	Host     string
	Port     int
	Username string
	Passn    string
}
