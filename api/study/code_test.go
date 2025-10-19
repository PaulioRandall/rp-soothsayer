package study

import (
	"testing"
)

func Test_Code_1(t *testing.T) {
	// GIVEN Valid code
	// WHEN Code is validated
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
	// GIVEN Code with all missing values
	// WHEN Code is validated
	// THEN return 4 type errors

	code := parseJson(`{}`)

	errors := ValidateCode(code)

	requireErrors(t, errors,
		"code.id: Expected 'string' but got 'unknown'",
		"code.name: Expected 'string' but got 'unknown'",
		"code.tags: Expected 'array' but got 'unknown'",
		"code.description: Expected 'string' but got 'unknown'",
	)
}
