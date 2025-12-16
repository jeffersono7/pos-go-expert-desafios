FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o service_b cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/service_b .
ENTRYPOINT [ "./service_b" ]
