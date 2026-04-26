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
docker logs #id
```

### Build context

```sh
docker build -t docker-basics .
```

#### .dockeringore

[Docs](https://docs.docker.com/build/concepts/context/#dockerignore-files)

```.dockerignore
*.DS_Store

application

.gitignore
```

### Nultistage builds

```dockerfile
# debian based image
FROM golang:1.24 AS builder

# mkdir /app && cd /app
WORKDIR /app

# cp from to
COPY . .

# exec command in current env
RUN go mod download

# buiid image
FROM GOOS=linux GOARCH=amd64 CGO_ENABLED=0 golang:1.24 AS builder

# mkdir /app && cd /app
WORKDIR /app

# cp from to
COPY . .

# exec command in current env
RUN go mod download

# go build
RUN go build -o application cmd/app/main.go

# deploy image
FROM alpine AS serve

# cp from builder to serve
COPY --from=builder /app/application .


ENV HTTP_SERVER_PORT=8080


# app entrypoint: 1st command executed by docker
ENTRYPOINT ["./application"]

# command params
CMD [""]
```

### Multi OS arch support

```dockerfile
# debian based image
FROM golang:1.24 AS builder

# args for building
ARG TARGETOS
ARG TARGETARCH

# set env vars
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH


# mkdir /app && cd /app
WORKDIR /app

# cp from to
COPY . .

# exec command in current env
RUN go mod download

# buiid image
FROM GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 golang:1.24 AS builder

# mkdir /app && cd /app
WORKDIR /app

# cp from to
COPY . .

# exec command in current env
RUN go mod download

# go build
RUN go build -o application cmd/app/main.go

# deploy image
FROM alpine AS serve

# cp from builder to serve
COPY --from=builder /app/application .


ENV HTTP_SERVER_PORT=8080


# app entrypoint: 1st command executed by docker
ENTRYPOINT ["./application"]

# command params
CMD [""]
```

```sh
docker build --build-arg TARGETOS=linux --build-arg TARGET_ARCH=arm64 -t docker-bascis:arm64 .
```

### Docker registry

[Docker registry](https://www.docker.com/products/docker-hub/)

```sh
docker tag docker-basics mariolazzari/docker-basics:latest
docker push
```

### Running contaieners

```sh
docker run -p 1234:8080 docker-basics
# detached
docker run -d -p 1234:8080 docker-basics
docker stop #id
docker ps -a
docker start #id
docker rm #id
docker rename from_name to_name
# named container
docker run -d --name my_container -p 1234:8080 docker-basics
# force remove
docker rm -f #id
docker logs
docker inspects #id
```

### Docker volumes
