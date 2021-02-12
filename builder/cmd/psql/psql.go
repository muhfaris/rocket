package psql

type psqlBuilder struct {
	path         *psqlPath
	functionName string
	options      *psqlOptionsBuilder
}

type psqlPath struct {
	parent       string
	sub          string
	host         string
	port         string
	username     string
	password     string
	databaseName string
	sslMode      string
}

func newPSQLPath() *psqlPath {
	return &psqlPath{
		parent:       "persistence",
		sub:          "database",
		host:         "host",
		port:         "port",
		username:     "username",
		password:     "password",
		databaseName: "name",
		sslMode:      "ssl_mode",
	}
}

func newPSQLBuilder() *psqlBuilder {
	return &psqlBuilder{
		options:      newPSQLOptionsBuilder(),
		path:         newPSQLPath(),
		functionName: "initPSQL",
	}
}

type psqlOptionsBuilder struct {
	name     string
	username string
	password string
	host     string
	port     int
	sslMode  string
}

func newPSQLOptionsBuilder() *psqlOptionsBuilder {
	return &psqlOptionsBuilder{
		name:     "rocket_database",
		username: "",
		password: "",
		host:     "127.0.0.1",
		port:     5342,
		sslMode:  "disable",
	}
}
