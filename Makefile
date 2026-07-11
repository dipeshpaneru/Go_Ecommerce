build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom


migration: 
	@migrate create -ext sql -dir cmd/migration/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migration/main.go up


migrate-down:
	@go run cmd/migrate/main.go down