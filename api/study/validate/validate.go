package validate

import (
	"fmt"
)

func NewErrorSlice() (*[]string, Err) {
	errors := []string{}

	err := func(msg string, args ...any) {
		msg = fmt.Sprintf(msg, args...)
		errors = append(errors, msg)
	}

	return &errors, err
}

func Validate(schema JsonObject, data JsonValue, dataName string) []string {
	pErrors, err := NewErrorSlice()

	validateStructure(
		err,
		schema,
		data,
		dataName,
	)

	return *pErrors
}
