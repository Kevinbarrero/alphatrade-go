postgres:
	sudo docker run --name alphatrade-go --network alphatrade-go -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4200 -d postgres
createdb:
	sudo docker exec -it alphatrade-go createdb --username=root --owner=root alphatrade-go
dropdb:
	sudo docker exec -it alphatrade-go dropdb alphatrade-go
migrateup:
	migrate -path db/migration -database "postgresql://root:4200@localhost:5432/alphatrade-go?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:4200@localhost:5432/alphatrade-go?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:4200@localhost:5432/alphatrade-go?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:4200@localhost:5432/alphatrade-go?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server
