DB_URL=postgresql://root:secret@localhost:5432/carzorro?sslmode=disable
MYSQL_URL=mysql://root@secret@localhost:3306/carzorro?sslmode=disable

sqlcinstall:
	sudo snap install sqlc
mysql:
	docker run --name mysql -p 3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=carzorro -d mysql:latest
czmysql:
	docker run --name mysql -p 3306 -e MYSQL_ROOT_PASSWORD=secret -d czmysql:latest

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root carzorro

dropdb:
	docker exec -it postgres dropdb carzorro

sqlc:
	sqlc generate

migrate: 
	# to create db migration files
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/czarro/db/sqlc Store

test:
	go clean -testcache && go test -v -cover ./...
.PHONY: sqlcinstall postgres createdb dropdb migrate migrateup migratedown migrateup1 migratedown1 sqlc test mysql mysqlcreatedb mysqldropdb mock