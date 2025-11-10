package validate

import (
	"strings"
)

type reference struct {
	path     string
	segments []string
	field    string
}

func makeReference(path string) reference {
	parts := strings.Split(path, "@")
	field := ""

	if len(parts) > 1 {
		field = parts[1]
	}

	segments := strings.Split(parts[0], ".")

	return reference{
		path:     path,
		segments: segments,
		field:    field,
	}
}

func (ref reference) existsWithin(data JsonObject, v string) bool {
	refValue := ref.findArray(data)

	if refValue == nil {
		return false
	}

	array, ok := refValue.(JsonArray)

	if !ok {
		return false
	}

	if ref.field != "" {
		return objectArrayContains(array, ref.field, v)
	}

	return arrayContains(array, v)
}

func (ref reference) findArray(data JsonObject) JsonValue {
	var result JsonValue = JsonValue(data)

	for _, segment := range ref.segments {
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

func objectArrayContains(haystack JsonArray, field string, needle string) bool {
	for _, item := range haystack {
		itemObj, ok := item.(JsonObject)
		if ok && itemObj[field] == needle {
			return true
		}
	}
	return false
}

func arrayContains(haystack JsonArray, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
