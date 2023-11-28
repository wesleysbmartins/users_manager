# Users Manager
Este é um projeto de estudos sobre Golang que contiste em uma API com funcionalidades como criação, edição, busca e remoção de usuários.

Recursos utilizados pela aplicação:
- Pacotes net/http e gorilla/mux para servidor API Rest.
- Banco de dados PostgreSQL.
- Gorm para operações no banco de dados e migrations (criação de tabelas do banco).
- Docker para criar imagem e rodar aplicação no contaier.
- .env para variáveis de ambiente.

# Pré-requisitos:
- Docker (utilize a documentação do docker para instalação).
- Postman ou Insomnia para testar os endpoits da aplicação.

# Clone o Repositório e Execute a Aplicação
- Rode o comando no seu terminal para clonar o projeto:
```json
git clone git@github.com:wesleysbmartins/users_manager.git
```

- Configure seu arquivo .env. Você pode apenas renomear o .env.example e inserir as informações:
```json
# porta onde a aplicação vai rodar
PORT="8000"
# origens autorizadas para conexão (por exemplo um outro projeto front-end)
ALLOWED_ORIGINS="*"
# usuário do banco de dados
DB_USER="postgres"
# senha do banco de dados
DB_PASSWORD="password"
# nome do banco de dados
DB_NAME="manager"
# host da maquina que hospeda o banco de dados
DB_HOST="10.11.12.13"
# porta do banco
DB_PORT="5432"
# modo ssl do banco
DB_SSLMODE="disable"
```

- Baixe imagem do postgres no docker:
```json
docker pull postgres
```

- Rode a imagem do postgres no docker:
```json
docker run --name postgresql -e POSTGRES_PASSWORD=<sua senha do banco> -p <porta do banco>:<porta do banco> -d postgres
```

- Acesse o banco de dados e crie o banco de dados da aplicação. Você pode utilizar a extensão DatabaseClient do VSCode. 

- Crie imagem da aplicação no Docker:
```json
docker build -t users_manager .
```

- Execute a imagem da aplicação:
```json
docker run -p <porta da aplicação>:<porta da aplicação> users_manager
```