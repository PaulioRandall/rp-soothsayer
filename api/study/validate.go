package study

import (
	"fmt"
)

type JsonType = string
type JsonValue = any
type JsonObject = map[string]JsonValue
type JsonArray = []JsonValue

type Err = func(msg string, args ...any)
type SchemaPropFields = map[string]JsonObject

const (
	TypeObject  JsonType = "object"
	TypeArray            = "array"
	TypeString           = "string"
	TypeNumber           = "number"
	TypeUnknown          = "unknown"
)

func hasPropValue[R any](schema JsonObject, name string) bool {
	_, ok := schema[name].(R)
	return ok
}

func getPropValue[R any](schema JsonObject, name string) R {
	return schema[name].(R)
}

func newErrorSlice() (*[]string, Err) {
	errors := []string{}

	err := func(msg string, args ...any) {
		msg = fmt.Sprintf(msg, args...)
		errors = append(errors, msg)
	}

	return &errors, err
}

func validateSchema(
	err Err,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	propType := getPropValue[JsonType](schema, "type")

	typeMatch := validateSchemaType(
		err,
		propType,
		prop,
		propName,
	)

	if !typeMatch {
		return
	}

	if propType == TypeObject {
		validateSchemaFields(err, schema, prop.(JsonObject), propName)
	}

	if propType == TypeArray {
		validateSchemaItems(err, schema, prop.(JsonArray), propName)
	}
}

func validateSchemaFields(
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
		validateSchema(
			err,
			subSchema,
			prop[name],
			propName+"."+name,
		)
	}
}

func validateSchemaItems(
	err Err,
	schema JsonObject,
	prop JsonArray,
	propName string,
) {
	itemSchema := getPropValue[JsonObject](schema, "items")

	for i, v := range prop {
		validateSchema(
			err,
			itemSchema,
			v,
			fmt.Sprintf("%s[%d]", propName, i),
		)
	}
}

func validateSchemaType(
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

func determineType(prop JsonValue) JsonType {
	switch prop.(type) {
	case float64:
		return TypeNumber
	case string:
		return TypeString
	case JsonObject:
		return TypeObject
	case JsonArray:
		return TypeArray
	default:
		return TypeUnknown
	}
}
