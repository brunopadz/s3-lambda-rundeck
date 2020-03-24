# S3-lambda-rundeck

## 🇧🇷
Esse é um exemplo bem simples de como utilizar Lambda para realizar chamadas para webhooks do Rundeck a partir de eventos no S3.

Essa função será invocada toda vez que um determinado tipo de evento ocorrer em um bucket. Nesse exemplo, será enviado via POST um json contendo o nome do bucket e do objeto que foi adicionado. No caso, o job do Rundeck e o webhook deverão receber esses dois parâmetros.

### Iniciando

- Compile a aplicação utilizando `GOOS=linux`
- Faça [deploy](https://docs.aws.amazon.com/pt_br/lambda/latest/dg/golang-package.html) da sua função
- A aplicação espera uma variável de ambiente chamada `RUNDECK_JOB_URL`. O valor é a URL do webhook do job do Rundeck.

### Quer ajudar?

Eu não sou desenvolvedor Go, ainda estou aprendendo. Criei essa função com o pouco que eu sei. Então se deseja melhorar algo, só abrir um Pull Request ou criar uma issue. :)

## 🇬🇧
This is a really simple example on how to use a Lambda function to call Rundeck webhooks from S3 events.

This function will be invoked everytime an event occurs in a S3 bucket. In this example, a json containing the bucket name and the object key will be sent in a POST request.

## Getting Started

- Compile the app using `GOOS=linux`
- [Deploy it](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html).
- The app needs `RUNDECK_JOB_URL` environment variable, set the webhook URL as its value.

## Wanna contribute?

I'm not a Go dev and still learning it. If you wish to contribute, fix something or add some new feature, feel free to open a Pull Request or an issue. :)
