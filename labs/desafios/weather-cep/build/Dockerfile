FROM golang:1.23 as build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o weather cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /build/weather /app/

EXPOSE 8080
ENTRYPOINT ["./weather"]
