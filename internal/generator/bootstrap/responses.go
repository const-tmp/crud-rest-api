package bootstrap

var Responses = []OpenAPI3Source{
	{
		Path: "api/src/responses/_index.yaml",
		Content: `NotFound:
  $ref: "./NotFound.yaml"
BadRequest:
  $ref: "./BadRequest.yaml"
Unauthorized:
  $ref: "./Unauthorized.yaml"
UnexpectedError:
  $ref: "./UnexpectedError.yaml"
NullResponse:
  $ref: "./NullResponse.yaml"
Created:
  $ref: "./Created.yaml"
OK:
  $ref: "./OK.yaml"
NoContent:
  $ref: "./NoContent.yaml"
`,
	},
	{
		Path: "api/src/responses/BadRequest.yaml",
		Content: `description: Bad request
content:
  application/json:
    schema:
      $ref : "../schemas/Error.yaml"
`,
	},
	{
		Path: "api/src/responses/Created.yaml",
		Content: `description: Created
content:
  application/json:
    schema:
      $ref : "../schemas/BaseModel.yaml"
`,
	},
	{
		Path:    "api/src/responses/NoContent.yaml",
		Content: `description: No content`,
	},
	{
		Path: "api/src/responses/NotFound.yaml",
		Content: `description: Resource not found
content:
  application/json:
    schema:
      $ref : "../schemas/Error.yaml"
`,
	},
	{
		Path:    "api/src/responses/NullResponse.yaml",
		Content: `description: Null response`,
	},
	{
		Path:    "api/src/responses/OK.yaml",
		Content: `description: OK`,
	},
	{
		Path: "api/src/responses/Unauthorized.yaml",
		Content: `description: Unauthorized
content:
  application/json:
    schema:
      $ref : "../schemas/Error.yaml"
`,
	},
	{
		Path: "api/src/responses/UnexpectedError.yaml",
		Content: `description: Unexpected error
content:
  application/json:
    schema:
      $ref : "../schemas/Error.yaml"
`,
	},
}
