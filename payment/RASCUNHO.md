### Rascunho

1- Instalar o gerenciador de bibliotecas do GO

`go mod init github.com/paulopitta97/fullcycle-imersao17/payment`

2- Comando para baixar dependencias:

`go mod tidy`

3- Subindo o container do Docker com o BD:

`docker-compose up -d`

4- Abrir a admin do [RabbitMQ](http://localhost:15672/) e criar uma fila `orders`.

4.1- Criar uma fila `orders_result` e vincular ela a exchange `amq.direct`

5- Rodando a aplicação:

`cd cmd/payment && go run main.go`