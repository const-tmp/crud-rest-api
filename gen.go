package main

//go:generate go generate github.com/nullc4t/crud-rest-api/api
//go:generate oapi-codegen -package crud -o crud/crud.gen.go api/build/openapi.yaml
