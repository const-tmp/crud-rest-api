package main

//go:generate go generate api/auth/api
//go:generate mkdir -p ../../pkg/auth
//go:generate oapi-codegen -package auth -o ../../pkg/auth/auth.gen.go api/build/openapi.yaml
