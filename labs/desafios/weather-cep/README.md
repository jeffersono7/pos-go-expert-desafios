# Weather API

Essa api é um desafio de labs da pós go expert

## Como usar

### local
Alterar no deployments/docker-compose.yml a key da api weather

```shell
docker-compose -f deployments/docker-compose.yml up --build -d

curl -v localhost:8080/cep/68903121/weather
```

### prd
```shell
curl -v https://weather-739284750138.us-east1.run.app/cep/68903121/weather
```

## Como testar

```shell
go test ./...
```
