run_services:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml up --bu -d

run_backend:
	docker-compose up --bu -d

test_database_run:
	docker pull postgres:11-alpine && docker run --name database_test -p 0.0.0.0:7010:5432 -e POSTGRES_USER=spkuser -e POSTGRES_PASSWORD=spkpassword -e POSTGRES_DB=spkdb -d postgres:11-alpine

test_database_container_remove:
	docker rm database_test

test_database_remove:
	docker kill database_test && docker rm database_test

test:
	go test ./tests/*.go -v

.PHONY: run_services run_backend test_database_run test_database_remove test
