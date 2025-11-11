package validate

import (
	"testing"
)

func testValidateStructure(schema JsonObject, data JsonValue, schemaName string) []string {
	pErrors, err := newErrorSlice()

	validateStructure(
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

	errors := testValidateStructure(schema, data, "test")
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

	errors := testValidateStructure(schema, data, "test")
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

	errors := testValidateStructure(schema, data, "test")
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

	errors := testValidateStructure(schema, data, "test")
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

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors)
}

func Test_Validate_6(t *testing.T) {
	// GIVEN Valid array schema
	// WHEN invalid data type used in array
	// THEN return type error

	schema := JsonObject{
		"type": "array",
		"items": JsonObject{
			"type": "string",
		},
	}

	data := []JsonValue{
		"Abc",
		float64(123),
	}

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors,
		"test[1]: Expected 'string' but got 'number'",
	)
}

func Test_Validate_7(t *testing.T) {
	// GIVEN Valid array within array schema
	// WHEN passed valid data
	// THEN returns no errors

	schema := JsonObject{
		"type": "array",
		"items": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
	}

	data := []JsonValue{
		[]JsonValue{
			"A",
			"B",
		},
		[]JsonValue{
			"C",
			"D",
		},
	}

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors)
}

func Test_Validate_8(t *testing.T) {
	// GIVEN Valid object schema with empty fields
	// WHEN validateSchema(...)
	// THEN returns no errors

	schema := JsonObject{
		"type":   "object",
		"fields": SchemaPropFields{},
	}

	data := parseJson(`{}`)

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors)
}

func Test_Validate_9(t *testing.T) {
	// GIVEN Valid object schema with no 'fields' member
	// WHEN validateSchema(...)
	// THEN return missing field error

	schema := JsonObject{
		"type": "object",
	}

	data := parseJson(`{}`)

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors,
		"test: Schema with type 'object' must have a 'fields' property",
	)
}

func Test_Validate_10(t *testing.T) {
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

	errors := testValidateStructure(schema, data, "test")
	requireErrors(t, errors)
}
