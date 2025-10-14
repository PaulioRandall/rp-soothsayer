package study

import (
	"fmt"
)

type JsonValue = any
type JsonObject = map[string]JsonValue
type JsonArray = []JsonValue

type Err = func(msg string, args ...any)
type PropFields = map[string]JsonObject

const (
	TypeObject  string = "object"
	TypeArray          = "array"
	TypeString         = "string"
	TypeNumber         = "number"
	TypeUnknown        = "unknown"
)

func getPropType(schema JsonObject) string {
	return schema["type"].(string)
}

func getPropFields(schema JsonObject) PropFields {
	return schema["fields"].(PropFields)
}

func getPropItems(schema JsonObject) JsonObject {
	return schema["items"].(JsonObject)
}

// TODO make generic
func getObject(obj JsonObject, name string) JsonObject {
	return obj[name].(JsonObject)
}
func getString(obj JsonObject, name string) string {
	return obj[name].(string)
}

var mediaSectionSchema = JsonObject{
	"type": "object",
	"fields": PropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"start": JsonObject{
			"type": "number",
		},
		"end": JsonObject{
			"type": "number",
		},
	},
}

func ValidateMediaSection(mediaSection JsonObject) []string {
	errors := []string{}
	err := func(msg string, args ...any) {
		msg = fmt.Sprintf(msg, args...)
		errors = append(errors, msg)
	}

	validateSchema(
		err,
		mediaSectionSchema,
		mediaSection,
		"media_section",
	)

	return errors
}

func validateSchema(
	err Err,
	schema JsonObject,
	prop JsonValue,
	propName string,
) {
	propType := getPropType(schema)

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
	for name, subSchema := range getPropFields(schema) {
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
	itemSchema := getPropItems(schema)

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
	expType string,
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

func determineType(prop JsonValue) string {
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
