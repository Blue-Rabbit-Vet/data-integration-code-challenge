const { startConsumer } = require("./consumer");
const express = require("express");
const { apiController, serveHome } = require("./controller");

const app = express();
app.use(express.json());

app.post("/", apiController);
app.get("/", serveHome);

app.listen(8080, async () => {
	console.log(`Listening on ${8080}`);
	await startConsumer();
});
