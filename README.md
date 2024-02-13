# ITU DevOps course, Spring 2024

## NB!
Make sure you have a `.env` file in the root folder, and that it contains the following fields:
```
# PostgreSQL database configuration
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USER=
DB_PASS=
```

## Run On Local Machine
```
make postgresinit
make createdb
make run
```

## Docker (NOT WORKING)
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