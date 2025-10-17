package study

import (
	"testing"
)

func Test_Study_1(t *testing.T) {
	// GIVEN Valid new/empty study
	// WHEN ValidateStudy()
	// THEN return no errors

	study := parseJson(`{
		"name": "My new study",
		"filepath": "/home/bob/studies/cheese",
		"media": [],
		"codes": [],
		"observations": [],
		"tags": []
	}`)

	errors := ValidateStudy(study)

	requireErrors(t, errors)
}

func Test_Study_2(t *testing.T) {
	// GIVEN Valid a populated study
	// WHEN ValidateStudy()
	// THEN return no errors

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
			}
		],
		"observations": [
			{
				"id": "o1",
				"tags": ["P1", "T1"],
				"start": 33,
				"quote": "",
				"description": "Misclicked: hit 'clear' instead of 'submit'"
			},
			{
				"id": "o2",
				"tags": ["P1", "T1"],
				"start": 66,
				"quote": "I really like cheese.",
				"description": "Big grin while expressing love for cheese"
			},
			{
				"id": "o3",
				"tags": ["P1", "T2"],
				"start": 149,
				"quote": "I don't like the colour of the background, it's not cheesy enough.",
				"description": ""
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
