# Desafio Final - API em Go 🚀

## ✨ Descrição
Este é o repositório do projeto Desafio Final do Bootcamp Arquiteto(a) de Software. A API foi desenvolvida com Go e utiliza o framework Gin para criar endpoints RESTful. O objetivo é simular um sistema de gerenciamento de clientes, com operações CRUD (Create, Read, Update, Delete) e contagem de clientes cadastrados.

## Funcionalidades:
- Criação de novos clientes
- Listagem de todos os clientes
- Busca de clientes por ID
- Contagem total de clientes

## ⚙️ Tecnologias Utilizadas
- **Golang**: Linguagem de programação principal.
- **Gin Framework**: Framework web para construção de APIs.
- **GORM**: ORM para interação com o banco de dados SQLite.
- **SQLite**: Banco de dados leve para persistência dos dados dos clientes.

## 🚀 Como Rodar o Projeto

### ⚠️ Pré-requisitos:
- Go (versão 1.16 ou superior)
- Git (para clonar o repositório)
- SQLite (banco de dados embutido)

### Passos para execução:
1. Clone o repositório:
    ```bash
    git clone https://github.com/DylanMafi/desafio-final-api-go.git
    cd desafio-final-api-go
    ```

2. Instale as dependências:
    O Go gerencia as dependências automaticamente com o comando `go mod`. Para instalar as dependências necessárias, use:
    ```bash
    go mod tidy
    ```

3. Inicie o servidor:
    Para rodar a aplicação localmente, use o comando:
    ```bash
    go run main.go
    ```
    O servidor estará disponível em [http://localhost:8080](http://localhost:8080).

## ❗ Endpoints Disponíveis
| Método | Endpoint            | Descrição                             |
|--------|---------------------|---------------------------------------|
| GET    | /clientes           | Retorna todos os clientes cadastrados.|
| GET    | /clientes/:id       | Retorna um cliente específico pelo ID.|
| POST   | /clientes           | Cria um novo cliente.                 |
| GET    | /clientes/contar    | Retorna o número total de clientes.   |

## ✏️ Exemplos de Uso

1. **Listar todos os clientes:**
    - Requisição: `GET http://localhost:8080/clientes`
    - Resposta:
    ```json
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
    ```

2. **Criar um novo cliente:**
    - Requisição:
    ```bash
    POST http://localhost:8080/clientes
    Content-Type: application/json
    ```
    - Corpo da Requisição:
    ```json
    {
      "nome": "Carlos Souza",
      "email": "carlos.souza@example.com"
    }
    ```
    - Resposta:
    ```json
    {
      "id": 3,
      "nome": "Carlos Souza",
      "email": "carlos.souza@example.com"
    }
    ```

3. **Contar o total de clientes:**
    - Requisição:
    ```bash
    GET http://localhost:8080/clientes/contar
    ```
    - Resposta:
    ```json
    {
      "total": 3
    }
    ```

## :octocat: Contribuindo
Contribuições são bem-vindas! Para contribuir com este projeto, siga os seguintes passos:

1. Faça um fork deste repositório.
2. Crie uma nova branch:
    ```bash
    git checkout -b minha-nova-feature
    ```
3. Faça as modificações e commit as mudanças:
    ```bash
    git commit -am 'Adicionando nova feature'
    ```
4. Envie para o repositório remoto:
    ```bash
    git push origin minha-nova-feature
    ```
5. Crie um Pull Request.

## :memo: Licença
Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para mais detalhes.

## :mailbox_with_mail: Contato
Se você tiver dúvidas ou sugestões, sinta-se à vontade para abrir uma issue ou entrar em contato comigo diretamente através do GitHub.

## :link: Recursos Adicionais
- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [SQLite Documentation](https://www.sqlite.org/docs.html)
