package study

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
	},
}

func ValidateMedia(media JsonObject) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		mediaSchema,
		media,
		"media",
	)

	return *pErrors
}
