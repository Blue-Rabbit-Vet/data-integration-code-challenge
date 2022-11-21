# Running Application
The docker-compose file that was supplied has been updated to run all parts of the requirements. I have added mongodb for my database, playerapp for my java backend, and react-ui for the frontend.

Navigate to root directory and run,
docker compose up

You will also find the a sh file to create a topic with the name of 'players' in root/src/main/resource/config

Once everything is up and running you can navigate to http://localhost:3000/

Here you will find a bare bones UI to take a first name, last name, and a number. 
This UI will submit the value to be created in the mongodb.

The rest api that the UI submits the players to is found here, http://localhost:8080.

POST - creates a player
/addplayer

Hitting the Endpoint 
curl --location --request PUT 'http://localhost:8080/addplayer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Robert",
    "lastName": "Wais",
    "number": 43
}'



PUT
/updateplayer - updates a players number

Hitting the Endpoint 
curl --location --request PUT 'http://localhost:8080/updateplayer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Robert",
    "lastName": "Wais",
    "number": 43
}'

## Technologies
I chose the technologies that I did because I was most familiar with them and I thought that they defined the work I have done most professionally.
I chose mongo because I thought it would be something fun to incorporate in the project that didn't add too much additional overhead.

### Backend
For the java side. I created a spring boot app that has both a Rest Controller to take calls from the frontend as well as a KafkaListener to listen to messages from the topic.

The Rest Controller calls a player service which then does most of the work. I create a key from the first name and last name. I use this value to make sure I don't create a similar player in the db. This was a simple way of using it with the mongodb but this is somewhere where I could definitely have a more robust solution. I could create a script to create a unique index on the key value to also have database validation. The uniqueness in this case is based off of the first name and last name. Having an email or username in the future would be an improvement as well.

The KafkaListener consumes the message and logs it to the console. Some improvements could be adding manual acknowledgement to my listener so that I am not depending on the container to commit the offset. Another improvement would be updating the container factory to have a few more custom values to align with the use of my listener such as a retry policy and/or a deadletter strategy.


### Frontend
For the front end I wrote a bare bones React app to submit to the api above. My experiences with react are extremely minimal but I thought it would be a good chance to try something a bit new. I added a few fields and a button to be able to submit to the backend api. A lot of room for improvements here. Some of them are basic styling, a responsive frontend on success or after submission, and/or the ability to update through the frontend.





# Blue Rabbit Data Integration Code Challenge

Fork this repo and create an app using languages and frameworks of your choice that 
*literally* introduces you to us. Submit your com.integration.response back to us here in the form of a pull 
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
