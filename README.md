Este é um projeto de API RESTful desenvolvido em Go (Golang) para gerenciar produtos. Ele utiliza o framework [Gin](https://github.com/gin-gonic/gin) para lidar com as rotas HTTP e o banco de dados PostgreSQL para persistência de dados.

## Funcionalidades

- **Listar Produtos**: Retorna todos os produtos cadastrados no banco de dados.
- **Criar Produto**: Permite criar um novo produto com nome e preço.
- **Buscar Produto por ID**: Retorna os detalhes de um produto específico com base no seu ID.

## Estrutura do Projeto

O projeto segue uma arquitetura limpa, separando responsabilidades em diferentes camadas:

- **Controller**: Gerencia as requisições HTTP e interage com a camada de usecase.
- **Usecase**: Contém a lógica de negócios da aplicação.
- **Repository**: Responsável pela interação com o banco de dados.
- **Model**: Define as estruturas de dados utilizadas no projeto.

### Principais Arquivos e Diretórios

- `cmd/main.go`: Ponto de entrada da aplicação.
- `controller/`: Contém os controladores responsáveis pelas rotas.
- `usecase/`: Contém a lógica de negócios.
- `repository/`: Contém a lógica de acesso ao banco de dados.
- `model/`: Contém as definições de modelos de dados.
- `db/conn.go`: Configuração e conexão com o banco de dados.
- `Dockerfile`: Configuração para criar a imagem Docker da aplicação.
- `docker-compose.yml`: Configuração para subir a aplicação e o banco de dados com Docker.

## Pré-requisitos

- [Go](https://golang.org/) 1.23.5 ou superior
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL

## Configuração e Execução

### Usando Docker

1. Certifique-se de que o Docker e o Docker Compose estão instalados.
2. Execute o comando abaixo para iniciar a aplicação e o banco de dados:

   ```bash
   docker-compose up --build
3. A API estará disponível em http://localhost:8000.

Sem Docker
1. Configure um banco de dados PostgreSQL com as credenciais definidas em db/conn.go.

2. Instale as dependências do projeto:

3. Execute a aplicação:

4. A API estará disponível em http://localhost:8000.

## Endpoints

### Testando com requests.http

O arquivo requests.http contém exemplos de requisições para testar a API. Você pode utilizá-lo com extensões como REST Client no VS Code.

### Exemplos de Endpoints

#### Ping: Verifica se a API está funcionando.

GET http://localhost:8000/ping

#### Listar Produtos:

GET http://localhost:8000/products

#### Criar Produto:

POST http://localhost:8000/product
Content-Type: application/json

{
  "name": "Coca-cola",
  "price": 5.55
}

#### Buscar Produto por ID:

GET http://localhost:8000/product/1

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework Web**: Gin
- **Banco de Dados**: PostgreSQL
- **Gerenciamento de Dependências**: Go Modules
- **Containerização**: Docker e Docker Compose

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.
