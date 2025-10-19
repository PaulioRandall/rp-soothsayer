package schema

import (
	"testing"
)

func testValidateReferences(schemaName string, schema JsonObject, data JsonValue) []string {
	pErrors, err := NewErrorSlice()

	ValidateReferences(
		err,
		schema,
		data.(JsonObject),
		data,
		schemaName,
	)

	return *pErrors
}

func Test_ValidateReferences_1(t *testing.T) {
	// GIVEN Valid
	// WHEN passed valid data
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"thingId": JsonObject{
				"type": "string",
				"ref":  "things",
			},
			"things": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "string",
				},
			},
		},
	}

	data := parseJson(`{
		"thingId": "123",
		"things": ["123"]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors)
}
