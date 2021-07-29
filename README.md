# Blockchain API

A REST API to get blockchain data

## Running

- **Option 1**: to run application asap you can use docker-compose, as the images have already been built and uploaded to docker hub:
```shell
    make docker_compose
```

- **Option 2**: to run the application using go:
```shell
    make first_run
```

## Documentation

- Swagger as the specification to document the API
- After running the application, you can check the webpage at:
    - http://localhost:8080/swagger/index.html

## Project's structure

```shell script
📦Project
 ┣ 📂cmd
 ┃ ┗ 📂rest
 ┃ ┃ ┣ 📂handler
 ┃ ┃ ┃ ┣ 📂block
 ┃ ┃ ┃ ┣ 📂transaction
 ┃ ┃ ┃ ┗ 📜health_check_handler.go
 ┃ ┃ ┣ 📂presenter
 ┃ ┃ ┗ 📜api.go
 ┣ 📂docs
 ┣ 📂internal
 ┃ ┣ 📂application
 ┃ ┃ ┣ 📂adapter
 ┃ ┃ ┃ ┗ 📂so-chain
 ┃ ┃ ┣ 📂mock
 ┃ ┃ ┣ 📜service.go
 ┃ ┣ 📂domain
 ┃ ┃ ┣ 📂mock
 ┃ ┗ 📂infra
 ┃ ┃ ┣ 📂config
 ┃ ┃ ┗ 📂so-chain
 ┃ ┃ ┃ ┣ 📂mock
 ┗ 📜main.go
```

## Code quality
- **golangci-lint**: 0 lint errors
- **unit testing**: 71.1% of code coverage

## Decisions

- **Design**: Clean architecture
- **Unit testing**: the libraries onsi/gomega and golang/mock are being used because their facility and popularity
- **Host application**: AWS ECS or AWS EKS as they are the most common and reliable solutions
- **Documentation**: created Swagger using swaggo and hosted on https://{host}/swagger
- **Scaling**:
  Using ECS or EKS, they will do this job for us, so the only concern is about:
    - Vertically scaling: set a minimum/maximum of CPU/RAM resources
    - Horizontally scaling set a minimum/maximum of replicas and specify the condition using such as: CPU/RAM and/or requests
- **Security**: it is a large topic (the CIA Triad), but I can think about:
    - Depending on the client:
        - Api-Key
        - Authentication using OAuth, like AWS Cognito or Keycloack
    - Rate limit
    - Container scan
    - SSL
    - Good practices in development such as coding scan, CI/CD, permissions
- **Deal with latency and delays in external API calls**: for this API is ok just to set a timeout such as 10s as we have only fetch calls,
  but some improvements are:
    - for a better time response, a cache could be included for subsequent calls
    - if the API will deal with operations, so an idempotency key would be needed
- **Caching**: in order to have an optimised solution, a dedicated cache solution would be needed, like Redis
- **Errors**: the logging is catching all errors and for the client can vary on response:
    - if is an internal error only, respond with a status code 500 without details
    - if is a business domain error, respond with a status code 4xx
- **Monitoring**: some solutions for this purpose: Dynatrace/Splunk/Elastic Stack/Slack
- **Observability**: some solutions for this purpose: Splunk or Elastic Stack
    - for logging, the application need:
        - standardization of the logs, like putting them in a json format
        - track more data about request

## TODO:

- Improve test coverage
- Finalize the implementation of the API making it RESTful, using HATEOAS as an example