# Go Web Scraper

Este é um projeto básico de **web scraping** em Go que busca o título de uma página HTML. Ele faz uma requisição HTTP, parseia o HTML e extrai informações específicas (como o título da página).

## Funcionalidades

- Realiza requisições HTTP para uma URL fornecida.
- Parseia o HTML da resposta.
- Extrai o título da página (ou outros dados, se desejado).

## Pré-requisitos

- **Go** instalado na sua máquina. [Instalar Go](https://golang.org/doc/install).
- Conexão com a internet para fazer requisições HTTP.

## Como executar

### 1. Clonar o repositório ou criar o diretório

```bash
git clone https://github.com/seu-usuario/go-web-scraper.git
cd go-web-scraper
```

### 1. Inicializar o projeto

Se você estiver criando o projeto do zero, use o comando abaixo:

```bash
go mod init go-web-scraper
```

### 2. Instalar dependências

Instale o pacote golang.org/x/net/html para permitir o parsing de HTML:
```bash
go get golang.org/x/net/html
```

### 3. Rodar o projeto

Após a configuração, rode o scraper com:

```bash
go run main.go
```

### Saída esperada

Você deve ver uma saída similar à seguinte:

```bash
Fetching URL: https://golang.org
Título da página: The Go Programming Language
```

