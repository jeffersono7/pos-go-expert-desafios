FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service_b

FROM scratch
WORKDIR /app
COPY --from=build /app/service_b .
ENTRYPOINT [ "service_b" ]
