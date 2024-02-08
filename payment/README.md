## Imersão 17 da FullCycle 🚀 <img src="https://go.dev/images/go-logo-white.svg" width="30" />

Projeto desenvolvido durante o Bootcamp Imersão 17 da FullCycle em 2024. Consiste em uma plataforma de E-commerce simples com BFF e Microserviços.

### Descrição:

- `/payment`:
    - Microserviço desenvolvido na linguagem **GO** que realiza o processamento de pagamentos com base nas filas no `RabbitMQ`.
    - Foi utilizado o **Docker** com container do **RabbitMQ** para consumir e publicar mensagens na fila para outros microserviços.

### Rodando o Projeto:

### 1- Pré Requisitos:

- Instalar o compilador da linguagem `GO` no Sistema Operacional: `https://go.dev/dl/`
- Docker

### 2- Configuração:

- Baixando as dependências...
```bash
go mod tidy
```

- Subindo o container do Docker com o RabbitMQ:
```bash
docker-compose up -d
```

- Abrir a admin do [RabbitMQ](http://localhost:15672/) e:
    - criar uma fila `orders`;
    - criar uma fila `orders_result` e vincular ela a exchange `amq.direct`;

### 3- Levantar a aplicação:

```bash
cd cmd/payment && go run main.go
```

### Rascunho

- [Visualizar Rascunho.](https://github.com/paulopitta97/fullcycle-imersao17/blob/main/payment/RASCUNHO.md)

### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)