## Imersão 17 da FullCycle 🚀 <img src="https://go.dev/images/go-logo-white.svg" width="30" alt="Nest Logo" />

Projeto desenvolvido durante o Bootcamp Imersão 17 da FullCycle em 2024. Consiste em uma plataforma de E-commerce simples com BFF e Microserviços.

### Descrição:

- `/goapi`:
    - API REST desenvolvida na linguagem **GO** que possui as rotas para CRUD de `category` e `product`. 
    - Foi utilizado o **Docker** para container do banco de dados MySQL.

### Rodando o Projeto:

### 1- Pré Requisitos:

- Instalar o compilador da linguagem `GO` no Sistema Operacional: `https://go.dev/dl/`
- Docker

### 2- Configuração:

- Baixando as dependências...
```bash
go mod tidy
```

- Subindo o container do Docker com o Banco de Dados:
```bash
docker-compose up -d
```

- Acessando o MySQL:
```bash
docker-compose exec mysql bash
```

- Conectando ao usuário no MySQL:
```cmd
mysql -uroot -p imersao17
```

- Rodar o SQL disponível no arquivo `db.sql`;<br><br>

- Rodar `show tables;` e verificar se as tabelas foram devidamente criadas;<br><br>

### 3- Levantar a aplicação:

```bash
cd cmd/catalog && go run main.go
```

### Rascunho

- [Visualizar Rascunho.](https://github.com/paulopitta97/fullcycle-imersao17/blob/main/goapi/RASCUNHO.md)

### 📝 Licença

- [MIT](https://github.com/paulopitta97/fullcycle-imersao17/blob/master/LICENSE) © [Paulo Pitta](https://github.com/paulopitta97)