package extenders

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

// WithFieldsLogger -
type WithFieldsLogger interface {
	TraceWithFields(fields Fields)
	PanicWithFields(fields Fields)
	FatalWithFields(fields Fields)
	ErrorWithFields(fields Fields)
	WarningWithFields(fields Fields)
	InfoWithFields(fields Fields)
	SuccessWithFields(fields Fields)
	DebugWithFields(fields Fields)
}
