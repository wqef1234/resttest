.PHONY: run
run:
	go run main.go

.PHONY:build
build:
	go build main.go

.PHONY:exec
exec:
	./main

.PHONY: docker_build
docker_build:
	sudo docker build -t resttest .

.PHONY: docker_run
docker_run:
	sudo docker run -p 8080:8081 -it resttest

.PHONY: images
images:
	sudo docker images


DEFAULT_GOAL := run