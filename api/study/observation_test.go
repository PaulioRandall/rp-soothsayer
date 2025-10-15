package study

import (
	"testing"
)

func Test_Observation_1(t *testing.T) {
	// GIVEN Valid observation
	// WHEN ValidateObservation()
	// THEN return no errors

	observation := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": 123,
		"quote": "Oops!",
		"description": "Made mistake"
	}`)

	errors := ValidateObservation(observation)

	requireErrors(t, errors)
}

func Test_Observation_2(t *testing.T) {
	// GIVEN Observation with missing description
	// WHEN ValidateObservation()
	// THEN return a single ID type error

	observation := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": 0,
		"quote": ""
	}`)

	errors := ValidateObservation(observation)

	requireErrors(t, errors,
		"observation.description: Expected 'string' but got 'unknown'",
	)
}

func Test_Observation_3(t *testing.T) {
	// GIVEN Observation with bad 'start' and 'quote' types.
	// WHEN ValidateObservation()
	// THEN return 2 type errors

	observation := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": "123",
		"quote": 123,
		"description": "Made mistake"
	}`)

	errors := ValidateObservation(observation)

	requireErrors(t, errors,
		"observation.start: Expected 'number' but got 'string'",
		"observation.quote: Expected 'string' but got 'number'",
	)
}
