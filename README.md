# 📝 CRUD de Tarefas em Go

Uma aplicação web simples de lista de tarefas (to-do list) construída com **Go puro** no backend e **HTML + CSS + JavaScript** no frontend. Sem frameworks, sem dependências externas — apenas a biblioteca padrão do Go.

![Print da aplicação](./go%20app.png)

---

## ✨ Funcionalidades

- ✅ **Criar** tarefas
- 📋 **Listar** todas as tarefas
- 🔍 **Buscar** tarefa por ID
- ✏️ **Atualizar** título e status (concluída/pendente)
- 🗑️ **Excluir** tarefas
- 💾 Armazenamento em memória (sem banco de dados)

---

## 🛠️ Tecnologias

| Camada | Tecnologia |
|---|---|
| Backend | Go 1.27.1 (biblioteca padrão `net/http`) |
| Frontend | HTML5, CSS3, JavaScript (Fetch API) |
| Armazenamento | Em memória (`slice` protegido por `sync.Mutex`) |

---

## 📁 Estrutura do Projeto

```
crud-go/
├── main.go              # Servidor HTTP + lógica do CRUD
├── go.mod               # Módulo Go
├── README.md            # Este arquivo
├── docs/
│   └── screenshot.png   # Print da aplicação
└── static/
    ├── index.html       # Interface
    ├── style.css        # Estilos
    └── app.js           # Lógica do frontend
```

---

## 🚀 Como executar

### Pré-requisitos

- **Go 1.27.1** ou superior instalado ([download](https://go.dev/dl/))
- Terminal Linux, macOS ou WSL no Windows

Verifique:

```bash
go version
# go version go1.27.1 linux/amd64
```

### Passos

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/crud-go.git
   cd crud-go
   ```

2. **Execute a aplicação:**

   ```bash
   go run main.go
   ```

3. **Abra no navegador:**

   ```
   http://localhost:8080
   ```

Pronto! A interface estará disponível e a API respondendo em `/tasks`.

---

## 🔌 API REST

Base URL: `http://localhost:8080`

| Método | Endpoint | Descrição | Body (JSON) |
|---|---|---|---|
| `GET` | `/tasks` | Lista todas as tarefas | — |
| `POST` | `/tasks` | Cria uma nova tarefa | `{"title": "...", "done": false}` |
| `GET` | `/tasks/{id}` | Busca tarefa por ID | — |
| `PUT` | `/tasks/{id}` | Atualiza tarefa | `{"title": "...", "done": true}` |
| `DELETE` | `/tasks/{id}` | Remove tarefa | — |

### Exemplos com `curl`

**Criar tarefa:**

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Estudar Go","done":false}'
```

**Listar tarefas:**

```bash
curl http://localhost:8080/tasks
```

**Atualizar tarefa:**

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Estudar Go Avançado","done":true}'
```

**Excluir tarefa:**

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 🧠 Como funciona

- O servidor Go registra três rotas:
  - `/` → serve os arquivos estáticos do frontend
  - `/tasks` → lida com listagem (GET) e criação (POST)
  - `/tasks/{id}` → lida com busca, atualização e exclusão por ID
- As tarefas ficam em um `slice` em memória, protegido por `sync.Mutex` para evitar condições de corrida.
- O frontend usa `fetch` para consumir a API e renderiza a lista dinamicamente.

> ⚠️ **Atenção:** as tarefas são perdidas ao reiniciar o servidor, pois não há persistência.

---

## 🧪 Testando

### Manualmente (navegador)

1. Adicione uma tarefa pelo formulário
2. Marque como concluída pelo checkbox
3. Exclua pelo botão vermelho

### Com `curl`

Use os exemplos da seção [API REST](#-api-rest).

---

## 🗺️ Roadmap

- [ ] Persistência com SQLite (`modernc.org/sqlite`)
- [ ] Edição inline do título
- [ ] Filtros: todas / pendentes / concluídas
- [ ] Testes automatizados (`net/http/httptest`)
- [ ] `//go:embed` para gerar binário único
- [ ] Docker

---

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch: `git checkout -b feature/minha-feature`
3. Commit: `git commit -m "feat: adiciona minha feature"`
4. Push: `git push origin feature/minha-feature`
5. Abra um Pull Request

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 👤 Autor

**Seu Nome**

- GitHub: [@seu-usuario](https://github.com/seu-usuario)
- Email: seu@email.com

---

⭐ Se este projeto te ajudou, deixe uma estrela no repositório!