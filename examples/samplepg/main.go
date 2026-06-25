package main

import "github.com/muhfaris/rocket/examples/samplepg/cmd"

// @title Laporkan API
// @version 1.0.0
// @host http://localhost
// @securityDefinitions.basic bearerAuth
// @securityDefinitions.basic noauthAuth

func main() {
	cmd.Execute()
}
