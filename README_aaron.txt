This is the source code for a programming exercise for Blue Rabbit. 

- Commands:
    - to create topic (on windows, from kafka directory); below the topic is
      called recordsTopic

$ bin/windows/kafka-topics.bat --create --topic recordsTopic --bootstrap-server localhost:9092
  
    - to start express (in server directory)

$ yarn dev 

    - to run client locally (in client directory); in production apps this is 
      unnecessary due to webpack, but I avoided webpack for simplicity.

$ yarn start


- Technologies chosen
  Client
    - Command line: api_client.sh, a bash script using curl to call api. Run 
                    with no arguments to see usage.
    - Kafka consumer: kafka_consumer.js, node script to consume kafka events
    - Web/React/Bootstrap: I have the most experience with these client 
                           technologies. They provide a modern, responsive user 
                           experience.

  Server
    - Node/Express: I have the most experience with these server technologies. 
    - SQLite: Provides a test database with zero installation overhead. I have 
              included my test database file (database.db); you can start with
              an empty database by deleting database.db. I typically use Postgres
              and an ORM, but both seemed like overkill for this exercise. 
    - Kafka: This was the messaging service mentioned in the exercise and seemed
             like a good choice. 

-  Challenges: I used a stripped down set of npm packages, usually I start with
               a large boilerplate setup that has everything I need. I ran into 
               weird, confusing errors which turned out to be due to missing 
               dependencies. 

- Future Directions: My next step would be to create a backend service that
                     consumes/listens for new kafka events, and provides 
                     notifications. Example: an invoice or receipt is generated
                     and the relevant users are notified. 
