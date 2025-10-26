package study

import (
	schema "soothsayer/api/study/schema"
)

var observationSchema = schema.JsonObject{
	"type": "object",
	"fields": schema.SchemaPropFields{
		"id": schema.JsonObject{
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
		"quote": schema.JsonObject{
			"type": "string",
		},
		"description": schema.JsonObject{
			"type": "string",
		},
		"codes": schema.JsonObject{
			"type": "array",
			"items": schema.JsonObject{
				"type": "string",
			},
		},
	},
}

func ValidateObservation(observation schema.JsonObject) []string {
	pErrors, err := schema.NewErrorSlice()

	schema.ValidateStructure(
		err,
		observationSchema,
		observation,
		"observation",
	)

	return *pErrors
}
