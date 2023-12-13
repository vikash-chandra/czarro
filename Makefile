DB_URL=postgresql://root:secret@localhost:5432/carzorro?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root carzorro
dropdb:
	docker exec -it postgres dropdb carzorro
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1


.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1