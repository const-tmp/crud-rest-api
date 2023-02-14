package bootstrap

var Schemas = []OpenAPI3Source{
	{
		Path: "api/src/schemas/_index.yaml",
		Content: `Any:
    $ref: ./Any.yaml
AnyNull:
    $ref: ./AnyNull.yaml
BaseModel:
    $ref: ./BaseModel.yaml
Error:
    $ref: ./Error.yaml
`,
	},
	{
		Path:    "api/src/schemas/Any.yaml",
		Content: `description: Can be any value - string, number, boolean, array or object.`,
	},
	{
		Path:    "api/src/schemas/AnyNull.yaml",
		Content: `description: Can be any value - string, number, boolean, array, object or null.`,
	},
	{
		Path: "api/src/schemas/BaseModel.yaml",
		Content: `type: object
required:
  - id
  - created_at
  - updated_at
  - deleted_at
properties:
  id:
    type: integer
    format: uint64
    readOnly: true
  created_at:
    type: string
    format: date-time
    readOnly: true
  updated_at:
    type: string
    format: date-time
    readOnly: true
  deleted_at:
    type: string
    format: date-time
    readOnly: true
`,
	},
	{
		Path: "api/src/schemas/Error.yaml",
		Content: `type: object
required:
  - code
  - message
properties:
  code:
    type: integer
    format: int32
  message:
    type: string
`,
	},
}
