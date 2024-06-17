FROM --platform=linux/amd64 golang:1.21-alpine
RUN apk update && apk upgrade zlib && apk upgrade apk-tools
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/scrumptious/weather-service
COPY . .
RUN go test -v
RUN go build -v -o app
CMD ["./app"]