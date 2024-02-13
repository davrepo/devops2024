# ITU DevOps course, Spring 2024

## NB!
Make sure you have a `.env` file in the root folder

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