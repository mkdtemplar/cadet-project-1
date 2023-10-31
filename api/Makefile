include pkg/config/app.env

postgres:
	docker run --name scalefocusdb -p $(DB_PORT):$(DB_PORT)  -e POSTGRES_USER=$(DB_USER)  -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:latest

createdb:
	docker exec -it scalefocusdb createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

migratecreate:
	migrate create -ext sql -dir api/db/migration -seq init_schema

migrateup:
	 migrate -path pkg/db/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

dropdb:
	docker exec -it scalefocusdb dropdb $(DB_NAME)

migratedown:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate