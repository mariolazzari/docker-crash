# Docker Crash Course for Beginners

[Github](https://github.com/leon-devs/course-docker_intro/)

## Introduction

### What is Docker

- Docker CLI
- Docker API
- Docker Descktop
- Docker deamon

### Installing Docker

- Engine (https://docs.docker.com/engine/)
- Desktop (https://docs.docker.com/desktop/)

## Docker in action

### First Dockerfile

```dockerfile
# debian based image
FROM golang:1.24 AS builder

# mkdir /app && cd /app
WORKDIR /app

# cp from to
COPY . .

# exec command in current env
RUN go mod download

# go build
RUN go build -o application cmd/app/main.go

# app entrypoint: 1st command executed by docker
ENTRYPOINT ["./application"]

# command params
CMD [""]
```

```sh
docker images
# exec container and and remove it aftef execution
docker run -rm docker-basics
docker ps
# bind port
docker run -rm -p 1234:8080 docker-basics
docke ps -a
docker start #id
docker stop #id
```
