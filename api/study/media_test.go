package study

import (
	"testing"
)

func Test_Media_1(t *testing.T) {
	// GIVEN Valid media with valid sections
	// WHEN Validating media
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
				"end": 200
			},
			{
				"id": "456",
				"name": "Xyz",
				"tags": ["P1", "T2"],
				"start": 200,
				"end": 400
			}
		],
		"observations": [
			{
				"id": "o1",
				"tags": ["P1", "T1"],
				"start": 33,
				"quote": "",
				"description": "Misclicked: hit 'clear' instead of 'submit'",
				"codes": ["c3"]
			},
			{
				"id": "o2",
				"tags": ["P1", "T1"],
				"start": 166,
				"quote": "I really like cheese.",
				"description": "Big grin while expressing love for cheese",
				"codes": []
			},
			{
				"id": "o3",
				"tags": ["P2", "T2"],
				"start": 222,
				"quote": "I don't like the colour of the background, it's not cheesy enough.",
				"description": "",
				"codes": ["c4"]
			}
		]
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors)
}

func Test_Media_2(t *testing.T) {
	// GIVEN Media with all missing values
	// WHEN Validating media
	// THEN returns 5 type errors

	media := parseJson(`{}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors,
		"media.id: Expected 'string' but got 'unknown'",
		"media.name: Expected 'string' but got 'unknown'",
		"media.tags: Expected 'array' but got 'unknown'",
		"media.filepath: Expected 'string' but got 'unknown'",
		"media.sections: Expected 'array' but got 'unknown'",
		"media.observations: Expected 'array' but got 'unknown'",
	)
}

func Test_Media_3(t *testing.T) {
	// GIVEN Valid media but with a section missing all values
	// WHEN Validating media
	// THEN returns 5 type errors

	media := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1"],
		"filepath": "/home/bob/studies/cheese/P1 - video.mp4",
		"sections": [
			{}
		],
		"observations": [
			{}
		]
	}`)

	errors := ValidateMedia(media)

	requireErrors(t, errors,
		"media.sections[0].id: Expected 'string' but got 'unknown'",
		"media.sections[0].name: Expected 'string' but got 'unknown'",
		"media.sections[0].tags: Expected 'array' but got 'unknown'",
		"media.sections[0].start: Expected 'number' but got 'unknown'",
		"media.sections[0].end: Expected 'number' but got 'unknown'",
		"media.observations[0].id: Expected 'string' but got 'unknown'",
		"media.observations[0].tags: Expected 'array' but got 'unknown'",
		"media.observations[0].start: Expected 'number' but got 'unknown'",
		"media.observations[0].quote: Expected 'string' but got 'unknown'",
		"media.observations[0].description: Expected 'string' but got 'unknown'",
		"media.observations[0].codes: Expected 'array' but got 'unknown'",
	)
}
