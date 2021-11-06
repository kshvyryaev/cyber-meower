# Meower service

Service for creaing meow messages

## Run postgres docker

`docker run -d --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432 postgres`

## Run nats docker

`docker network create nats`
`docker run -d --name nats --network nats -p 4222:4222 -p 8222:8222 nats`

## Create blank migration

`migrate create -ext sql -dir ./pkg/migration -seq migration_name`

## Apply migrations

`migrate -path ./pkg/migration -database 'postgres://postgres:postgres@localhost:5432/cybermeowerdb?sslmode=disable' up`
