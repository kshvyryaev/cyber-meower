# Run postgres docker

`docker run -d --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432 postgres`

# Run nats docker

`docker network create nats`
`docker run -d --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats`

# Create blank migration for meow service

`migrate create -ext sql -dir .internal/meow-service/migration -seq migration_name`

# Apply migration for meow service

migrate -path ./internal/meow-service/migration -database 'postgres://postgres:postgres@localhost:5432/cybermeowerdb?sslmode=disable' up
