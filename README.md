## Go United - API Gateway

### Generate code for REST API

#### Oapi-codegen

1) Install [oapi-codegen](https://github.com/deepmap/oapi-codegen)
```
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```
2) Command to run code generator:
```
$ oapi-codegen -generate types,server -package rest api/openapi/openapi.yaml > internal/api/rest/rest.gen.go
```

#### Swagger codegen
```
$ wget https://raw.githubusercontent.com/swagger-api/swagger-codegen/master/standalone-gen-dev/run-in-docker.sh -O swagger-codegen.sh
```


### Run containers for dev environment
```
$ sudo docker compose -f ./compose/docker-compose.dev.yml build
$ sudo docker compose -f ./compose/docker-compose.dev.yml -p gounited up -d
```