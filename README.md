# Desafio Final - API em Go

## :sparkles: Descrição

Este é o repositório do projeto **Desafio Final** do Bootcamp **Arquiteto(a) de Software**. A API foi desenvolvida com **Go** e utiliza o framework **Gin** para criar endpoints RESTful. O projeto tem como objetivo simular um sistema de gerenciamento de **clientes**, incluindo operações de **CRUD** (Create, Read, Update, Delete), e a contagem de clientes cadastrados.

### Funcionalidades:
- **Criação** de novos clientes
- **Listagem** de todos os clientes
- **Busca** de clientes por ID
- **Contagem** total de clientes

---

## :gear: Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal.
- **Gin Framework**: Framework web para construção de APIs.
- **GORM**: ORM utilizado para interação com o banco de dados SQLite.
- **SQLite**: Banco de dados leve para persistência dos dados dos clientes.

---

## :rocket: Como Rodar o Projeto

### :warning: Pré-requisitos
- **Go** (versão 1.16 ou superior)
- **Git** (para clonar o repositório)
- **SQLite** (banco de dados embutido)

### Passos para execução

1. **Clone o repositório**:
    ```bash
    git clone https://github.com/DylanMafi/desafio-final-api-go.git
    cd desafio-final-api-go
    ```

2. **Instale as dependências**:
    - O Go gerencia as dependências automaticamente com `go mod`. Você pode instalar as dependências necessárias com o comando:
    ```bash
    go mod tidy
    ```

3. **Inicie o servidor**:
    - Para rodar a aplicação localmente, use:
    ```bash
    go run main.go
    ```

    O servidor estará disponível em [http://localhost:8080](http://localhost:8080).

---

## :exclamation: Endpoints Disponíveis

| Método | Endpoint                | Descrição                              |
|--------|-------------------------|----------------------------------------|
| GET    | `/clientes`             | Retorna todos os clientes cadastrados. |
| GET    | `/clientes/:id`         | Retorna um cliente específico pelo ID. |
| POST   | `/clientes`             | Cria um novo cliente.                 |
| GET    | `/clientes/contar`      | Retorna o número total de clientes.    |

### :pencil2: Exemplos de Uso

#### **1. Listar todos os clientes:**

**Requisição**:
```http
GET http://localhost:8080/clientes
Resposta:

json
Copiar código
[
  {
    "id": 1,
    "nome": "João Silva",
    "email": "joao.silva@example.com"
  },
  {
    "id": 2,
    "nome": "Maria Oliveira",
    "email": "maria.oliveira@example.com"
  }
]
2. Criar um novo cliente:
Requisição:

http
Copiar código
POST http://localhost:8080/clientes
Content-Type: application/json

{
  "nome": "Carlos Souza",
  "email": "carlos.souza@example.com"
}
Resposta:

json
Copiar código
{
  "id": 3,
  "nome": "Carlos Souza",
  "email": "carlos.souza@example.com"
}
3. Contar o total de clientes:
Requisição:

http
Copiar código
GET http://localhost:8080/clientes/contar
Resposta:

json
Copiar código
{
  "total": 3
}
:octocat: Contribuindo
Contribuições são bem-vindas! Para contribuir com este projeto, siga os passos abaixo:

Faça um fork deste repositório.
Crie uma nova branch (git checkout -b minha-nova-feature).
Faça as modificações e commit as mudanças (git commit -am 'Adicionando nova feature').
Envie para o repositório remoto (git push origin minha-nova-feature).
Crie um Pull Request.
:memo: Licença
Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para mais detalhes.

:mailbox_with_mail: Contato
Se você tiver dúvidas ou sugestões, sinta-se à vontade para abrir uma issue ou entrar em contato comigo diretamente através do GitHub.

:link: Recursos Adicionais
Gin Documentation
GORM Documentation
SQLite Documentation
:trophy: Conclusão
Com este README.md, seu repositório está mais organizado e com uma aparência atraente. Ele agora fornece detalhes importantes, como o uso da API, contribuições e tecnologias, de forma clara e concisa, ajudando qualquer pessoa que acessar seu projeto a entender rapidamente o que ele faz e como usá-lo.

Se precisar de mais sugestões ou ajustes, me avise! 🚀

markdown
Copiar código

### **O que foi adicionado e melhorado:**
1. **Emojis**: Usei emojis nas seções para tornar o texto mais atraente visualmente, sem prejudicar a clareza.
2. **Exemplos e Formatação**: Todos os exemplos de código (requisições e respostas) estão bem formatados.
3. **Tabela de Endpoints**: Uma tabela foi criada para listar os endpoints de forma clara e organizada.
4. **Links**: Adicionei links úteis para documentação das tecnologias usadas no projeto.
5. **Instruções claras para contribuição**: A seção de contribuição foi organizada para ser fácil de entender.

### **Conclusão:**
Esse **README.md** torna o repositório muito mais atraente e facilita a navegação de outros desenvolvedores ou recrutadores. Ele contém todos os detalhes essenciais de maneira clara, permitindo que qualquer pessoa entenda como o projeto funciona e como contribuir.

Se precisar de mais alguma coisa, estou à disposição! 🚀
