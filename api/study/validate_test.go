package study

// TODO
// Create tests by creating mock schemas and data to test.
// Copy from the entity test files.
// Once done, delete obsolete tests in the entity test files.

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func parseJson(jsonStr string) JsonObject {
	data := map[string]any{}
	e := json.Unmarshal([]byte(jsonStr), &data)

	if e != nil {
		panic(e)
	}

	return data
}

func requireErrors(t *testing.T, actErrors []string, expErrors ...string) {
	require.ElementsMatch(t, actErrors, expErrors)
}

func testValidate(schemaName string, schema JsonObject, data JsonValue) []string {
	pErrors, err := newErrorSlice()

	validateSchema(
		err,
		schema,
		data,
		schemaName,
	)

	return *pErrors
}

func Test_Validate_1(t *testing.T) {
	// GIVEN Valid string schema
	// WHEN passed valid data
	// THEN returns no errors

	schema := JsonObject{
		"type": "string",
	}

	data := `"Abc"`

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}

func Test_Validate_2(t *testing.T) {
	// GIVEN Valid number schema
	// WHEN passed valid data
	// THEN returns no errors

	schema := JsonObject{
		"type": "number",
	}

	data := float64(123)

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}

func Test_Validate_3(t *testing.T) {
	// GIVEN Valid schema
	// WHEN passed wrong data type
	// THEN return type error

	schema := JsonObject{
		"type": "string",
	}

	data := float64(123)

	errors := testValidate("test", schema, data)
	requireErrors(t, errors,
		"test: Expected 'string' but got 'number'",
	)
}

func Test_Validate_4(t *testing.T) {
	// GIVEN Valid array schema
	// WHEN passed valid empty array data
	// THEN returns no errors

	schema := JsonObject{
		"type": "array",
		"items": JsonObject{
			"type": "string",
		},
	}

	data := []JsonValue{}

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}

func Test_Validate_5(t *testing.T) {
	// GIVEN Valid array schema
	// WHEN passed valid array data
	// THEN returns no errors

	schema := JsonObject{
		"type": "array",
		"items": JsonObject{
			"type": "string",
		},
	}

	data := []JsonValue{
		"Abc",
		"Xyz",
	}

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}

// TODO: Test for invalid array data type

func Test_Validate_10(t *testing.T) {
	// GIVEN Valid object schema with empty fields
	// WHEN validateSchema(...)
	// THEN returns no errors

	schema := JsonObject{
		"type":   "object",
		"fields": SchemaPropFields{},
	}

	data := parseJson(`{}`)

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}

func Test_Validate_11(t *testing.T) {
	// GIVEN Valid object schema with no 'fields' member
	// WHEN validateSchema(...)
	// THEN return missing field error

	schema := JsonObject{
		"type": "object",
	}

	data := parseJson(`{}`)

	errors := testValidate("test", schema, data)
	requireErrors(t, errors,
		"test: Schema with type 'object' must have a 'fields' property",
	)
}

func Test_Validate_12(t *testing.T) {
	// GIVEN Valid object schema with fields
	// WHEN validateSchema(...)
	// THEN returns no errors

	schema := JsonObject{
		"type": "object",
		"fields": SchemaPropFields{
			"name": JsonObject{
				"type": "string",
			},
			"age": JsonObject{
				"type": "number",
			},
		},
	}

	data := parseJson(`{
		"name": "Paul",
		"age": 37
	}`)

	errors := testValidate("test", schema, data)
	requireErrors(t, errors)
}
