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
