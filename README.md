
<div  align="center">
    <img  src="https://miro.medium.com/max/700/1*XtU2ByqGUYD4O5TQuS2CxQ.png"  alt="marvel"  width="320"/>
</div>

# Go Challenge 20200923 - Coodesh
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)


### API REST utilizando Go + MongoDB

API REST para consulta de informações do banco de dados open source 'Open Food Facts', feita utilizando Go como 
linguagem server-side e MongoDB para a persistência de dados.

#### Bibliotecas utilizadas

- Mongo-go-driver `go get go.mongodb.org/mongo-driver`
- Cron v3 `go get github.com/robfig/cron/v3`
- Gorilla Mux `go get github.com/gorilla/mux`
- net/http `Built-in`
- Godotenv `go get github.com/joho/godotenv`

### Instruções (Instalação e utilização)

#### Ferramentas necessárias :

| Ferramenta  |  Windows  |   Linux(Ubuntu)  |    Web  | 
| ------------------- | ------------------- | ------------------- |  ------------------- | 
|  Go |   [Instalação](https://golang.org/doc/install) |  [Instalação](https://golang.org/doc/install) |   X |
|  Docker |  [Instalação](https://br.openfoodfacts.org/data) | [Instalação](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04-pt)  |  X |
|  MongoDB Atlas | X|X | [Criação de conta](https://www.mongodb.com/try) | 
|  Insomnia |  [Instalação](https://insomnia.rest/download/) | [Instalação](https://snapcraft.io/insomnia)  |   X |
|  GIT |  [Instalação](https://git-scm.com/downloads) | [Instalação](https://git-scm.com/downloads)  |   X |

#### Obtendo o projeto e rodando:

Tendo o Docker instalado e configurado em seu ambiente, assim como o Git e o Insomnia(ou Postman, caso prefira), 
clone o repositório para sua máquina local através do seguinte comando no terminal `git clone https://github.com/ThyagoFr/coodesh.git`.
Navegue através do terminal até a pasta raiz do repositório anteriormente clonado e com o Docker configurado utilize os seguintes comandos em sequência :
- `docker build --tag gomongo:1.0 .`
- `docker run --p 8000:8080 --d --name gomongo-api gomongo:1.0`
- `docker container ps -a`

Após o ultimo comando você verá no terminal um container em execução chamado **gomongo-api** e isso significa que a aplicação já esta em 
execução em sua máquina local e pode ser acessada através das seguintes rotas :

- `GET http://localhost:8080/api/v1/`: Detalhes da API, se conexão leitura e escritura com a base de dados está OK, horário da última vez que o CRON foi executado, tempo online e uso de memória.
- `PUT http://localhost:8080/api/v1/products/:code`: Rota para atualização de dados de um produto específico.
- `DELETE http://localhost:8080/api/v1/products/:code`: Rota para mudar o status de um produto específico para `trash`
- `GET http://localhost:8080/api/v1/products/:code`: Rota para obter a informação somente de um produto da base de dados.
- `GET http://localhost:8080/api/v1/products?size=10&page=0`: Rota para listar produtos cadastrados na aplicação. Por DEFAULT tal rota só irá retornar (se possuir) os 
10 primeiros produtos da base de dados. Caso queira retornar mais produtos, utilizar o query param `size` sendo a quantidade de produtos
desejada e `page` para especificar à "página" contendo os produtos. 

#### Teste da API utilizando Insomnia:

Logo abaixo segue imagens de alguns testes utilizando o REST Client Insomnia, incluindo testes que provocam
respostas com HTTP Status como `NOT FOUND (404)` e `BAD REQUEST(400)`.

- Buscando produtos
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/GETALL.png?raw=true)
- Buscando produtos (Paginação com size específico)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/GETALL-PAGINATE.png?raw=true)
- Buscando produto por código (OK)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/GETPRODUCT-OK.png?raw=true)
- Buscando produto por código (NOT FOUND)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/GETPRODUCT-NOTFOUND.png?raw=true)
- Removendo produto por codigo (OK)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/DELETEPRODUCT-OK.png?raw=true)
- Removendo produto por codigo (NOT FOUND)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/DELETEPRODUCT-NOTFOUND.png?raw=true)
- Editando produto por código (NOT FOUND)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/PUTPRODUCT-NOTFOUND.png?raw=true)
- Editando produto por codigo (OK)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/PUTPRODUCT-OK.png?raw=true)
- Editando produto por codigo (BAD REQUEST)
![Buscando todos os produtos](https://github.com/thyagofr/coodesh/blob/main/imgs/PUTPRODUCT-BADREQUEST.png?raw=true)
 