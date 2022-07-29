#!/bin/bash

NAME=improve-your-vocabulary/$1
VERSION=$2

echo $NAME: Compilando o micro-serviço $NAME
# env GOOS=linux GOARCH=amd64 go build -o dist/$NAME
go build -o dist/$NAME

echo $NAME: Escrevendo o Dockerfile
CAT <<EOF > Dockerfile
    FROM alpine:3.7

    COPY ./dist/$NAME /opt/$NAME

    WORKDIR /opt
    EXPOSE 8080

    ENTRYPOINT ./$NAME
EOF

echo $NAME: Construindo a imagem
docker build -t $NAME .

echo $NAME: Removendo artefatos
rm ./Dockerfile
rm -rf ./dist
