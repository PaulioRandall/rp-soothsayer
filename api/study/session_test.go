package study

import (
	"testing"
)

func Test_Session_1(t *testing.T) {
	// GIVEN Valid session
	// WHEN ValidateSession()
	// THEN returns no errors

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"observations": []
	}`)

	errors := ValidateSession(media)

	requireErrors(t, errors)
}

func Test_Session_2(t *testing.T) {
	// GIVEN Session with observations
	// WHEN ValidateSession()
	// THEN return no errors

	session := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "P2"],
		"observations": [
			{
				"id": "123",
				"tags": ["P1", "T1"],
				"start": 123,
				"quote": "Oops!",
				"description": "Made mistake"
			},
			{
				"id": "456",
				"tags": ["P2", "T2"],
				"start": 456,
				"quote": "Yey!",
				"description": "Figured out shortcut to completing task"
			}
		]
	}`)

	errors := ValidateSession(session)

	requireErrors(t, errors)
}

func Test_Session_3(t *testing.T) {
	// GIVEN Session with observations and 2nd observation is missing 'start'
	// WHEN ValidateSession()
	// THEN returns 'observation[1].start' type error

	session := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "P2"],
		"observations": [
			{
				"id": "123",
				"tags": ["P1", "T1"],
				"start": 123,
				"quote": "Oops!",
				"description": "Made mistake"
			},
			{
				"id": "456",
				"tags": ["P2", "T2"],
				"start": "456",
				"quote": "Yey!",
				"description": "Figured out shortcut to completing task"
			}
		]
	}`)

	errors := ValidateSession(session)

	requireErrors(t, errors,
		"session.observations[1].start: Expected 'number' but got 'string'",
	)
}
