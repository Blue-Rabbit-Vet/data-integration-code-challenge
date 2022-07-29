#!/bin/bash

# Initialize database
echo "Initializing database..."
sqlite3 ./dbstore/audio.sqlite3 < init.sql

echo "Running docker-compose"
docker-compose -f docker-compose.yml up -d

echo "Building pub sub components"
go build -o consumer ./consumer
go build -o publisher

echo "Running pub service"
./publisher &

echo "Running consumer"
./consumer/consumer &