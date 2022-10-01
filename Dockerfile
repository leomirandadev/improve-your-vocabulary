FROM alpine:3.7

COPY ./dist/improve-your-vocabulary/main /app/improve-your-vocabulary/main
COPY ./cmd/config.json /app/improve-your-vocabulary/main

WORKDIR /app

# timezone
RUN apk update && apk --no-cache add tzdata 

EXPOSE 8080

ENTRYPOINT ./improve-your-vocabulary/main
