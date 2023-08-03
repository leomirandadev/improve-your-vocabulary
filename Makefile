
.PHONY: docs

NAME = "improve-your-vocabulary"
DB_CONNECTION = "root:root@(127.0.0.1:3306)/improve_your_vocabulary?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s&parseTime=true"

docs:
	@swag init -g cmd/api/main.go

install: 
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
	@GOOS=linux go build -o dist/$(NAME)/api cmd/api/*.go 
	@echo $(NAME): Construindo a imagem
	@docker build -t $(NAME) .

docker-up: 
	@docker compose -f "docker/docker-compose.yml" up -d --build

docker-down: 
	@docker compose -f "docker/docker-compose.yml" down

local-up: 
	@docker compose -f "docker/db_dev/docker-compose.yml" up -d --build
	@docker compose -f "docker/memcached_dev/docker-compose.yml" up -d --build
	@docker compose -f "docker/tracer_dev/docker-compose.yml" up -d --build

local-down: 
	@docker compose -f "docker/db_dev/docker-compose.yml" down
	@docker compose -f "docker/memcached_dev/docker-compose.yml" down
	@docker compose -f "docker/tracer_dev/docker-compose.yml" down

run: local-up
	@go run cmd/api/*.go

run-watch: local-up
	@cd cmd && nodemon --exec go run cmd/api/*.go --signal SIGTERM

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
	@go test -v -coverpkg=./... -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out

test-cover: test
	@go tool cover -html=coverage.out 

open-swagger:
	open http://localhost:8080/swagger/index.html

open-jaeger:
	open http://localhost:16686/search