## Go United - API Gateway

### Generate code for REST API
1) Install [oapi-codegen](https://github.com/deepmap/oapi-codegen)
```
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```
2) Command to run code generator:
```
$ oapi-codegen -generate types,server -package rest api/openapi/openapi.yaml > internal/api/rest/rest.gen.go
```