# crud-rest-api
REST API for Generic CRUD operations and OpenAPI 3 specification. Library and codegen tool. 
# Installation
```
go install github.com/nullc4t/crud-rest-api@latest
```
# Prerequisites
- https://github.com/deepmap/oapi-codegen
- https://www.npmjs.com/package/swagger-cli
# Usage
## Bootstrap
First, we need some bootstrap code
```
crud-rest-api bootstrap
```
This will generate files:
```
api
├── build
├── gen.go
└── src
    ├── openapi.yaml
    ├── parameters
    │   ├── _index.yaml
    │   ├── path
    │   │   └── id.yaml
    │   └── query
    │       ├── limit.yaml
    │       ├── offset.yaml
    │       └── sort.yaml
    ├── paths
    ├── responses
    │   ├── BadRequest.yaml
    │   ├── Created.yaml
    │   ├── NoContent.yaml
    │   ├── NotFound.yaml
    │   ├── NullResponse.yaml
    │   ├── OK.yaml
    │   ├── Unauthorized.yaml
    │   ├── UnexpectedError.yaml
    │   └── _index.yaml
    └── schemas
        ├── Any.yaml
        ├── AnyNull.yaml
        ├── BaseModel.yaml
        ├── Error.yaml
        └── _index.yaml
```
## Generate
```
crud-rest-api generate example-config.yaml
```
TODO: see [gen.go](gen.go) for example
## Implement server
See [server implementation](internal/server/impl.go) and [run example](example/main.go)
