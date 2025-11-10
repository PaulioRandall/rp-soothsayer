package validate

var mediaSectionSchema = JsonObject{
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
		"end": JsonObject{
			"type": "number",
		},
	},
}

func ValidateMediaSection(mediaSection JsonObject) []string {
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		mediaSectionSchema,
		mediaSection,
		"media_section",
	)

	return *pErrors
}
