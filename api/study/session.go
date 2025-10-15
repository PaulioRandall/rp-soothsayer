package study

var sessionSchema = JsonObject{
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
		"observations": JsonObject{
			"type":  "array",
			"items": observationSchema,
		},
	},
}

func ValidateSession(session JsonObject) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		sessionSchema,
		session,
		"session",
	)

	return *pErrors
}
