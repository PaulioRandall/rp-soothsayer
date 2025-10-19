package study

import (
	schema "soothsayer/api/study/schema"
)

var mediaSectionSchema = schema.JsonObject{
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
		"start": schema.JsonObject{
			"type": "number",
		},
		"end": schema.JsonObject{
			"type": "number",
		},
	},
}

func ValidateMediaSection(mediaSection schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		mediaSectionSchema,
		mediaSection,
		"media_section",
	)

	return *pErrors
}
