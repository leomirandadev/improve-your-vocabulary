
VERSION = $(shell git branch --show-current)
NAME = $(shell echo $(CURRENTNAME) | sed 's/\(.\)\([A-Z]\)/\1-\2/g' | tr '[:upper:]' '[:lower:]')

DB_CONNECTION = "root:root@(127.0.0.1:3306)/improve_your_vocabulary?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s&parseTime=true"

build: 
	./docker/build-image.sh $(NAME) $(VERSION) $(DB_CONNECTION) $(CACHE_URL) $(CACHE_EXP)

run:
	@go run main.go

run-watch:
	@nodemon --exec go run main.go --signal SIGTERM

mig-create: 
	@goose -dir ./migrations create $(MIG_NAME) sql 

mig-status: 
	@goose mysql $(DB_CONNECTION) status

mig-up: 
	@goose -dir ./migrations mysql $(DB_CONNECTION) up

mig-down: 
	@goose -dir ./migrations mysql $(DB_CONNECTION) down

generate-mock: 
	go generate ./...

test:
	go test -v -p 1 -cover -failfast ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

test-cover: test
	go tool cover -html=coverage.out 
