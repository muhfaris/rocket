package main

import "github.com/muhfaris/rocket/examples/books-api/cmd"

// @title Books API
// @description Demonstrates rocket generator features: route groups, path/query params, request body, response structs, and schema references.
// @version 1.0.0
// @host http://localhost:8080
// @securityDefinitions.basic bearerAuth
// @securityDefinitions.basic noauthAuth

func main() {
	cmd.Execute()
}
