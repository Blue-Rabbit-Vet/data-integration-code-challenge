const express = require("express");
const bodyParser = require("body-parser");
const sqlite3 = require("sqlite3").verbose();
const { Kafka } = require("kafkajs");
const cliendId = "my-app";
const brokers = ["localhost:9092"];
const topic = "recordsTopic";

const app = express();
const port = 5000;
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);
app.use(bodyParser.json());

const db = new sqlite3.Database(
  "./database.db",
  sqlite3.OPEN_READWRITE,
  (error) => {
    if (error) return console.error(error.message);
  }
);

let createTable = `CREATE TABLE if not exists petRecords(id INTEGER PRIMARY KEY, petName, ownerEmail, procedure, cost, isPaid)`;
db.run(createTable);

const kafka = new Kafka({ cliendId, brokers });
const producer = kafka.producer({});

app.post("/api", async (req, res) => {
  //persist data to sqlite
  try {
    const { petName, ownerEmail, procedure, cost, isPaid } = req.body;
    let insert = `INSERT INTO petRecords(petName, ownerEmail, procedure, cost, isPaid) VALUES (?,?,?,?,?)`;
    db.run(insert, [petName, ownerEmail, procedure, cost, isPaid], (error) => {
      if (error) return console.error(error.message);
    });
    //send data to topic 
    try {
      await producer.connect();
      await producer.send({
        topic: topic,
        messages: [{
          value: `${ownerEmail}'s pet, ${petName}, had a ${procedure} that cost ${cost} - Has Paid: ${isPaid}`
        }]
      });
      console.log("Successful append to topic")
    } catch (error) {
      console.log("Error: ", error);
    }
    return res.status(201).json({ record: req.body });
  } catch (error) {
    console.log("Error: ", error);
    return res.status(422);
  }
});

app.get("/api", async (req, res) => {
  let queryTable = `SELECT * FROM petRecords`;
  db.all(queryTable, [], (error, rows) => {
    if (error) return console.error(error.message);
    return res.status(200).json({ records: rows });
  });
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
