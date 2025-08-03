# 🧰 Go API Produtos

API simples desenvolvida com Go, utilizando o framework **Gin**, integrada com **PostgreSQL** e empacotada com **Docker**.

> Este projeto realiza operações de **CRUD** e tem como objetivo servir como **modelo para aplicações mais robustas** em Go.

---

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

---

## 🐳 Como rodar o projeto com Docker

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/go-api-produtos.git
   cd go-api-produtos
   ```

2. Suba os containers com o Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Acesse a API na porta `8000`. Para testar se está rodando:
   ```
   GET http://localhost:8000/ping
   ```
   Retorna:
   ```json
   "pong"
   ```

---

## 🔗 Principais Endpoints

| Método | Rota                      | Descrição                        |
|--------|---------------------------|----------------------------------|
| GET    | `/products`               | Lista todos os produtos          |
| POST   | `/product`                | Cria um novo produto             |
| GET    | `/product/:productId`     | Busca um produto por ID          |
| PUT    | `/product/:productId`     | Atualiza um produto por ID       |
| DELETE | `/product/:productId`     | Deleta um produto por ID         |

---

## 📦 Payload de exemplo para criar produto

```json
{
  "name": "Tenis",
  "price": 59.29
}
```

---

## 👤 Autor

- Nome: João Miguel Fortunato Rita
- GitHub: https://github.com/JoaoMiguelRita
- Linkedin: https://www.linkedin.com/in/jo%C3%A3o-miguel-fortunato-rita-623962219/


---

## 📄 Licença

Este projeto está sob a licença que você preferir.  
Sinta-se livre para reutilizar e modificar!