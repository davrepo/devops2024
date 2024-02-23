#!/bin/zsh

# Stop script execution on any error
set -e

# Drop the database
make dropdb

# Create the database
make createdb

# Change directory to 'api'
cd api

# Run the Go API
go run api.go
