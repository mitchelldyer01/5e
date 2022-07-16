tidy:
	go mod tidy

test:
	go test ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -o out/5e ./cmd

up: clean
	docker-compose up

clean:
	docker rmi 5e_5e -f || true