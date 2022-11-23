postgres:
	docker run --name scalefocusdb -p 5432:5432  -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=i76248591M -d postgres:latest

createdb:
	docker exec -it scalefocusdb createdb --username=root --owner=root ships_data

migratecreate:
	migrate create -ext sql -dir api/db/migration -seq init_schema

migrateup:
	 migrate -path db/migrations -database "postgresql://root:i76248591M@localhost:5432/ships_data?sslmode=disable"-verbose up

dropdb:
	docker exec -it scalefocusdb dropdb ships_data

migratedown:
	migrate -path db/migrations -database "postgresql://root:i76248591M@localhost:5432/ships_data?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate