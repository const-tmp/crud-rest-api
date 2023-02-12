package paths

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func Put(dot generator.ResourceTemplate) (io.Reader, error) {
	return generator.RenderTemplate("put", `summary: Replace {{ .Schema }} if exists
tags:
  - {{ .Tag }}
description: Optional extended description in CommonMark or HTML
parameters:
  - $ref: '../../../parameters/path/id.yaml'
requestBody:
  required: true
  content:
    application/json:
      schema:
        $ref: '../../../schemas/{{ .Schema }}.yaml'
responses:
  200:
    $ref: '../../../responses/OK.yaml'
  401:
    $ref: '../../../responses/Unauthorized.yaml'
  404:
    $ref: '../../../responses/NotFound.yaml'
`, dot)
}
