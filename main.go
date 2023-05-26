package main

import (
	"api/config"
	"api/database"
	"api/di"
	"api/graphqlServer"
	"api/managedTalent"
	"log"
)

func main() {
	// Initialize dig
	di.InitializeContainer([]di.RegisterFunction{
		config.Register,
		database.Register,
		managedTalent.Register,
	})

	// Run graphql server
	if err := graphqlServer.RunServer([]graphqlServer.InitializerFunc{
		managedTalent.InitializeQuery,
	}); err != nil {
		log.Fatalf("failed running graphql server: %v", err)
	}
}
