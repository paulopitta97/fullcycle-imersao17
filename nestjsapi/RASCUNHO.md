### Rascunho

1- Link NESTJS: docs.nestjs.com

2- Instalar o NestJS Global:

`npm install -g @nestjs/cli`

3- Inicializando o Projeto:

`nest new nestjsapi`

3.1- Sugestão: instalar as extensões Prettier + ESLint.

`npm run start:dev`

3.2- Configurando os módulos através do Nest.

`nest g module products`
    - cria apenas o módulo

`nest g resource` 
    - digitar `products` / `REST API` / `Y` -> gera CRUD

3.3- Instalando o TypeORM:

`npm install typeorm @nestjs/typeorm mysql2`

3.4- Configurando o Docker:
`docker compose up`
`docker compose exec db bash`
`mysql -uroot -proot`

3.5- Refazer o banco de dados padrão:
`npm run fixture`

3.6- Instalando dependências de validação:
`npm install class-validator class-transformer`

3.7- Instalando JWT para trabalhar com autenticação por token:
`npm install @nestjs/jwt`

3.7.1- Configurando o 'guardião' com módulo JWT/auth:
`nest g guard auth`

3.8- Acessando ADMIN do RabbitMQ: http://localhost:15672/

3.9- Instalando lib RabbitMQ para mensageria:
`npm install @golevelup/nestjs-rabbitmq` 
    - sugeriu usar essa e não a lib nativa do NestJS por questões de configs...

3.9.1- Configurando módulo para RabbitMQ:
`nest g module rabbitmq`

- Na Admin: linkar a fila (queue/streams) com a exchange:
    - From exchange: `amq.direct`
    - Routing key: `OrderCreated`

##### Vídeo: https://www.youtube.com/watch?v=A6dYq1ce34o