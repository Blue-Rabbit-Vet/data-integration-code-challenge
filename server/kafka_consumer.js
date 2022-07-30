const kafka = require("kafka-node");

(Consumer = kafka.Consumer),
  (client = new kafka.KafkaClient()),
  (consumer = new Consumer(
    client,
    [
      { topic: "recordsTopic", partition: 0 },
    ],
    {
      autoCommit: false,
    }
  ));

  consumer.on('message', function (message) {
    console.log(`Topic: ${message.topic} \n Value: ${message.value} \n Offset: ${message.offset}\n\n`);
});