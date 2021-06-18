-include ./environment/local.env

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

compile:
	go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

postgres:
	docker pull postgres:${POSTGRES_VERSION} && docker run --name ${POSTGRES_CONTAINER_NAME} -p ${POSTGRES_HOST}:${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_USER=${POSTGRES_USER} -d postgres:${POSTGRES_VERSION}

postgres_delete:
	docker stop ${POSTGRES_CONTAINER_NAME} && docker rm ${POSTGRES_CONTAINER_NAME}

postgres_logs:
	docker logs ${POSTGRES_CONTAINER_NAME} -f --tail=30

redis:
	docker pull redis && docker run --name spk_redis -p ${REDIS_HOST}:${REDIS_PORT}:6379 -d redis

redis_delete:
	docker kill spk_redis && docker rm spk_redis

redis_logs:
	docker logs spk_redis -f --tail=30

migrate_up:
	migrate -path ${MIGRATION_PATH} -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -verbose up

migrate_down:
	migrate -path ${MIGRATION_PATH} -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -verbose down

generate:
	sqlc generate

test:
	go test -v ./...

server:
	go run main.go

sql_next_migrate:
	migrate create -ext sql -dir ./database/postgres/ -seq -digits 5 nameit

services_:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml up --remove-orphans --bu -d

services_log:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml logs -f --tail=30

services_stop:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml stop

services_rm:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml rm

images_all_rm:
	docker rmi $(docker images -aq)

images_postgres_rm:
	docker rmi postgres:11

#back:
#	env HOST=0.0.0.0 docker-compose up --bu

backend:
	env HOST=0.0.0.0 docker-compose up --bu -d

backend_log:
	env HOST=0.0.0.0 docker-compose logs -f --tail=30

backend_stop:
	env HOST=0.0.0.0 docker-compose stop

backend_rm:
	env HOST=0.0.0.0 docker-compose rm

volume:
	docker volume prune

.PHONY: postgres postgres_delete postgres_logs postgres_up postgres_down generate test compile clean lint-prepare lint server sql_next_migrate
.PHONY: services_ services_stop services_rm backend backend_stop backend_rm services_log
.PHONY: images_all_rm images_postgres_rm volume
