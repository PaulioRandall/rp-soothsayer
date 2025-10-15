package study

import (
	"testing"
)

func Test_Code_1(t *testing.T) {
	// GIVEN Valid code
	// WHEN ValidateCode()
	// THEN returns no errors

	code := parseJson(`{
		"id": "123",
		"name": "Slow cursor speed",
		"tags": [],
		"description": "Annoyed at cursor/pointer speed"
	}`)

	errors := ValidateCode(code)

	requireErrors(t, errors)
}

func Test_Code_2(t *testing.T) {
	// GIVEN Code with missing description
	// WHEN ValidateCode()
	// THEN return 1 type error

	code := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"]
	}`)

	errors := ValidateCode(code)

	requireErrors(t, errors,
		"code.description: Expected 'string' but got 'unknown'",
	)
}
