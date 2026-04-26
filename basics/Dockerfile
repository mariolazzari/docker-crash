# debian based image
FROM golang:1.24 AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

# mkdir /app && cd /app
WORKDIR /app
# cp from to
COPY . .
# exec command in current env
RUN go mod download
# go build
RUN CGO_ENABLED=0 go build -o application cmd/app/main.go

FROM alpine AS serve

COPY --from=builder /app/application .

ENV HTTP_SERVER_PORT=8080
# app entrypoint: 1st command executed by docker
ENTRYPOINT ["./application"]
# command params
CMD [""]
