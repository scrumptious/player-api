#syntax=docker/dockerfile:1
#my player-api microsevice in Go packaged into a container image
FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
ENV GOOS=linux
RUN go build .
FROM alpine:latest
WORKDIR /usr/home
COPY --from=builder / /usr/home/player-api
EXPOSE 8080
CMD ["/usr/home/player-api"]