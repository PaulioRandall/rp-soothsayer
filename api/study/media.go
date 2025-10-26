package study

import (
	schema "soothsayer/api/study/schema"
)

var mediaSchema = schema.JsonObject{
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
		"filepath": schema.JsonObject{
			"type": "string",
		},
		"sections": schema.JsonObject{
			"type":  "array",
			"items": mediaSectionSchema,
		},
		"observations": schema.JsonObject{
			"type":  "array",
			"items": observationSchema,
		},
	},
}

func ValidateMedia(media schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		mediaSchema,
		media,
		"media",
	)

	return *pErrors
}
