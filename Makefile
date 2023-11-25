
build:
	go build
test:
	go test ./...

#docker commands
docker_login:
	docker login -u jamvicky3094424 -p Bullshit@123

docker_ps:
	docker ps