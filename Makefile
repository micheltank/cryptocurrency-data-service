deps:
	go mod tidy

run:
	set -o allexport && source env.example && go run main.go

first_run: deps run

lint:
	golangci-lint run

docker_build:
	docker build -t micheltank/cryptocurrency-data-service .

test:
	go test -v -cover ./...

swagger:
	swag init -g cmd/rest/api.go

docker_compose:
	docker-compose up -d

all: deps lint docker_build test swagger docker_compose