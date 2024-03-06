postgresinit:
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine
postgrestestinit:
	docker run --name postgres-test -p 5433:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -d postgres:16-alpine
postgres:
	docker exec -it postgres16 psql

createdbtest:
	docker exec -it postgres-test createdb --username=test --owner=test minitwit-test
dropdbtest:
	docker exec -it postgres16 dropdb minitwit-test
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root minitwit

dropdb:
	docker exec -it postgres16 dropdb minitwit

migrateup:
	migrate -path src/database/migrations -database "postgresql://root:password@localhost:5433/minitwit?sslmode=disable" -verbose up

migratedown:
	migrate -path src/database/migrations -database "postgresql://root:password@localhost:5433/minitwit?sslmode=disable" -verbose down

run:
	go run src/main.go .

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown