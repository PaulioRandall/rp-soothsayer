package validate

import (
	"testing"
)

func Test_Study_1(t *testing.T) {
	// GIVEN Valid new/empty study
	// WHEN Validating study
	// THEN return no errors

	study := parseJson(`{
		"name": "My new study",
		"filepath": "/home/bob/studies/cheese",
		"media": [],
		"codes": [],
		"tags": []
	}`)

	errors := ValidateStudy(study)

	requireErrors(t, errors)
}

func Test_Study_2(t *testing.T) {
	// GIVEN Study missing all fields
	// WHEN Validating study
	// THEN returns 6 type errors

	study := parseJson(`{}`)

	errors := ValidateStudy(study)

	requireErrors(t, errors,
		"study.name: Expected 'string' but got 'unknown'",
		"study.filepath: Expected 'string' but got 'unknown'",
		"study.media: Expected 'array' but got 'unknown'",
		"study.codes: Expected 'array' but got 'unknown'",
		"study.tags: Expected 'array' but got 'unknown'",
	)
}

func Test_Study_3(t *testing.T) {
	// GIVEN Valid populated study
	// WHEN Validating study
	// THEN returns no errors

	study := parseJson(`{
		"name": "My new study",
		"filepath": "/home/bob/studies/cheese",
		"media": [
			{
				"id": "m1",
				"name": "Media 1",
				"tags": ["P1"],
				"filepath": "/home/bob/studies/cheese/P1 - video.mp4",
				"sections": [
					{
						"id": "m1.ms1",
						"name": "Task 1",
						"tags": ["P1", "T1"],
						"start": 0,
						"end": 100
					},
					{
						"id": "m1.ms2",
						"name": "Task 2",
						"tags": ["P1", "T2"],
						"start": 100,
						"end": 200
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
			},
			{
				"id": "m2",
				"name": "Media 2",
				"tags": ["P2"],
				"filepath": "/home/bob/studies/cheese/P2 - video.mp4",
				"sections": [
					{
						"id": "m2.ms1",
						"name": "Task 1",
						"tags": ["P2", "T1"],
						"start": 0,
						"end": 150
					},
					{
						"id": "m2.ms2",
						"name": "Task 2",
						"tags": ["P2", "T2"],
						"start": 150,
						"end": 300
					}
				],
				"observations": [
					{
						"id": "m2.o1",
						"tags": ["P2", "T1"],
						"start": 33,
						"quote": "",
						"description": "Misclicked: hit 'clear' instead of 'submit'",
						"codes": ["c3"]
					},
					{
						"id": "m2.o2",
						"tags": ["P2", "T2"],
						"start": 166,
						"quote": "I really like cheese.",
						"description": "Big grin while expressing love for cheese",
						"codes": []
					}
				]
			}
		],
		"codes": [
			{
				"id": "c1",
				"name": "Unresponsive cursor (speed)",
				"tags": ["Cursor"],
				"description": "Annoyed at unresponsive cursor/pointer speed"
			},
			{
				"id": "c2",
				"name": "Sensitive cursor (speed)",
				"tags": ["Cursor"],
				"description": "Annoyed at sensitive cursor/pointer speed"
			},
			{
				"id": "c3",
				"name": "Misclick",
				"tags": ["Cursor"],
				"description": "Accidentally clicked an element whilst trying to click another"
			},
			{
				"id": "c4",
				"name": "Dislike colour scheme",
				"tags": [],
				"description": "Dislikes some aspect of colour scheme"
			}
		],
		"tags": [
			{
				"id": "t1",
				"name": "Cursor",
				"description": "Relating to the cursor/pointer, likely associated with mouse/trackpad"
			},
			{
				"id": "t2",
				"name": "P1",
				"description": "Involves participant 1"
			},
			{
				"id": "t3",
				"name": "P2",
				"description": "Involves participant 2"
			},
			{
				"id": "t4",
				"name": "T1",
				"description": "Associated with task 1"
			},
			{
				"id": "t5",
				"name": "T2",
				"description": "Associated with task 2"
			}
		]
	}`)

	errors := ValidateStudy(study)

	requireErrors(t, errors)
}
