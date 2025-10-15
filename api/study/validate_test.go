package study

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func parseJson(jsonStr string) JsonObject {
	data := map[string]any{}
	e := json.Unmarshal([]byte(jsonStr), &data)

	if e != nil {
		panic(e)
	}

	return data
}

func requireErrors(t *testing.T, actErrors []string, expErrors ...string) {
	require.ElementsMatch(t, actErrors, expErrors)
}
