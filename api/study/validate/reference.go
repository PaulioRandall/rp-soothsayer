package validate

import (
	"fmt"
)

func validateReferences(
	err Err,
	checkRef refChecker,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	if hasPropValue[string](schema, "ref") {
		validateRef(err, checkRef, schema, prop, propName)
		return
	}

	propType := getPropValue[string](schema, "type")

	if propType == TypeObject {
		validateObjectFieldReferences(err, checkRef, schema, prop.(JsonObject), propName)
		return
	}

	if propType == TypeArray {
		validateArrayItemReferences(err, checkRef, schema, prop.(JsonArray), propName)
		return
	}
}

func validateRef(
	err Err,
	checkRef refChecker,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	if getPropValue[string](schema, "type") != TypeString {
		err("%s: Properties that reference another may only be strings", propName)
		return
	}

	ref := getPropValue[string](schema, "ref")
	if !checkRef(ref, prop.(string)) {
		err("%s: Reference doesn't exist at '%s'", propName, ref)
	}
}

func validateObjectFieldReferences(
	err Err,
	checkRef refChecker,
	schema JsonObject,
	obj JsonObject,
	propName string,
) {
	fields := getPropValue[SchemaPropFields](schema, "fields")

	for fieldName, subSchema := range fields {
		fieldPropName := propName + "." + fieldName
		fieldValue := obj[fieldName]
		validateReferences(err, checkRef, subSchema, fieldValue, fieldPropName)
	}
}

func validateArrayItemReferences(
	err Err,
	checkRef refChecker,
	schema JsonObject,
	array JsonArray,
	propName string,
) {
	itemsSchema := getPropValue[JsonObject](schema, "items")

	for i, v := range array {
		itemPropName := fmt.Sprintf("%s[%d]", propName, i)
		validateReferences(err, checkRef, itemsSchema, v, itemPropName)
	}
}
