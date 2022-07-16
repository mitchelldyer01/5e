tidy:
	go mod tidy

test:
	go test ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -o out/characters-5e ./cmd

up: clean
	docker-compose up

clean:
	docker rmi characters-5e_characters-5e -f