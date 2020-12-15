run_services:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml up --bu -d

down_services:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml down

run_backend:
	docker-compose up --bu -d

test_database_run:
	docker pull postgres:11-alpine && docker run --name database_test -p ${POSTGRES_HOST}:${POSTGRES_PORT}:5432 -d postgres:11-alpine

test_database_container_remove:
	docker rm database_test

test_database_remove:
	docker kill database_test && docker rm database_test

test:
	go test ./tests/*.go -v

migrate_up:
	migrate -path ./db/postgre/migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -verbose up

migrate_down:
	migrate -path ./db/postgre/migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -verbose down

env_variables:
	echo $(POSTGRES_USER)

.PHONY: run_services run_backend test_database_run test_database_remove test migrate_up migrate_down

