# Desafio MultiThreading

### Como usar
```sh
go run main.go <cep> <cep> <cep>....
```

### Como funciona
Vai instanciar 3 workers para processar os ceps mais rapidamente.

Em cada worker vai disparar 2 requests simultânea para brasilapi e viacep, quem responder mais rápido ele retorna a resposta, tudo limitado em 1s no máximo para ter uma resposta do cep consultado.
