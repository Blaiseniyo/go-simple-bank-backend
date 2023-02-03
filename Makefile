postgres:
	docker run --name postegres15 -p 5433:5432 -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose down

.PHONY:postgres createdb dropdb