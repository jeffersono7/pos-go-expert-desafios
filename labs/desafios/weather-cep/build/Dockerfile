FROM golang:1.23 as build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=build /build/weather /app/

EXPOSE 8080
ENTRYPOINT ["./weather"]
