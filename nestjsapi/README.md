## Imersão 17 da FullCycle 🚀 <img src="https://nestjs.com/img/logo-small.svg" width="30" alt="Nest Logo" />

Projeto desenvolvido durante o Bootcamp Imersão 17 da FullCycle em 2024. Consiste em uma plataforma de E-commerce simples com BFF e Microserviços.

### Descrição:

- `/nestjsapi`: 
    - API REST desenvolvida em **NodeJS** com **TypeScript** e framework **NestJS** com **TypeORM**. 
    - Foi utilizado o **Docker** para container do banco de dados MySQL e também o **RabbitMQ** para publicar mensagens na fila para outros microserviços.

### Rodando o Projeto:

### 1- Pré Requisitos:

- NodeJS
- Docker

### 2- Configuração:

- Instalando as dependências...
```bash
npm install
```

- Subindo o container do Docker com o Banco de Dados e RabbitMQ:
```bash
docker-compose up -d
```

- Para preparar/resetar o Banco de Dados:
```bash
npm run fixture
```

### 3- Rodar a aplicação:

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

### Rascunho

- [Visualizar Rascunho.](https://github.com/paulopitta97/fullcycle-imersao17/blob/main/nestjsapi/RASCUNHO.md)

### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)