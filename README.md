## Imersão 17 da FullCycle 🚀

Projeto desenvolvido durante o Bootcamp Imersão 17 da FullCycle em 2024. Consiste em uma plataforma de E-commerce simples com BFF e Microserviços.

<img src="https://camo.githubusercontent.com/2e3480588544bcaa81f8d123c8985384d60deb014611646d5d30ed80f25b78cb/68747470733a2f2f6576656e74732d66756c6c6379636c652e73332e616d617a6f6e6177732e636f6d2f6576656e74732d66756c6c6379636c652f7374617469632f736974652f696d672f677275706f5f343431372e706e67"/>

### Consiste em:

- `/goapi`: 
    - _Resumo:_ `API para Categorias e Produtos`
    - API REST desenvolvida na linguagem **GO** que possui as rotas para CRUD de `category` e `product`. 
    - Foi utilizado o **Docker** para container do banco de dados MySQL.
- `/nestjsapi`: 
    - _Resumo:_ `API para Pedidos e autenticação do usuário/cliente`
    - API REST desenvolvida em **NodeJS** com **TypeScript** e framework **NestJS** com **TypeORM**. 
    - Foi utilizado o **Docker** para container do banco de dados MySQL e também o **RabbitMQ** para publicar mensagens na fila para outros microserviços.
- `/nextjs-frontend`: 
    - _Resumo:_ `FRONT-END com NextJS aproveitando as vantagens do server-side components e cache`
    - Front-end desenvolvido em **NodeJS** com **TypeScript** e framework **NextJS** com **Material Design/UI**.
	- Foi implementado _server components_ e _client components_, além de classes de serviço e as _server actions_.
- `/payment`: 
    - _Resumo:_ `Microserviço para processar pagamentos`
    - Microserviço desenvolvido na linguagem **GO** que realiza o processamento de pagamentos com base nas filas no `RabbitMQ`.
    - Foi utilizado o **Docker** com container do **RabbitMQ** para consumir e publicar mensagens na fila para outros microserviços.

### Detalhes:

As instruções para executar se encontram no `README.md` de cada subpasta.

- Ordem de execução dos projetos:

1. goapi
1. nestjsapi
1. payment-service
1. nextjs-frontend

- Executando o projeto:

```sh
## Levantando a GO API:
cd goapi
docker compose up -d
go run ./cmd/catalog/main.go

## Levantando a NEST JS API:
cd nestjsapi
docker compose up -d
npm install
npm run fixture
npm run start:dev

## Levantando o NEXT JS FRONTEND:
cd nextjs-frontend
npm install
npm run dev
npx serve -l 9000 images

## Levantando o MSVC PAYMENT: 
## (obs: não precisa levantar novo container para o RabbitMQ com o docker aqui não)
cd ./payment/cmd/payment && go run main.go
```

<details open>
 <summary>Imagem da Aplicação</summary>

 ![imagem-projeto](https://github.com/paulopitta97/fullcycle-imersao17/assets/52472087/1ce1d7ca-0ec9-4ed1-a762-63f9b986f7db)
</details>

<details>
 <summary>Configuração RabbitMQ</summary>

![filas-rabbitmq-01](https://github.com/paulopitta97/fullcycle-imersao17/assets/52472087/7de88f42-cc0a-4777-b6f7-fdd78198e0d0)

![filas-rabbitmq-02-orders](https://github.com/paulopitta97/fullcycle-imersao17/assets/52472087/ca19ecef-1d9e-4232-8e8d-d7d99abcfa9b)

![filas-rabbitmq-03-payments](https://github.com/paulopitta97/fullcycle-imersao17/assets/52472087/08e5c1cd-2582-4ddb-b765-bf032b8feae2)

</details>


### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)
