package paths

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func Post(dot generator.ResourceTemplate) (io.Reader, error) {
	return generator.RenderTemplate("post", `summary: Create {{ .Schema }}
tags:
  - {{ .Tag }}
description: Optional extended description in CommonMark or HTML
requestBody:
  required: true
  content:
    application/json:
      schema:
        $ref: '../../schemas/{{ .Schema }}.yaml'
responses:
  201:
    $ref: '../../responses/Created.yaml'
  400:
    $ref: '../../responses/BadRequest.yaml'
  401:
    $ref: '../../responses/Unauthorized.yaml'
`, dot)
}
