package study

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
	for _, e := range expErrors {
		require.Contains(t, actErrors, e)
	}

	require.Equal(t, len(actErrors), len(expErrors))
}

func Test_1(t *testing.T) {
	// GIVEN Valid media section
	// WHEN ValidateMediaSection()
	// THEN return no errors

	mediaSection := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1, T1"],
		"start": 0,
		"end": 100
	}`)

	errors := ValidateMediaSection(mediaSection)

	requireErrors(t, errors)
}

func Test_2(t *testing.T) {
	// GIVEN Media section with bad ID type
	// WHEN ValidateMediaSection()
	// THEN return a single ID type error

	mediaSection := parseJson(`{
		"id": 123,
		"name": "Abc",
		"tags": ["P1, T1"],
		"start": 0,
		"end": 100
	}`)

	errors := ValidateMediaSection(mediaSection)

	requireErrors(t, errors,
		"media_section.id: Expected 'string' but got 'number'",
	)
}

func Test_3(t *testing.T) {
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
