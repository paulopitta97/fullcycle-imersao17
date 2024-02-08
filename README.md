## Imersão 17 da FullCycle 🚀

Projeto desenvolvido durante o Bootcamp Imersão 17 da FullCycle em 2024. Consiste em uma plataforma de E-commerce simples com BFF e Microserviços.

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

### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)