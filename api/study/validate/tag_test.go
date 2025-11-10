package validate

import (
	"testing"
)

func Test_Tag_1(t *testing.T) {
	// GIVEN Valid tag
	// WHEN Validating tag
	// THEN returns no errors

	tag := parseJson(`{
		"id": "123",
		"name": "P1",
		"description": "Participant 1"
	}`)

	errors := ValidateTag(tag)

	requireErrors(t, errors)
}

func Test_Tag_2(t *testing.T) {
	// GIVEN Tag missing all fields
	// WHEN Validating tag
	// THEN return 3 type errors

	tag := parseJson(`{}`)

	errors := ValidateTag(tag)

	requireErrors(t, errors,
		"tag.id: Expected 'string' but got 'unknown'",
		"tag.name: Expected 'string' but got 'unknown'",
		"tag.description: Expected 'string' but got 'unknown'",
	)
}
