# ITU DevOps course, Spring 2024

## Run On Local Machine
Make sure you have a `.env` file in the root folder, and that it contains the following fields:
```
# PostgreSQL database configuration
DB_HOST=localhost
DB_PORT=5433
DB_DATABASE=minitwit
DB_USER=root
DB_PASS=password
```
Then run the following commands:
```
make postgresinit
make createdb
make run
```

## Docker
Make sure you have a `.env` file in the root folder, and that it contains the following fields:
```
# PostgreSQL database configuration
DB_HOST=db
DB_PORT=5432
DB_DATABASE=minitwit
DB_USER=root
DB_PASS=password
```
Start service:
```
docker-compose build
docker-compose up -d
```

Stop service:
```
docker-compose down
```

Stop service and remove container:
```
docker-compose down --volumes
```

### Why Docker is not working?
Seems like the `docker-compose` is not able to connect to the database.
