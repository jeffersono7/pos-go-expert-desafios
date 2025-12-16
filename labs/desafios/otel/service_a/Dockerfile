FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o service_a cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/service_a .
ENTRYPOINT [ "./service_a" ]
