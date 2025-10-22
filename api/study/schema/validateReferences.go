package schema

import (
	//"fmt"
	"strings"
)

func ValidateReferences(
	err Err,
	data JsonObject,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	if hasPropValue[string](schema, "ref") {

	}

	/*
		for fieldName, subSchema := range data["fields"] {
			fieldPropName := propName + "." + fieldName
			fieldValue := data[fieldName]

			if hasPropValue[string](subSchema, "ref") {
				findAndCheckRef(err, subSchema, data, fieldValue, fieldPropName)
				continue
			}

			if getPropValue[string](subSchema, "type") == TypeObject {
				ValidateReferences(err, subSchema, fieldValue, fieldPropName)
			}
		}
	*/
}

/*
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

		if determineType(prop) != TypeString {
			err("%s: Reference fields must always be a string", propName)
			return
		}

		ref := getPropValue[string](schema, "ref")

		if !refExists(data, prop.(string), ref) {
			err("%s: Reference doesn't exist at '%s'", propName, ref)
		}
	}

func refExists(

	data JsonObject,
	value string,
	ref string,

	) bool {
		path := strings.Split(ref, ".")
		pArray := getRefArray(data, path...)

		fmt.Printf("pArray: %s", *pArray)

		if pArray == nil {
			return false
		}

		array, ok := (*pArray).(JsonArray)

		if !ok {
			return false
		}

		return arrayContains(array, value)
	}
*/
func findArray(data JsonObject, ref string) JsonValue {
	var result JsonValue = JsonValue(data)
	path := strings.Split(ref, ".")

	for _, segment := range path {
		obj, isObject := result.(JsonObject)
		if !isObject {
			return nil
		}

		if v, ok := obj[segment]; ok {
			result = v
		} else {
			return nil
		}
	}

	return result
}

func arrayContains(haystack JsonArray, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
