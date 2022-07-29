package main

import (
	"audioTest/helpers"
	"audioTest/src/application"
	"audioTest/src/infrastructure/db"
	"audioTest/src/infrastructure/filesystem"
	interfaces "audioTest/src/interfaces/repositories"
	"audioTest/src/interfaces/routes"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"os"
)

func main() {
	//Setup sqllite
	workingDir, err := os.Getwd()
	if err != nil {
		// Fail fast
		log.Fatal(err)
	}
	// Connect to a server
	nats_port := helpers.GetEnvVar("NATS_PORT", false)
	if nats_port == "" {
		nats_port = nats.DefaultURL
	}
	nc, err := nats.Connect(nats_port)
	if err != nil {
		log.Fatal("Unable to connect to the nats server")
	}

	// Dependency management
	sqlStorePath := fmt.Sprintf("%s/%s", workingDir, "dbstore/audio.sqlite")
	sqlHandler := db.NewSqliteHandler(sqlStorePath)

	audioRepo := interfaces.AudioRepositoryFactory(sqlHandler, filesystem.FileSystemFactory(), nc)
	audioService := application.AudioServiceFactory(audioRepo)

	router := routes.Routes(
		audioService,
	)

	// Start the http server
	servicePort := helpers.GetEnvVar("SERVICE_PORT", false)
	if servicePort == "" {
		servicePort = "8080"
	}
	err = http.ListenAndServe(":"+servicePort, router)
	if err != nil {
		log.Fatal(err)
	}

}
