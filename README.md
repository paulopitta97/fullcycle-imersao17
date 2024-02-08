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

### Detalhes:

As instruções para executar se encontram no `README.md` de cada subpasta.

- Ordem de execução dos projetos:

1. goapi
1. nestjsapi
1. payment-service
1. nextjs-frontend

### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)