FROM alpine:3.7

WORKDIR /app

COPY ./dist/improve-your-vocabulary/api ./
COPY ./config.json ./

# timezone
RUN apk update && apk --no-cache add tzdata 

EXPOSE 8080

ENTRYPOINT ./api
