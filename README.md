# Run postgres docker

`docker run -d --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432 postgres`

# Create blank migration for meow service

`migrate create -ext sql -dir .internal/meow-service/migration -seq migration_name`

# Apply migration for meow service

migrate -path ./internal/meow-service/migration -database 'postgres://postgres:postgres@localhost:5432/cybermeowerdb?sslmode=disable' up
