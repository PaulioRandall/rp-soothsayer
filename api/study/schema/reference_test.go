package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reference_findArray_1(t *testing.T) {
	// GIVEN Data with array at top level
	// WHEN searching for that array
	// THEN returns it

	data := parseJson(`{
		"ids": ["abc", "123"]
	}`)

	ref := makeReference("ids")
	act := ref.findArray(data)
	exp := []JsonValue{"abc", "123"}
	require.EqualValues(t, exp, act)
}

func Test_reference_findArray_2(t *testing.T) {
	// GIVEN Data with array not at top level
	// WHEN searching for that array
	// THEN returns it

	data := parseJson(`{
		"things": {
			"stuff": {
				"ids": ["abc", "123"]
			}
		}
	}`)

	ref := makeReference("things.stuff.ids")
	act := ref.findArray(data)
	exp := []JsonValue{"abc", "123"}
	require.EqualValues(t, exp, act)
}

func Test_reference_findArray_3(t *testing.T) {
	// GIVEN Search ref not matching data structure
	// WHEN searching for referenced array
	// THEN returns nil

	data := parseJson(`{
		"things": {
			"thongs": {
				"ids": ["abc", "123"]
			}
		}
	}`)

	ref := makeReference("things.stuff.ids")
	act := ref.findArray(data)
	require.EqualValues(t, nil, act)
}
