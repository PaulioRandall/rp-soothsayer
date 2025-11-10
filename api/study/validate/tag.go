package validate

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
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		tagSchema,
		tag,
		"tag",
	)

	return *pErrors
}
