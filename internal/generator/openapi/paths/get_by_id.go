package paths

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func GetByID(dot generator.ResourceTemplate) (io.Reader, error) {
	return generator.RenderTemplate("get_by_id", `summary: Get {{ .Schema }} by ID
tags:
  - {{ .Tag }}
description: Optional extended description in CommonMark or HTML
parameters:
  - $ref: '../../../parameters/path/id.yaml'
responses:
  200:
    description: {{ .Schema }} object
    content:
      application/json:
        schema:
          $ref: '../../../schemas/{{ .Schema }}.yaml'
  401:
    $ref: '../../../responses/Unauthorized.yaml'
  404:
    $ref: '../../../responses/NotFound.yaml'
`, dot)
}
