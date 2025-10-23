package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findArray_1(t *testing.T) {
	// GIVEN Data with array at top level
	// WHEN searching for that array
	// THEN returns it

	data := parseJson(`{
		"ids": ["abc", "123"]
	}`)

	act := findArray(data, "ids")
	exp := []JsonValue{"abc", "123"}
	require.EqualValues(t, exp, act)
}

func Test_findArray_2(t *testing.T) {
	// GIVEN Data with array not at top level
	// WHEN searching for that array
	// THEN returns it

	data := parseJson(`{
		"things": {
			"stuff": {
				"ids": ["abc", "123"]
			}
		}
	}`)

	act := findArray(data, "things.stuff.ids")
	exp := []JsonValue{"abc", "123"}
	require.EqualValues(t, exp, act)
}

func Test_findArray_3(t *testing.T) {
	// GIVEN Search ref not matching data structure
	// WHEN searching for referenced array
	// THEN returns nil

	data := parseJson(`{
		"things": {
			"thongs": {
				"ids": ["abc", "123"]
			}
		}
	}`)

	act := findArray(data, "things.stuff.ids")
	require.EqualValues(t, nil, act)
}

func Test_arrayContains_1(t *testing.T) {
	// GIVEN array contains "c3"
	// WHEN checking if "c3" is contained in array
	// THEN returns true

	data := []JsonValue{
		"a1",
		"b2",
		"c3",
		"d4",
	}

	contains := arrayContains(data, "c3")
	require.EqualValues(t, true, contains)
}

func Test_arrayContains_2(t *testing.T) {
	// GIVEN array contains "c3"
	// WHEN checking if "e5" is contained in array
	// THEN returns false

	data := []JsonValue{
		"a1",
		"b2",
		"c3",
		"d4",
	}

	contains := arrayContains(data, "e5")
	require.EqualValues(t, false, contains)
}

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
