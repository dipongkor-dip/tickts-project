package main

import (
	"gotickts/internal/config"
	"gotickts/internal/server"
)

func main() {
	// load environment variables
	cfg := config.LoadEnv()

	// database connection
	db := config.ConnectDatabase(cfg)

	// start server
	server.Start(db, cfg)
}
