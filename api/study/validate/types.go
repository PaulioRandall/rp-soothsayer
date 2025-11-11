package validate

// TODO: Much of this can be packag private.

type JsonValue = any
type JsonObject = map[string]JsonValue
type JsonArray = []JsonValue
type JsonType = string

type Err = func(msg string, args ...any)
type SchemaPropFields = map[string]JsonObject

const (
	TypeObject  JsonType = "object"
	TypeArray            = "array"
	TypeString           = "string"
	TypeNumber           = "number"
	TypeUnknown          = "unknown"
)

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
