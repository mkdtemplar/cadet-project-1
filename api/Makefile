postgres:
	docker run --name scalefocusdb -p 5432:5432  -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=i76248591M -d postgres:latest

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

createdb:
	docker exec -it scalefocusdb createdb --username=root --owner=root ships_data

createtestdb:
	docker exec -it scalefocusdb createdb --username=root --owner=root ships_data_test

dropdb:
	docker exec -it scalefocusdb dropdb ships_data

migrateup:
	migrate -path db/migrations -database "postgresql://root:i76248591M@localhost:5432/ships_data?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:i76248591M@localhost:5432/ships_data?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate