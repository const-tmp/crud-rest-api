package paths

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func Delete(dot generator.ResourceTemplate) (io.Reader, error) {
	return generator.RenderTemplate("delete", `summary: Delete {{ .Schema }} by ID
tags:
  - {{ .Tag }}
description: Optional extended description in CommonMark or HTML
parameters:
  - $ref: '../../../parameters/path/id.yaml'
responses:
  204:
    $ref: '../../../responses/NoContent.yaml'
  401:
    $ref: '../../../responses/Unauthorized.yaml'
  404:
    $ref: '../../../responses/NotFound.yaml'
`, dot)
}
