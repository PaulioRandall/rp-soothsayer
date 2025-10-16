package study

import (
	"testing"
)

func Test_Tag_1(t *testing.T) {
	// GIVEN Valid tag
	// WHEN ValidateTag()
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
	// GIVEN Tag with missing description
	// WHEN ValidateTag()
	// THEN return 1 type error

	tag := parseJson(`{
		"id": "123",
		"name": "P1"
	}`)

	errors := ValidateTag(tag)

	requireErrors(t, errors,
		"tag.description: Expected 'string' but got 'unknown'",
	)
}
