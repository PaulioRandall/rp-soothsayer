package study

var observationSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"start": JsonObject{
			"type": "number",
		},
		"quote": JsonObject{
			"type": "string",
		},
		"description": JsonObject{
			"type": "string",
		},
	},
}

func ValidateObservation(observation JsonObject) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		observationSchema,
		observation,
		"observation",
	)

	return *pErrors
}
