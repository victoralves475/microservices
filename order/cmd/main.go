package main

import (
	"fmt"
	"github.com/victoralves475/microservices/order/config"
	"github.com/victoralves475/microservices/order/internal/adapters/db"
	"github.com/victoralves475/microservices/order/internal/adapters/grpc"
	"log"
	"os"

	//"github.com/victoralves475/microservices/order/internal/adapters/rest"
	"github.com/victoralves475/microservices/order/internal/application/core/api"
)

func main() {
	fmt.Printf("DEBUG: DATA_SOURCE_URL=%q\n", os.Getenv("DATA_SOURCE_URL"))

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf(" Failed to connect to database . Error : %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
