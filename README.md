# Case Eulabs Luanderson

> Olá nesse projeto temos como objetivo demonstrar conhecimentos de desenvolvimento em Golang utilizando echo frammework, 
  o sistema consiste em um CRUD de produtos.

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

- Utilizamos a versão `<v1.22.0>` do Go.
- Utilize o Backup do banco de dados e valide as configs de acordo com o arquivo em: configs/config.go

## 🚀 Instalando <Projeto>

Para instalar o <sistema> basta ter os requisitos acima e fazer o clone do projeto na sua máquina.


## ☕ Usando <Sistema>

Para usar <Sistema>, siga estas etapas:

dentro da raiz, rode o seguinte comando

```
<go run main.go>
```
Aqui você já estará com o BackEnd rodando na sua máquina na porta 9000.

## 🚀 decisões do projeto

Utilizei alguns conceitos para tornar o projeto mais escalável: 

- UUID ao invés de Id serial, como identificador primário decidi utilizar uuid baseado no artigo a seguir: https://espigah.medium.com/id-vs-uuid-vs-publickey-5f7e19455b90
  com o id universal e unico(uuid) adicionamos uma camada a mais de proteção para a nossa Api evitando que usuários possam varrer nossa base de dados e pegar alguns dados que podem ser sensíveis, além disso tornamos a validação por chave primária bem mais robusta e sólida;
- Montagem das configs de acordo com um arquivo toml, para processos onde utilizamos nuvem e dockerização da aplicação sempre é importante poder controlar isso e maneira externa ao código, para isso decidi construir as configs do projeto a partir de um arquivo toml.
