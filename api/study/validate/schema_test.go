package validate

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

	errors := Validate(mediaSchema, media, "media")

	requireErrors(t, errors)
}

func Test_Media_2(t *testing.T) {
	// GIVEN Media with all missing values
	// WHEN Validating media
	// THEN returns 5 type errors

	media := parseJson(`{}`)

	errors := Validate(mediaSchema, media, "media")

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

	errors := Validate(mediaSchema, media, "media")

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

func Test_MediaSection_1(t *testing.T) {
	// GIVEN Valid media section
	// WHEN Validating media section
	// THEN return no errors

	mediaSection := parseJson(`{
		"id": "123",
		"name": "Abc",
		"tags": ["P1", "T1"],
		"start": 0,
		"end": 100
	}`)

	errors := Validate(mediaSectionSchema, mediaSection, "media_section")

	requireErrors(t, errors)
}

func Test_MediaSection_2(t *testing.T) {
	// GIVEN Media section missing all values
	// WHEN Validating media section
	// THEN return 5 type errors

	mediaSection := parseJson(`{}`)

	errors := Validate(mediaSectionSchema, mediaSection, "media_section")

	requireErrors(t, errors,
		"media_section.id: Expected 'string' but got 'unknown'",
		"media_section.name: Expected 'string' but got 'unknown'",
		"media_section.tags: Expected 'array' but got 'unknown'",
		"media_section.start: Expected 'number' but got 'unknown'",
		"media_section.end: Expected 'number' but got 'unknown'",
	)
}

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

	errors := Validate(observationSchema, observation, "observation")

	requireErrors(t, errors)
}

func Test_Observation_2(t *testing.T) {
	// GIVEN Observation missing all fields
	// WHEN Validating media section
	// THEN returns 7 type errors

	observation := parseJson(`{}`)

	errors := Validate(observationSchema, observation, "observation")

	requireErrors(t, errors,
		"observation.id: Expected 'string' but got 'unknown'",
		"observation.tags: Expected 'array' but got 'unknown'",
		"observation.start: Expected 'number' but got 'unknown'",
		"observation.quote: Expected 'string' but got 'unknown'",
		"observation.description: Expected 'string' but got 'unknown'",
		"observation.codes: Expected 'array' but got 'unknown'",
	)
}

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

	errors := Validate(codeSchema, code, "code")

	requireErrors(t, errors)
}

func Test_Code_2(t *testing.T) {
	// GIVEN Code with all missing values
	// WHEN Code is validated
	// THEN return 4 type errors

	code := parseJson(`{}`)

	errors := Validate(codeSchema, code, "code")

	requireErrors(t, errors,
		"code.id: Expected 'string' but got 'unknown'",
		"code.name: Expected 'string' but got 'unknown'",
		"code.tags: Expected 'array' but got 'unknown'",
		"code.description: Expected 'string' but got 'unknown'",
	)
}

func Test_Tag_1(t *testing.T) {
	// GIVEN Valid tag
	// WHEN Validating tag
	// THEN returns no errors

	tag := parseJson(`{
		"id": "123",
		"name": "P1",
		"description": "Participant 1"
	}`)

	errors := Validate(tagSchema, tag, "tag")

	requireErrors(t, errors)
}

func Test_Tag_2(t *testing.T) {
	// GIVEN Tag missing all fields
	// WHEN Validating tag
	// THEN return 3 type errors

	tag := parseJson(`{}`)

	errors := Validate(tagSchema, tag, "tag")

	requireErrors(t, errors,
		"tag.id: Expected 'string' but got 'unknown'",
		"tag.name: Expected 'string' but got 'unknown'",
		"tag.description: Expected 'string' but got 'unknown'",
	)
}
