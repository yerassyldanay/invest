BINARY=rubi

POSTGRES_HOST=0.0.0.0
POSTGRES_PORT=7001
POSTGRES_PASSWORD=simple
POSTGRES_USER=simple
POSTGRES_DB_NAME=simple
POSTGRES_CONTAINER_NAME=invest_postgres
POSTGRES_VERSION=11-alpine

REDIS_HOST=0.0.0.0
REDIS_PORT=7002

services_run:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml up --bu -d

services_down:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml down

backend_run:
	docker-compose up --bu -d

postgres:
	docker pull postgres:${POSTGRES_VERSION} && docker run --name ${POSTGRES_CONTAINER_NAME} -p ${POSTGRES_HOST}:${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_DB=${POSTGRES_DB_NAME} -d postgres:${POSTGRES_VERSION}

postgres_remove:
	docker kill ${POSTGRES_CONTAINER_NAME} && docker rm ${POSTGRES_CONTAINER_NAME}

test:
	go test ./tests/*.go -v

migrate_up:
	migrate -path ./db/postgre/migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable -verbose up

migrate_down:
	migrate -path ./db/postgre/migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable -verbose down

env:
	echo $(POSTGRES_USER)

.PHONY: services_run services_down backend_run postgres postgres_remove test migrate_up migrate_down env
