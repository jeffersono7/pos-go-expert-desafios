FROM golang:1.23 as build

WORKDIR /build

COPY . /build/

RUN GOOS=linux GOARCH=amd64 go build -o weather cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /build/weather /app/

EXPOSE 8080
ENTRYPOINT ["/app/weather"]
