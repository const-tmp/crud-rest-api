package paths

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func Get(dot generator.ResourceTemplate) (io.Reader, error) {
	return generator.RenderTemplate("get", `summary: Returns a list of {{ .Schema }}
tags:
  - {{ .Tag }}
description: Optional extended description in CommonMark or HTML
parameters:
  - $ref: '../../parameters/query/offset.yaml'
  - $ref: '../../parameters/query/limit.yaml'
  - $ref: '../../parameters/query/sort.yaml'
responses:
  200:
    description: A JSON array of {{ .Schema }}
    content:
      application/json:
        schema:
          type: array
          items:
            $ref: '../../schemas/{{ .Schema }}.yaml'
`, dot)
}
