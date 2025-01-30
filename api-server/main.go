package main

import (
	"database/sql"
	"gateway-router/api"
	"gateway-router/utils"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver            = "postgres"
// 	dbSource            = "postgresql://postgres:admin@localhost:5432/gateway_router?sslmode=disable"
// 	advertismentAddress = "0.0.0.0:5000"
// )

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatalf("unable to load the config, %e", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Panic("unable to open the database connection")
	}
	server := api.New(conn)
	go server.ListenForResponses()
	server.Start(config.ListenAddress)
	// sqlc.InitializeInterfaces()
	// controllers.CreateVlanInterface()
}
