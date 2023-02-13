package openapi

import (
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"io"
)

func Get(dot TemplateData) (io.Reader, error) {
	return generator.RenderTemplate("get", `summary: Return list of {{ .Schema }}
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

func Post(dot TemplateData) (io.Reader, error) {
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

func GetByID(dot TemplateData) (io.Reader, error) {
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

func Put(dot TemplateData) (io.Reader, error) {
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

func Patch(dot TemplateData) (io.Reader, error) {
	return generator.RenderTemplate("patch", `summary: Update specified fields of {{ .Schema }}
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

func Delete(dot TemplateData) (io.Reader, error) {
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
