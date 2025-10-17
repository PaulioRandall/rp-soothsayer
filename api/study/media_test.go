package study

import (
	"testing"
)

func Test_Media_1(t *testing.T) {
	// GIVEN Valid media
	// WHEN ValidateMedia()
	// THEN returns no errors

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"filepath": "/home/bob/studies/cheese/P1 - video.mp4",
		"sections": []
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors)
}

func Test_Media_2(t *testing.T) {
	// GIVEN Media section with missing filepath
	// WHEN ValidateMedia()
	// THEN return 1 type error

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"sections": []
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors,
		"media.filepath: Expected 'string' but got 'unknown'",
	)
}

func Test_Media_3(t *testing.T) {
	// GIVEN Media with sections and 2nd section is missing 'end'
	// WHEN ValidateMedia()
	// THEN returns no errors

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"filepath": "/home/bob/studies/cheese/P1 - video.mp4",
		"sections": [
			{
				"id": "123",
				"name": "Abc",
				"tags": ["P1", "T1"],
				"start": 0,
				"end": 100
			},
			{
				"id": "456",
				"name": "Xyz",
				"tags": ["P1", "T2"],
				"start": 100,
				"end": 200
			}
		]
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors)
}

func Test_Media_4(t *testing.T) {
	// GIVEN Media with sections and 2nd section is missing 'end'
	// WHEN ValidateMedia()
	// THEN returns 1 type error

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"filepath": "/home/bob/studies/cheese/P1 - video.mp4",
		"sections": [
			{
				"id": "123",
				"name": "Abc",
				"tags": ["P1", "T1"],
				"start": 0,
				"end": 100
			},
			{
				"id": "456",
				"name": "Xyz",
				"tags": ["P1", "T2"],
				"start": 100
			}
		]
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors,
		"media.sections[1].end: Expected 'number' but got 'unknown'",
	)
}
