package schema

import (
// "fmt"
)

func ValidateReferences(
	err Err,
	schema JsonObject,
	data JsonObject,
	prop JsonValue,
	propName string,
) {
	if hasPropValue[string](schema, "ref") {
		findAndCheckRef(err, schema, data, prop, propName)
		return
	}
}

func findAndCheckRef(
	err Err,
	schema JsonObject,
	data JsonObject,
	prop JsonValue,
	propName string,
) {
	if getPropValue[string](schema, "type") != TypeString {
		err("%s: Reference fields must always be a string", propName)
		return
	}

	ref := getPropValue[string](schema, "ref")

	if !refExists(data, prop.(string), ref) {
		err("%s: Reference fields must always be a string", propName)
	}
}

func refExists(
	data JsonObject,
	value string,
	ref string,
) bool {
	// TODO:
	//   1: split ref by '.'
	//   2: descend 'data' to find ref
	//   3: if can't descend to end of ref then return false
	//   4: if id/value found that matches 'ref' and 'value' then return true
	//   5: else return false

	return false
}
