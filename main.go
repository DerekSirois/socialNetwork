package main

import (
	"project/socialnetwork/routes"
	"project/socialnetwork/server"
)

func main() {
	srv := server.New()
	srv.OpenDatabase()

	routes.Routes(srv)
}
