postgres:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

createmigration:
	migrate create -ext sql -dir db/migrations -seq add_users
tests:
	go test -v -cover ./...

server:
	go run main.go

.PHONY:postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 tests server