package validate

var codeSchema = JsonObject{
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
		"description": JsonObject{
			"type": "string",
		},
	},
}

func ValidateCode(code JsonObject) []string {
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		codeSchema,
		code,
		"code",
	)

	return *pErrors
}
