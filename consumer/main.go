package main

import (
	"audioTest/helpers"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// Connect to a server
	nats_port := helpers.GetEnvVar("NATS_PORT", false)
	if nats_port == "" {
		nats_port = nats.DefaultURL
	}
	nc, err := nats.Connect(nats_port)
	if err != nil {
		log.Fatal("Unable to connect to the nats server")
	}

	// Log file
	absPath, _ := filepath.Abs("./consumer/log.txt")
	log.Printf("Absolute Path to log file: %s", absPath)
	logFile, err := os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// wait for up to 10 messages
	wg := sync.WaitGroup{}
	wg.Add(1000)
	// Listen for upload subject
	_, err = nc.Subscribe("file.upload", func(m *nats.Msg) {
		var msg = string(m.Data)
		if _, err = logFile.WriteString(fmt.Sprintf("\n%s", msg)); err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}
