DB_URL=postgresql://root:password@localhost:5432/nijimas?sslmode=disable

postgres:
	docker run --name nijimas-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16.2-alpine3.19 

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

server:
	go run main.go

test:
	go test -v -cover -short ./...

mockdb:
	mockgen -package mockdb  -destination db/mock/repository.go  github.com/isaki-kaji/nijimas-api/db/sqlc Repository

mockuserservice:
	mockgen -package mockservice  -destination service/mock/user_service.go  github.com/isaki-kaji/nijimas-api/service UserService

mockpostservice:
	mockgen -package mockservice  -destination service/mock/post_service.go  github.com/isaki-kaji/nijimas-api/service PostService


.PHONY: db_docs db_schema createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc postgres server test mockdb mockuserservice mockpostservice