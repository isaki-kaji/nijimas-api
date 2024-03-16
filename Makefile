DB_URL=postgresql://nijimas:password@localhost:5432/nijimas?sslmode=disable

postgres:
	docker run --name nijimas-postgres -p 5432:5432 -e POSTGRES_USER=nijimas -e POSTGRES_PASSWORD=password -d postgis/postgis:16-3.4 

createdb:
	docker exec -it nijimas-postgres createdb --username=root nijimas

dropdb:
	docker exec -it nijimas-postgres dropdb --username=root nijimas

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate


.PHONY: db_docs db_schema createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc postgres