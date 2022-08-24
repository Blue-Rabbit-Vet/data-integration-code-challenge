const { Kafka } = require("kafkajs");

const startConsumer = async () => {
	try {
		// Allows the containers to be fully ready before connection is attempted
		setTimeout(() => { }, 30_000);

		const kafka = new Kafka({
			groupId: "names-group",
			clientId: "data-int-app",
			brokers: ["localhost:9092"],
			retry: {
				retries: 100,
			},
		});

		const consumer = kafka.consumer({
			groupId: "names-group",
			retry: { retries: 100 },
		});
		await consumer.connect();
		await consumer.subscribe({ topic: "names-topic" });

		await consumer.run({
			eachMessage: async ({ message }) => {
				console.log(message);
			},
		});
	} catch (error) {
		console.log(error);
	}
};

module.exports = {
	startConsumer,
};
