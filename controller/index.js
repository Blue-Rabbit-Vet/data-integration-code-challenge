const { openDb } = require("../util");
const path = require("path");

const { Kafka } = require("kafkajs");

const kafka = new Kafka({
	clientId: "data-int-app",
	brokers: ["localhost:9092"],
	retry: { retries: 100 },
});

const producer = kafka.producer({ retry: { retries: 100 } });

const apiController = async (req, res) => {
	try {
		const { body } = req;

		if (!body.name) {
			throw "Expected 'name' information.";
		}

		const db = await openDb();

		const dateInMinutes = new Date().getTime() / 60_000;

		await db.run(
			"INSERT INTO info (name, date_in_minutes) VALUES (?, ?)",
			body.name,
			dateInMinutes
		);

		await producer.connect();
		producer.send({
			topic: "names-topic",
			messages: [
				{
					name: body.name,
					dateInMinutes,
				},
			],
		});

		res.json({ message: "Successfuly entered information." });
	} catch (error) {
		res.json({ error });
	}
};

const serveHome = async (_req, res) => {
	res.sendFile(path.join(__dirname, "../ui/index.html"));
	return;
};

module.exports = {
	apiController,
	serveHome,
};
