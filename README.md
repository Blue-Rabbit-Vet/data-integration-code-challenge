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


## Setup/Explanation
Staying true to your request of a couple of hours, this is the result. I couldn't get Kafka working 100%.
I believe my issue was my web app was trying to connect to the kafka container before the kafka container
was ready, and my setTimeout/sleeps weren't having any effect on that behavior (hence the very high retry
rate of 100 on the connections, which still seemed to not work). In order to run the app to it's fullest
extent, run `docker compose build`, and once that is finished, you can run `docker compose up` and wait
for that to spin everything up. Once finished, you can go to localhost:8080 in a browser to find a very
simple (and not that pretty) interface where you can enter a name that will be submitted to a local
SQLite database (data.db). From there it was supposed to push to the Kafka producer, and then the consumer
was set to log it as the messages were received. But as previously mentioned, I could not get those
connections working properly in the time allotted. Thank you very much for the opportunity, and I hope
to hear from you soon.

Ryan Devenney
