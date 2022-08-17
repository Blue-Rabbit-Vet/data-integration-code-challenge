# Blue Rabbit Data Integration Code Challenge

Fork this repo and create an app using languages and frameworks of your choice that 
*literally* introduces you to us. Submit your response back to us here in the form of a pull 
request or submit it to us privately. Please don't spend more than a couple of hours on it. It's ok
if you don't finish, just tackle the requirements in order and take it as far as you can in the time frame.

Include A README with instructions on how to build/run the app. Use the README to let us know
why you chose the technologies you did. Notes on design patterns, challenges, or aspects
of your stack that you find interesting are also appreciated!

Provided is a `docker-compose-yml` file to help you start kafka. You are welcome to use other messaging services instead.

### Requirements
1. Create an API with an endpoint or operation that we can call and pass data to, save the request to a database. The shape of the data and the storage mechanism are up to you.
2. Create a sh script or add to README the commands to create topic/queue.
3. Publish API data to a topic/queue.
4. Add a consumer to your API to consume from the topic/queue and perform an operation of your choice with the message, .i.e. log to console, write to database, write to file.
5. Create a minimal frontend that calls your api.


####
- [X] Create an API with an endpoint or operation that we can call and pass data to, save the request to a database. The shape of the data and the storage mechanism are up to you.
- [X] Create a sh script or add to README the commands to create topic/queue.
- [X] Publish API data to a topic/queue.
- [X] Add a consumer to your API to consume from the topic/queue and perform an operation of your choice with the message, .i.e. log to console, write to database, write to file.
- [] Create a minimal frontend that calls your api.

## Setup

### Building locally

First Navigate your terminal of choice to the root of the project folder

Run the following command: `docker compose build`

### Running locally

After following the instructions to build the project complete you can then run the project with `docker compose up`

This command will bring up all the required containers for the project.

On first run the database will be empty, for ease of demoing (not recommended for production use) navigate to `localhost:5000/setup` in your browser. This will create the database schema for the web api. 

Next follow the steps in the Creating a topic section to create your desired topic

#### Use
The best way to currently interface with the api is using the docs page at `localhost:5000/docs/ui`. This will allow you to see the data structure for endpoints as well as send requests using the "Try now" button.

### Creating a topic

In a command line prompt in the project root folder running the following command with the name of the topic substituted in:

`docker-compose exec broker kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic <name of topic>`

### Topic consumer
The consumer is implemented in worker.py and is started in its own container. This container will consume starting on a 30 second offset. Currently it is only listening on a sample topic set in `worker.py`. It will output to standard out in the docker logs.