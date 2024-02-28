#!/bin/bash

# Stop the containers
echo "Stopping containers..."
docker stop minitwit-postgres-instance minitwit-app-instance minitwit-api-instance

# Remove the containers
echo "Removing containers..."
docker rm minitwit-postgres-instance minitwit-app-instance minitwit-api-instance

# Remove the images
echo "Removing images..."
docker rmi minitwit-postgres minitwit-app minitwit-api

echo "Cleanup complete."
