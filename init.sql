CREATE TABLE files
(
    id         INTEGER PRIMARY KEY,
    filename   VARCHAR(500),
    filepath   VARCHAR(1000),
    duration   INTEGER,
    sampleRate INTEGER,
    numChans   INTEGER,
    bitDepth   INTEGER
)

