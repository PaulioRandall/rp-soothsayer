package validate

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
		"tags": JsonObject{
			"type":  "array",
			"items": tagSchema,
		},
	},
}

func ValidateStudy(study JsonObject) []string {
	pErrors, err := NewErrorSlice()

	ValidateStructure(
		err,
		studySchema,
		study,
		"study",
	)

	return *pErrors
}

func MakeStudy() {

}
