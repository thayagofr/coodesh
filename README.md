
<div  align="center">
    <img  src="https://miro.medium.com/max/700/1*XtU2ByqGUYD4O5TQuS2CxQ.png"  alt="marvel"  width="320"/>
</div>

# Go Challenge 20200923 - Coodesh
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)


### API REST utilizando Go + MongoDB

API REST para consulta de informações do banco de dados open source 'Open Food Facts', feita utilizando Go como 
linguagem server-side e MongoDB para a persistência de dados.

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
 

## O projeto
 
- Criar um banco de dados MongoDB usando Atlas: https://www.mongodb.com/cloud/atlas ou algum Banco de Dados SQL se não sentir confortável com NoSQL;
- Criar uma REST API usando GoLang com as melhores práticas de desenvolvimento.
- Integrar a API com o banco de dados criado para persistir os dados
- Recomendável usar Drivers oficiais para integração com o DB
- Desenvolver Testes Unitários

### Modelo de Dados:

Para a definição do modelo, consultar o arquivo [products.json](./products.json) que foi exportado do Open Food Facts, um detalhe importante é que temos dois campos personalizados para poder fazer o controle interno do sistema e que deverão ser aplicados em todos os alimentos no momento da importação, os campos são:

- `imported_t`: campo do tipo Date com a dia e hora que foi importado;
- `status`: campo do tipo Enum com os possíveis valores draft, trash e published;

### Sistema do CRON

Para prosseguir com o desafio, precisaremos criar na API um sistema de atualização que vai importar os dados para a Base de Dados com a versão mais recente do [Open Food Facts](https://br.openfoodfacts.org/data) uma vez ao día. Adicionar aos arquivos de configuração o melhor horário para executar a importação.

A lista de arquivos do Open Food, pode ser encontrada em: 

- https://static.openfoodfacts.org/data/delta/index.txt

Onde cada linha representa um arquivo que está disponível em https://static.openfoodfacts.org/data/delta/{filename}. O nome do arquivo contém o timestamp UNIX da primeira e última alteração contida no arquivo JSON, para que os arquivos possam ser importados (após extracção) ordenados.

É recomendável utilizar uma Collection secundária para controlar os históricos das importações e facilitar a validação durante a execução.

Ter em conta que:

- Todos os produtos deverão ter os campos personalizados `imported_t` e `status`.
- Limitar a importação a somente 100 produtos;

### A REST API

Na REST API teremos um CRUD com os seguintes endpoints:

 - `GET /`: Detalhes da API, se conexão leitura e escritura com a base de dados está OK, horário da última vez que o CRON foi executado, tempo online e uso de memória.
 - `PUT /products/:code`: Será responsável por receber atualizações do Projeto Web
 - `DELETE /products/:code`: Mudar o status do produto para `trash`
 - `GET /products/:code`: Obter a informação somente de um produto da base de dados
 - `GET /products`: Listar todos os produtos da base de dados, adicionar sistema de paginação para não sobrecarregar o `REQUEST`.

## Extras

- **Diferencial 1** Front End con ReactJs, configurar um projeto web para listar os produtos cadastrados na REST API.
- **Diferencial 2** Configurar Docker no Projeto para facilitar o Deploy da equipe de DevOps;
- **Diferencial 3** Configurar um sistema de alerta se tem algum falho durante o Sync dos produtos;
- **Diferencial 4** Descrever a documentação da API utilizando o conceito de Open API 3.0;
- **Diferencial 5** Escrever Unit Tests para os endpoints  GET e PUT do CRUD;


## Readme do Repositório
 
- Deve conter o título de cada projeto
- Uma descrição de uma frase
- Como instalar e usar o projeto (instruções)
- Não esqueça o [.gitignore](https://www.toptal.com/developers/gitignore)
 
## Finalização 

Avisar sobre a finalização e enviar para correção em: [https://coodesh.com/review-challenge](https://coodesh.com/review-challenge) 
 
Após essa etapa será marcado a apresentação/correção do projeto.

## Instruções para a Apresentação: 

1. Será necessário compartilhar a tela durante a vídeo chamada;
2. Deixe todos os projetos de solução previamente abertos em seu computador antes de iniciar a chamada;
3. Deixe os ambientes configurados e prontos para rodar; 
4. Prepara-se pois você será questionado sobre cada etapa e decisão do Challenge;
5. Prepare uma lista de perguntas, dúvidas, sugestões de melhorias e feedbacks (caso tenha).


## Suporte

Use o nosso canal no slack: http://bit.ly/32CuOMy para tirar dúvidas sobre o processo ou envie um e-mail para contato@coodesh.com. 
