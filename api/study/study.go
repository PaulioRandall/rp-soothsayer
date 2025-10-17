package study

var studySchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"name": JsonObject{
			"type": "string",
		},
		"filepath": JsonObject{
			"type": "string",
		},
		"media": JsonObject{
			"type":  "array",
			"items": mediaSchema,
		},
		"codes": JsonObject{
			"type":  "array",
			"items": codeSchema,
		},
		"observations": JsonObject{
			"type":  "array",
			"items": observationSchema,
		},
		"tags": JsonObject{
			"type":  "array",
			"items": tagSchema,
		},
	},
}

func ValidateStudy(study JsonObject) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		studySchema,
		study,
		"study",
	)

	return *pErrors
}
