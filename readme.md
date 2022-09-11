# Coinbit

This project is intended to be a test from coinbit to create a service where user can deposit money into their wallet and fetch their current balance using event-driven with golang.
![architecture-flow](/doc/image.png)

## Installation
This project require kafka and redis to run, but down worry, I make you a docker-compose. Just running `docker-compose up` and everything will be setup.

Install the dependencies and create env file from env.example.yaml then compile the apps
```sh
go mod download
cp env.example.yaml env.yaml
make build-api
make build-processor
```

Run apps with this command
```sh
chmod +x api
chmod +x processor
./api
./processor
```

## Route
Here are the list of the API

| METHOD |           ENDPOINT                           | BODY                           |
| ------ | -------------------------------------------- | ------------------------------ |
|  POST  | http://localhost:3000/v1/coinbit/deposit     |{"wallet_id":"1", "amount":1000}|
|  GET   | http://localhost:3000/v1/coinbit/balance/:id |                                |

> Funfact: if you requested with `X-Request-ID` included on header, it will be activated imdempotency for the request, so if you invoke again with same request ID it will be returned same result as before.


## Development

##### Updating Proto
The contract is stored in the `proto` folder. If you want to change the proto files, you need to compile it again with `make proto` but before that make sure `protoc` and `protoc-gen-go` (just install `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`) is installed on your machine. Please refer this https://grpc.io/docs/protoc-installation/

##### Testing
The testing already included on `domain` folder. I just doing unit testing for the domain, because all of the bussiness logic are implemented on domain, to run the test just `make test`

##### Develop the applications
There are two applications on this project first is `api` it's located on `infrastructure/api` and another on is `processor` it's located on `insfrastructure/processor`. If you want to run this application just `make run-api` or `make run-processor`

## Project Structure
This project implement clean architecture, the purposes is to organize a project more clearly and store business logic so that it remains independent, clean, and extensible.
There are 3 layers in this project:
- Domain: models of the business context will be implemented on this layer. There can also be methods, for example, for validation.
- Application: this layer contains of abstraction of the application and orchestrator of the business logic, like interface, config, and usecases.
- Infrastructure: it's outer layer which changes frequently based on the technologies, updates like database, frameworks, and third party technology. Also this layer acts as a communicator to convert data into desired format for storing into external sources like database, file system, 3rd parties, and convert data for use cases or business logic.
 
##### `application/config`
- configuration of the application from env is registered here
##### `application/external_services`
- interfaces to use external_service inside the application 
##### `application/usecases`
- usecases is an application logic to processing/orchestrating domain
##### `domain`
- domain is the smallest unit on this application, there is no dependency to another layer, bussiness logic is implemented here so it can be easily to create unit test
##### `insfrastructure/api`
- api server is created in this folder, api is an adapters that necessarily do the conversion of data in both ways, convert data for usecases or business logic
##### `insfrastructure/api`
- same like api, processor is also an adapters to handling data flow and convert data for usecases or business logic
##### `insfrastructure/external_service`
- third parties technologies stored here with interface implemented from application, it aims to make sure application running well even if we change the technologies.

# Demo
<video width="640" height="360" controls>
  <source src="(/doc/demo.mp4" type="video/mp4">
</video>