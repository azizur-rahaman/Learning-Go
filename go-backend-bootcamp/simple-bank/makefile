mgup:
	migrate -path db/migration -database "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable" --verbose up

mgdown:
	migrate -path db/migration -database "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: mgup mgdown sqlc test server