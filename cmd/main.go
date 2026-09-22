package main

import (
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/config"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/server"
)

func main() {
	cfg := config.LoadEnv()              // load env
	db := config.ConnectionDatabase(cfg) // database database

	// start the server
	server.Start(db, cfg)

}