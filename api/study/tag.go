package study

import (
	schema "soothsayer/api/study/schema"
)

var tagSchema = schema.JsonObject{
	"type": "object",
	"fields": schema.SchemaPropFields{
		"id": schema.JsonObject{
			"type": "string",
		},
		"name": schema.JsonObject{
			"type": "string",
		},
		"description": schema.JsonObject{
			"type": "string",
		},
	},
}

func ValidateTag(tag schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		tagSchema,
		tag,
		"tag",
	)

	return *pErrors
}
