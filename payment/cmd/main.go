package main

import (
	"log"

	"github.com/victoralves475/microservices/payment/config"
	dbAdapter "github.com/victoralves475/microservices/payment/internal/adapters/db"
	grpcServer "github.com/victoralves475/microservices/payment/internal/adapters/grpc"
	api "github.com/victoralves475/microservices/payment/internal/application/core/api"
)

func main() {
	dbA, err := dbAdapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	app := api.NewApplication(dbA)

	server := grpcServer.NewServer(app, config.GetApplicationPort())
	server.Start()
}
