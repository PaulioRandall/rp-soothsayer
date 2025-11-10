package validate

var observationSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
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
		"codes": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
	},
}

func ValidateObservation(observation JsonObject) []string {
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		observationSchema,
		observation,
		"observation",
	)

	return *pErrors
}
