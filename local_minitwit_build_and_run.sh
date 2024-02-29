#!/bin/bash

# Step 1: Create the Docker network (if it doesn't already exist)
docker network ls | grep minitwit-network >/dev/null || docker network create minitwit-network

# Step 2: Build and run the PostgreSQL database container
docker build -t minitwit-postgres -f database/Dockerfile .
docker run --network=minitwit-network -d --name minitwit-postgres-instance -p 5432:5432 -v $(pwd)/database/init:/docker-entrypoint-initdb.d minitwit-postgres

# Optional: Wait for PostgreSQL to be fully up before proceeding
# This requires the PostgreSQL client installed or another way to check the DB status
echo "Waiting for PostgreSQL to start..."
sleep 10 # Simple sleep, replace with a loop that checks actual DB readiness if needed

# Step 3: Build and run the web app container
docker build -t minitwit-app .
docker run --network=minitwit-network -d --name minitwit-app-instance -p 8080:8080 minitwit-app

# Step 4: Build and run the API container
docker build -t minitwit-api -f api/Dockerfile .
docker run --network=minitwit-network -d --name minitwit-api-instance -p 5000:5000 minitwit-api

echo "All containers are up and running."
