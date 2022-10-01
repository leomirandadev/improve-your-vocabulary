
.PHONY: docs

NAME = "improve-your-vocabulary"
DB_CONNECTION = "root:root@(127.0.0.1:3306)/improve_your_vocabulary?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s&parseTime=true"

docs:
	@swag init -g cmd/main.go

setup: 
	@echo "installing nodemon..."
	@npm install -g nodemon
	@echo "installing goose..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "installing swaggo..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "downloading project dependencies..."
	@go mod tidy

build: 
	@echo $(NAME): Compilando o micro-serviço
	@go build -o dist/$(NAME)/main cmd/main.go 
	@echo $(NAME): Construindo a imagem
	@docker build -t $(NAME) .

docker-up: 
	@docker compose -f "docker/docker-compose.yml" up -d --build

docker-down: 
	@docker compose -f "docker/docker-compose.yml" down

up-local: 
	@docker compose -f "docker/db_dev/docker-compose.yml" up -d --build
	@docker compose -f "docker/memcached_dev/docker-compose.yml" up -d --build
	@docker compose -f "docker/tracer_dev/docker-compose.yml" up -d --build

down-local: 
	@docker compose -f "docker/db_dev/docker-compose.yml" down
	@docker compose -f "docker/memcached_dev/docker-compose.yml" down
	@docker compose -f "docker/tracer_dev/docker-compose.yml" down

run: up-local
	@go run cmd/main.go

run-watch: up-local
	@nodemon --exec go run cmd/main.go --signal SIGTERM

mig-create: 
	@goose -dir ./migrations create $(MIG_NAME) sql 

mig-status: 
	@goose mysql $(DB_CONNECTION) status

mig-up: 
	@goose -dir ./migrations mysql $(DB_CONNECTION) up

mig-down: 
	@goose -dir ./migrations mysql $(DB_CONNECTION) down

mock: 
	@go generate ./...

test:
	@go test -v -p 1 -cover -failfast ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

test-cover: test
	@go tool cover -html=coverage.out 
