package bootstrap

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"strings"
)

type OpenAPI3Source struct {
	Path, Content string
}

var (
	Spec = OpenAPI3Source{
		Path: "api/src/openapi.yaml",
		Content: `openapi: 3.0.3
info:
    title: CRUD API
    description: CRUD API
    version: 1.0.0
servers:
    - url: http://localhost:8080
      description: Dev
paths:

components:
    parameters:
        $ref: ./parameters/_index.yaml
    responses:
        $ref: ./responses/_index.yaml
    schemas:
        $ref: ./schemas/_index.yaml
`,
	}

	Gen = OpenAPI3Source{
		Path: "api/gen.go",
		Content: `package api

//go:generate swagger-cli bundle src/openapi.yaml --outfile build/openapi.yaml --type yaml
`,
	}
)

func Files(path string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	var src []OpenAPI3Source

	src = append(src, Parameters...)
	src = append(src, Schemas...)
	src = append(src, Responses...)
	src = append(src, Spec, Gen)

	for _, source := range src {
		if err := generator.WriteFile(path+source.Path, strings.NewReader(source.Content)); err != nil {
			return err
		}
	}

	return nil
}
