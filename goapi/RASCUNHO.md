### Rascunho

1- Instalar o GO no sistema operacional: `https://go.dev/dl/`

2- Instalar a extensão GO no VSCODE.

- `CTRL + P` e `>go:Install/Update Tools`
- Marcar todos e instalar.

3- Instalar o gerenciador de bibliotecas do GO

`go mod init github.com/paulopitta97/fullcycle-imersao17/goapi`

- Considerações: 
    - em GO não tem classes e sim Struct, é o mais equivalente...
    - a pasta Internal vai armazenar tudo que for interno da aplicação...

4- Comando para baixar dependencias:

`go mod tidy`

5- Subindo o container do Docker com o BD:

`docker-compose up -d`

- Obs: o erro ao levantar o container referente a porta 3306 foi por estar com o serviço do MySQL já iniciado na máquina local
    - Error response from daemon: Ports are not available: exposing port TCP 0.0.0.0:3306 -> 0.0.0.0:0: listen tcp 0.0.0.0:3306: bind: 
    - Normalmente é permitida apenas uma utilização de cada endereço de soquete (protocolo/endereço de rede/porta).

`docker-compose logs mysql`
`docker-compose exec mysql bash`
`mysql -uroot -p imersao17`
- Rodar o SQL e `show tables;`

6- Configurando TESTES DE REQUISIÇÕES HTTP.

`cd cmd/catalog && go run main.go`

- Instalar a extensão REST Client >> https://marketplace.visualstudio.com/items?itemName=humao.rest-client

##### Vídeo: https://www.youtube.com/watch?v=WCsVUYXeBUA

7- Apagando o banco de dados:

`sudo rm -rf ./.docker/mysql`
`docker compose down` para destruir o container do docker.