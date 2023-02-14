package bootstrap

var Parameters = []OpenAPI3Source{
	{
		Path: "api/src/parameters/_index.yaml",
		Content: `# path
id:
  $ref: './path/id.yaml'

# query
limit:
  $ref: './query/limit.yaml'
offset:
  $ref: './query/offset.yaml'
sort:
  $ref: './query/sort.yaml'
`,
	},
	{
		Path: "api/src/parameters/path/id.yaml",
		Content: `name: id
in: path
required: true
description: The id of the resource to retrieve
schema:
  type: integer
  format: uint64
`,
	},
	{
		Path: "api/src/parameters/query/limit.yaml",
		Content: `name: limit
in: query
description: How many items to return at one time
required: false
schema:
  type: integer
  format: uint32
`,
	},
	{
		Path: "api/src/parameters/query/offset.yaml",
		Content: `name: offset
in: query
description: How many items to skip
required: false
schema:
  type: integer
  format: uint32
`,
	},
	{
		Path: "api/src/parameters/query/sort.yaml",
		Content: `name: sort
in: query
required: false
description: Sort order
schema:
  type: object
  additionalProperties:
    type: string
    enum: [ asc, desc ]
`,
	},
}
