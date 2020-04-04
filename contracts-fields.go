package logging

import (
	"encoding/json"
)

// Fields -
type Fields map[string]interface{}

func (thisRef Fields) String() string {
	bytes, err := json.Marshal(thisRef)
	if err != nil {
		return ""
	}

	return string(bytes)
}
