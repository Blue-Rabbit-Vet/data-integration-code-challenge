const sqlite3 = require("sqlite3");
const { open } = require("sqlite");

const openDb = async () => {
	const db = await open({
		filename: "./data.db",
		driver: sqlite3.Database,
	});

	await migrations(db);

	return db;
};

// "migrations" to make sure the database is setup and ready
const migrations = async (db) => {
	await db.exec(
		`CREATE TABLE IF NOT EXISTS info
		(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT(255), date_in_minutes BIGINT);`
	);
};

module.exports = {
	openDb,
};
