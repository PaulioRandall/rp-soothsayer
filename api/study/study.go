package study

import (
	schema "soothsayer/api/study/schema"
)

var studySchema = schema.JsonObject{
	"type": "object",
	"fields": schema.SchemaPropFields{
		"name": schema.JsonObject{
			"type": "string",
		},
		"filepath": schema.JsonObject{
			"type": "string",
		},
		"media": schema.JsonObject{
			"type":  "array",
			"items": mediaSchema,
		},
		"codes": schema.JsonObject{
			"type":  "array",
			"items": codeSchema,
		},
		"tags": schema.JsonObject{
			"type":  "array",
			"items": tagSchema,
		},
	},
}

func ValidateStudy(study schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		studySchema,
		study,
		"study",
	)

	return *pErrors
}
