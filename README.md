# Testes de stress

Este projeto é uma apllicação command-line interface (CLI) desenvolvida em Go para realização de testes de stress que utiliza a biblioteca Cobra.

## Flags

-  -u, --url: Uma string que representa a URL para onde serão enviadas as requisições.
-  -r, --requests: Valor inteiro que representa a quantidade de requisições a serem realizadas. Senão for informado será utilizado o valor padrão de 1. 
-  -c, --concurrency: Valor inteiro que representa a quantidade de requisições concorrentes. Senão for informado será utilizado o valor padrão de 1.


## Utilização

Crie uma imagem Docker com a seguinte instrução:

```bash
docker build -t nomeImagem
```
Execute a imagem Docker em um container com a seguinte instrução:

```bash
docker run nomeImagem -u "http://exemplo.com" -r 10 -c 10
```
