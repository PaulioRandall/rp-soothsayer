package validate

import (
	"testing"
)

func Test_Observation_1(t *testing.T) {
	// GIVEN Valid observation
	// WHEN ValidateObservation()
	// THEN return no errors

	observation := parseJson(`{
		"id": "o1",
		"tags": ["P1", "T1"],
		"start": 123,
		"quote": "Oops!",
		"description": "Made mistake",
		"codes": ["c1"]
	}`)

	errors := ValidateObservation(observation)

	requireErrors(t, errors)
}

func Test_Observation_2(t *testing.T) {
	// GIVEN Observation missing all fields
	// WHEN Validating media section
	// THEN returns 7 type errors

	observation := parseJson(`{}`)

	errors := ValidateObservation(observation)

	requireErrors(t, errors,
		"observation.id: Expected 'string' but got 'unknown'",
		"observation.tags: Expected 'array' but got 'unknown'",
		"observation.start: Expected 'number' but got 'unknown'",
		"observation.quote: Expected 'string' but got 'unknown'",
		"observation.description: Expected 'string' but got 'unknown'",
		"observation.codes: Expected 'array' but got 'unknown'",
	)
}
