<div id="top">

<!-- ESTILO DO CABEÇALHO: MODERNO -->
<div align="left" style="position: relative; width: 100%; height: 100%; ">

<img src="./assets/frete-rapido-logo.png" width="30%" style="position: absolute; top: 0; right: 0;" alt="Logo Frete Rápido"/>

# SHIPPING-METRICS-API

<em><em>

<em>Construído com as ferramentas e tecnologias:</em>

<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Gin-008ECF.svg?style=for-the-badge&logo=Gin&logoColor=white" alt="Gin">
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=for-the-badge&logo=Docker&logoColor=white" alt="Docker">
<img src="https://img.shields.io/badge/Postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="Docker">

</div>
</div>
<br clear="right">

---

## Índice

I. [Índice](#indice)<br>
II. [Visão Geral](#visao-geral)<br>
III. [Estrutura do Projeto](#estrutura-do-projeto)<br>
IV. [Começando](#comecando)<br>
&nbsp;&nbsp;&nbsp;&nbsp;IV.a. [Pré-requisitos](#pre-requisitos)<br>
&nbsp;&nbsp;&nbsp;&nbsp;IV.b. [Instalação](#instalacao)<br>
&nbsp;&nbsp;&nbsp;&nbsp;IV.c. [Uso](#uso)<br>
&nbsp;&nbsp;&nbsp;&nbsp;IV.d. [Testes](#testes)<br>

---

## Visão Geral

Este projeto foi desenvolvido entre os dias 31 de julho à 4 de agosto de 2025, com base no teste técnico fornecido pela [Frete Rápido](https://freterapido.com.br)

---

## Estrutura do Projeto

```sh
└── shipping-metrics-api/
    ├── Dockerfile
    ├── README.md
    ├── cli
    ├── cmd
    │   ├── api
    │   └── docs
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── domain
    │   ├── handler
    │   ├── mapper
    │   ├── repository
    │   └── service
    ├── migrations
    │   ├── 000001_init.down.sql
    │   └── 000001_init.up.sql
    └── pkg
        ├── config
        ├── database
        ├── httputil
        └── validator
```

---

## Começando

### Pré-requisitos

Este projeto requer as seguintes dependências:

- **Linguagem de Programação:** Go
- **Gerenciador de Pacotes:** Go modules
- **Runtime de Contêiner:** Docker

### Instalação

Construa o shipping-metrics-api a partir do código-fonte e instale as dependências:

1. **Clone o repositório:**

   ```sh
   ❯ git clone https://github.com/deepzS2/shipping-metrics-api
   ```

2. **Navegue até o diretório do projeto:**

   ```sh
   ❯ cd shipping-metrics-api
   ```

3. **Replique o arquivo `.env.example` com as suas variáveis de ambiente:**

   ```sh
   ❯ cp .env.example .env # Linux
   # OU
   ❯ copy .env.example .env # Windows
   ```

### Uso

Execute o projeto com:

**Usando [docker-compose](https://docs.docker.com/compose/):**

```sh
docker-compose up -d
```

**Usando [go modules](https://golang.org/):**

```sh
go run ./cmd/api/main.go
```

### Testes

`shipping-metrics-api` usa o framework de testes **testify**. Execute a suíte de testes com:

**Usando [go modules](https://golang.org/):**

```sh
go test ./...
```

<div align="right">

[![][voltar-ao-topo]](#top)

</div>

[voltar-ao-topo]: https://img.shields.io/badge/-VOLTAR_AO_TOPO-151515?style=flat-square

---
