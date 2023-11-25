
PROJECT := infinity

build:
	go build

test:
	go test ./...

#docker commands
docker_login:
	docker login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD}

docker_ps:
	docker ps

start_db:
	docker-compose	-f ./infra/docker-compose.db.yml	--project-name	$(PROJECT) up -d

stop_db:
	docker-compose	-f ./infra/docker-compose.db.yml	--project-name	$(PROJECT) down -v

