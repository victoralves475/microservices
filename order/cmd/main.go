package main

import (
	"github.com/victoralves475/microservices/order/config"
	dbAdapter "github.com/victoralves475/microservices/order/internal/adapters/db"
	grpcAdapter "github.com/victoralves475/microservices/order/internal/adapters/grpc"
	payAdapter "github.com/victoralves475/microservices/order/internal/adapters/payment"
	api "github.com/victoralves475/microservices/order/internal/application/core/api"
	"log"
)

func main() {
	dbA, err := dbAdapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	pA, err := payAdapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Payment stub error: %v", err)
	}

	app := api.NewApplication(dbA, pA)
	grpcAdapter := grpcAdapter.NewAdapter(app, config.GetApplicationPort())
	grpcAdapter.Run()
}
