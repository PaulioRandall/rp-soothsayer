package validate

import (
	"fmt"
)

func Validate(schema JsonObject, data JsonObject, dataName string) []string {
	pErrors, err := newErrorSlice()

	validateStructure(
		err,
		schema,
		data,
		dataName,
	)

	return *pErrors
}

func newErrorSlice() (*[]string, Err) {
	errors := []string{}

	err := func(msg string, args ...any) {
		msg = fmt.Sprintf(msg, args...)
		errors = append(errors, msg)
	}

	return &errors, err
}
