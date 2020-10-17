build:
	go build -o bin

run:
	go run main.go

docker-build:
	docker build -t jstone28/mars-viking:latest .

docker-push:
	docker push jstone28/mars-viking:latest

publish: docker-build docker-push
