package main

import (
	"log"
	"os"

	"github.com/Chandra5468/cfp-Products-Service/cmd/api"
	"github.com/Chandra5468/cfp-Products-Service/cmd/grpcserver"
	"github.com/Chandra5468/cfp-Products-Service/internal/config"
	"github.com/Chandra5468/cfp-Products-Service/internal/services/database/postgresql"
)

func main() {

	// Loading env configs
	err := config.MustLoad()

	if err != nil {
		log.Fatalf("Error while loading env file %v", err)
	}

	// Connect to psql database here
	db, err := postgresql.NewPostgres(os.Getenv("POSTGRESQL_STRING"))
	if err != nil {
		log.Fatalf("Unable to connect with postgresql database---- %v", err)
	}

	// Connect to Mongodb database here

	// Call GRPC Server here. No need to pass
	grpcServer := grpcserver.NewGrpcServer(":9002", db)
	go grpcServer.Run()
	// Calling HTTP API Server
	server := api.NewApiServer(os.Getenv("HTTP_ADDRESS"), db)

	server.RUN()
	// set APP_ENV=development
	//  go run .\cmd\main.go
}
