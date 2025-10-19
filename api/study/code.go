package study

import (
	schema "soothsayer/api/study/schema"
)

var codeSchema = schema.JsonObject{
	"type": "object",
	"fields": schema.SchemaPropFields{
		"id": schema.JsonObject{
			"type": "string",
		},
		"name": schema.JsonObject{
			"type": "string",
		},
		"tags": schema.JsonObject{
			"type": "array",
			"items": schema.JsonObject{
				"type": "string",
			},
		},
		"description": schema.JsonObject{
			"type": "string",
		},
	},
}

func ValidateCode(code schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		codeSchema,
		code,
		"code",
	)

	return *pErrors
}
