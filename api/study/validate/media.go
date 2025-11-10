package validate

var mediaSchema = JsonObject{
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
		"filepath": JsonObject{
			"type": "string",
		},
		"sections": JsonObject{
			"type":  "array",
			"items": mediaSectionSchema,
		},
		"observations": JsonObject{
			"type":  "array",
			"items": observationSchema,
		},
	},
}

func ValidateMedia(media JsonObject) []string {
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		mediaSchema,
		media,
		"media",
	)

	return *pErrors
}
