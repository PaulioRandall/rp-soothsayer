package study

import (
	"testing"
)

func Test_MediaSection_1(t *testing.T) {
	// GIVEN Valid media section
	// WHEN ValidateMediaSection()
	// THEN return no errors

	mediaSection := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": 0,
		"end": 100
	}`)

	errors := ValidateMediaSection(mediaSection)

	requireErrors(t, errors)
}

func Test_MediaSection_2(t *testing.T) {
	// GIVEN Media section with bad ID type
	// WHEN ValidateMediaSection()
	// THEN return a single ID type error

	mediaSection := parseJson(`{
		"id": 123,
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": 0,
		"end": 100
	}`)

	errors := ValidateMediaSection(mediaSection)

	requireErrors(t, errors,
		"media_section.id: Expected 'string' but got 'number'",
	)
}

func Test_MediaSection_3(t *testing.T) {
	// GIVEN Media section with no fields
	// WHEN ValidateMediaSection()
	// THEN return 5 type errors

	mediaSection := parseJson(`{}`)

	errors := ValidateMediaSection(mediaSection)

	requireErrors(t, errors,
		"media_section.id: Expected 'string' but got 'unknown'",
		"media_section.name: Expected 'string' but got 'unknown'",
		"media_section.tags: Expected 'array' but got 'unknown'",
		"media_section.start: Expected 'number' but got 'unknown'",
		"media_section.end: Expected 'number' but got 'unknown'",
	)
}
