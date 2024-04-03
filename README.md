```markdown
# API de Cotação de Frete

Este é um projeto Go que implementa uma API de cotação de frete fictícia, conforme especificado no desafio.

## Requisitos

- Go (versão 1.11 ou posterior)
- Docker (opcional)

## Configuração do Ambiente

1. Certifique-se de ter Go instalado em seu sistema. Você pode baixar e instalar a partir do site oficial: https://golang.org/dl/
2. Opcionalmente, se você quiser executar o projeto em um container Docker, certifique-se de ter o Docker instalado em seu sistema: https://www.docker.com/get-started

## Executando o Projeto

1. Clone este repositório para o seu ambiente local:

```
git clone https://github.com/seu_usuario/nome_do_repositorio.git
cd nome_do_repositorio
```

2. Inicie o servidor executando o seguinte comando:

```
go run main.go
```

3. O servidor estará sendo executado em http://localhost:8080.

## Rotas da API

### 1. POST /quote

Esta rota permite que você envie uma solicitação de cotação de frete.

**Exemplo de Corpo da Solicitação:**

```json
{
   "recipient":{
      "address":{
         "zipcode":"01311000"
      }
   },
   "volumes":[
      {
         "category":7,
         "amount":1,
         "unitary_weight":5,
         "price":349,
         "sku":"abc-teste-123",
         "height":0.2,
         "width":0.2,
         "length":0.2
      },
      {
         "category":7,
         "amount":2,
         "unitary_weight":4,
         "price":556,
         "sku":"abc-teste-527",
         "height":0.4,
         "width":0.6,
         "length":0.15
      }
   ]
}
```

### 2. GET /metrics?last_quotes={?}

Esta rota permite consultar métricas das cotações armazenadas no banco de dados.

**Parâmetros:**
- `last_quotes`: (opcional) Número de cotações para retornar.

## Estrutura do Projeto

- `main.go`: Arquivo principal que inicia o servidor e configura as rotas.
- `handlers.go`: Contém os manipuladores de rota HTTP.
- `quote.go`: Contém a lógica para cotação de frete.
- `database.go`: Contém a lógica para armazenamento em banco de dados.

## Considerações Finais

Este é um projeto de exemplo para demonstrar como construir uma API em Go para cotação de frete. Sinta-se à vontade para explorar e expandir este projeto conforme necessário.
```

Esse README.md fornece uma introdução ao projeto, explicando como executá-lo, as rotas disponíveis na API, a estrutura do projeto e algumas considerações finais. Certifique-se de personalizar as seções com detalhes específicos do seu projeto, como links para o repositório, instruções detalhadas de configuração e informações adicionais sobre as rotas da API.