.PHONY: run, migration_create, migration_up, migration_down, migration_force

run:
	go run cmd/main.go

migration_create:
	migrate create -ext sql -dir ./internal/migrations/ -seq users

migration_up:
	migrate -path ./internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/jwt_simple_echo?sslmode=disable" -verbose up	

migration_down: 
	migrate -path ./internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/jwt_simple_echo?sslmode=disable" -verbose down
