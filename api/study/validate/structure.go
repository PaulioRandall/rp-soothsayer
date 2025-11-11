package validate

import (
	"fmt"
)

func validateStructure(
	err Err,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	propType := getPropValue[JsonType](schema, "type")

	typeMatch := validateStructureType(
		err,
		propType,
		prop,
		propName,
	)

	if !typeMatch {
		return
	}

	if propType == TypeObject {
		validateStructureFields(err, schema, prop.(JsonObject), propName)
	}

	if propType == TypeArray {
		validateStructureItems(err, schema, prop.(JsonArray), propName)
	}
}

func hasPropValue[R any](schema JsonObject, name string) bool {
	_, ok := schema[name].(R)
	return ok
}

func getPropValue[R any](schema JsonObject, name string) R {
	return schema[name].(R)
}

func validateStructureFields(
	err Err,
	schema JsonObject,
	prop JsonObject,
	propName string,
) {
	if !hasPropValue[SchemaPropFields](schema, "fields") {
		err("%s: Schema with type 'object' must have a 'fields' property", propName)
		return
	}

	fields := getPropValue[SchemaPropFields](schema, "fields")

	for name, subSchema := range fields {
		validateStructure(
			err,
			subSchema,
			prop[name],
			propName+"."+name,
		)
	}
}

func validateStructureItems(
	err Err,
	schema JsonObject,
	prop JsonArray,
	propName string,
) {
	itemSchema := getPropValue[JsonObject](schema, "items")

	for i, v := range prop {
		validateStructure(
			err,
			itemSchema,
			v,
			fmt.Sprintf("%s[%d]", propName, i),
		)
	}
}

func validateStructureType(
	err Err,
	expType JsonType,
	prop JsonValue,
	propName string,
) bool {
	actType := determineType(prop)

	if actType == expType {
		return true
	}

	err("%s: Expected '%s' but got '%s'", propName, expType, actType)
	return false
}
