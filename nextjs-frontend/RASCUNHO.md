### Rascunho

1- Inicializando o projeto:
`npx create-next-app --typescript nextjs-frontend`
`npm run dev`

2- Instalando dependências de estilização:
`npm install @mui/material-nextjs @emotion/cache @emotion/react @emotion/styled`
`npm install @mui/material`
`npm install @mui/icons-material`

3- Instalando dependência para hook formulário:
`npm install react-hook-form`
`npm install @hookform/resolvers yup`

4- Levantando servidor de imagens produtos:
`npx serve -l 9000 images`

5- Exemplo de publicação de uma mensagem na fila para mudança de status do pedido:
```
{
  "order_id": "634dadfb-fa9c-461b-9b5c-3896e445f923",
  "status": "PAID"
}
```