# Introduction

This is the source code for a programming exercise by Aaron Mitchell for Blue Rabbit. 

## Commands:

To run kafka, follow the instructions [here](https://kafka.apache.org/quickstart).

To create the topic (from kafka directory), see below. Here the topic is called `recordsTopic`

```
$ bin/kafka-topics.sh --create --topic recordsTopic --bootstrap-server localhost:9092
```
  
To install the necessary npm packages, run the following:

```
$ cd client
$ yarn install
$ cd ../server
$ yarn install
```

To start the Node express server, run the following.

```
$ cd ../server
$ yarn dev 
```

To run the React client locally, see below. In production apps this is unnecessary due to webpack. I avoided webpack in this exercise for simplicity.

```
$ cd ../client
$ yarn start
```

To run the kafka consumer command-line app, the following should be executed in the server directory.

```
$ node kafka_consumer.js 
```

To run the command-line Node api client app, do something like the following (from the top-level directory).

```
$ chmod 755 api_client.sh
$ ./api_client.sh Fluffy dan@email.com "annual+checkup" \$100 true" 
```

## Technologies chosen

### Client

* Command line: api_client.sh, a bash script using curl to call api. Run with no arguments to see usage.
* Kafka consumer: kafka_consumer.js, node script to consume kafka events
* Web/React/Bootstrap: I have the most experience with these client technologies. They provide a modern, responsive user experience.

### Server
* Node/Express: I have the most experience with these server technologies. 
* SQLite: Provides a test database with zero installation overhead. I have included my test database file (database.db); you can start with an empty database by deleting database.db. I typically use Postgres and an ORM, but both seemed like overkill for this exercise. 
* Kafka: This was the messaging service mentioned in the exercise and seemed like a good choice. 

## Challenges

I used a stripped down set of npm packages, usually I start with a large boilerplate setup that has everything I need. I ran into weird, confusing errors which turned out to be due to missing dependencies. 

## Future Directions

My next step would be to create a backend service that consumes/listens for new kafka events, and provides notifications. Example: an invoice or receipt is generated and the relevant users are notified. 