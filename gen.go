package crud_api

//go:generate go generate crud-api/api
//go:generate oapi-codegen -package crud -o crud/crud.gen.go api/build/openapi.yaml
