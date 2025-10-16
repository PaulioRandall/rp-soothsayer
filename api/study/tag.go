package study

var tagSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"description": JsonObject{
			"type": "string",
		},
	},
}

func ValidateTag(tag JsonObject) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		tagSchema,
		tag,
		"tag",
	)

	return *pErrors
}
