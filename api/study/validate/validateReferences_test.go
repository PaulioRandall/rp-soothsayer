package validate

import (
	"testing"
)

func testValidateReferences(schemaName string, schema JsonObject, data JsonValue) []string {
	pErrors, err := NewErrorSlice()
	checkRef := NewRefChecker(data.(JsonObject))

	ValidateReferences(
		err,
		checkRef,
		schema,
		data,
		schemaName,
	)

	return *pErrors
}

func Test_ValidateReferences_1(t *testing.T) {
	// GIVEN Invalid schema, ref used with non-string type
	// WHEN validating references
	// THEN returns error

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"thingId": JsonObject{
				"type": "number",
				"ref":  "thingIds",
			},
			"thingIds": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "string",
				},
			},
		},
	}

	data := parseJson(`{
		"thingId": 456,
		"thingIds": ["123", "456", "789"]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors,
		"test.thingId: Properties that reference another may only be strings",
	)
}

func Test_ValidateReferences_2(t *testing.T) {
	// GIVEN Value not in reference array
	// WHEN validating references
	// THEN returns error

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"thingId": JsonObject{
				"type": "string",
				"ref":  "thingIds",
			},
			"thingIds": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "string",
				},
			},
		},
	}

	data := parseJson(`{
		"thingId": "abc",
		"thingIds": ["123", "456", "789"]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors,
		"test.thingId: Reference doesn't exist at 'thingIds'",
	)
}

func Test_ValidateReferences_3(t *testing.T) {
	// GIVEN Valid schema, data, and reference
	// WHEN validating references
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"thingId": JsonObject{
				"type": "string",
				"ref":  "thingIds",
			},
			"thingIds": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "string",
				},
			},
		},
	}

	data := parseJson(`{
		"thingId": "456",
		"thingIds": ["123", "456", "789"]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors)
}

func Test_ValidateReferences_4(t *testing.T) {
	// GIVEN ref src and dst within objects
	// WHEN validating references
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"things": JsonObject{
				"type": "object",
				"fields": SchemaPropFields{
					"thingId": JsonObject{
						"type": "string",
						"ref":  "stuff.thingIds",
					},
				},
			},
			"stuff": JsonObject{
				"type": "object",
				"fields": SchemaPropFields{
					"thingIds": JsonObject{
						"type": "array",
						"items": JsonObject{
							"type": "string",
						},
					},
				},
			},
		},
	}

	data := parseJson(`{
		"things": {
			"thingId": "456"
		},
		"stuff": { 
			"thingIds": ["123", "456", "789"]
		}
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors)
}

func Test_ValidateReferences_5(t *testing.T) {
	// GIVEN ref is an item within an array
	// WHEN validating references
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"things": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "object",
					"fields": SchemaPropFields{
						"thingId": JsonObject{
							"type": "string",
							"ref":  "thingIds",
						},
					},
				},
			},
			"thingIds": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "string",
				},
			},
		},
	}

	data := parseJson(`{
		"things": [
			{
				"thingId": "123"
			},
			{
				"thingId": "456"
			},
			{
				"thingId": "789"
			}
		],
		"thingIds": [
			"123",
			"456",
			"789"
		]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors)
}

func Test_ValidateReferences_6(t *testing.T) {
	// GIVEN an object array used as a target for references
	// WHEN validating a reference exists within the array
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"thingId1": JsonObject{
				"type": "string",
				"ref":  "things@id",
			},
			"thingId2": JsonObject{
				"type": "string",
				"ref":  "things@id",
			},
			"thingId3": JsonObject{
				"type": "string",
				"ref":  "things@id",
			},
			"things": JsonObject{
				"type": "array",
				"items": JsonObject{
					"type": "object",
					"fields": SchemaPropFields{
						"id": JsonObject{
							"type": "string",
						},
					},
				},
			},
		},
	}

	data := parseJson(`{
		"thingId1": "123",
		"thingId2": "456",
		"thingId3": "789",
		"things": [
			{
				"id": "123"
			},
			{
				"id": "456"
			},
			{
				"id": "789"
			}
		]
	}`)

	errors := testValidateReferences("test", schema, data)
	requireErrors(t, errors)
}
