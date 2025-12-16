FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service_a

FROM scratch
WORKDIR /app
COPY --from=build /app/service_a .
ENTRYPOINT [ "service_a" ]
