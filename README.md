# Audio Media  Server

This repo uses a mixture of Uncle Bob's clean architecture and Domain Driven Design.


# How to run app

Fastest to get the apps runs would to run the `init.sh` bash script. Before doing that consult the dependencies section. 

The project uses go modules, so an appropriate version is needed, I recommend 1.12 and above.

If you want to build the manual way, use `go run main.go` at the root of the project, or to `go build` .
The app will run on port  `8080` or that can be easily changed using the `SERVICE_PORT` env variable like so `SERVICE_PORT=8000 go run main.go`.


# Consumer
The consumer is a slim application, that just listens for a subject and logs that even data to the filesystem.

# Fun stuff

This projects uses Nats which can be considered a messaging service, but more than anything it's a tool that enables distributed development.


# Directory structure explanation

- Domain

    - Had the domain objects currently very simple aggregate structure without much behavior

- Dto

    - Simple DTO helpers

- Interfaces

    - Contains the repository interface implementation,and also the the route handlers. This project uses `Chi`

- Usecases

    - Contains the application service and top level application orchestration logic

# Database 

This app uses `SQLite` for the metadata and the filepath storage of audio filles (currently just wav format)



# Dependencies

- SQLite 3: You will need to have this installed locally
- Nats Messaging Server: will be provided with docker-compose
- Golang: Needs to be installed locally, go version 1.12 and up 

# How to run tests

-  Unit tests

    - Best done the idiomatic way using `go test ./...` at the root
        

# Triggering the API Endpoint

```
curl --location --request POST 'localhost:8089/upload?name=ping.wav' \
--header 'Content-Type: audio/wave' \
--data-binary '{path to your wav file}'
```