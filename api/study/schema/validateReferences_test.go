package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testValidateReferences(schemaName string, schema JsonObject, data JsonValue) []string {
	pErrors, err := NewErrorSlice()

	ValidateReferences(
		err,
		data.(JsonObject),
		schema,
		data,
		schemaName,
	)

	return *pErrors
}

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

/*
func Test_ValidateReferences_1(t *testing.T) {
	// GIVEN Valid
	// WHEN passed valid data
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
		"thingId": "12",
		"thingIds": ["123"]
	}`)

	errors := testValidateReferences("test", schema, data)
	t.Fail()
	requireErrors(t, errors)
}
*/
