package dataframe

import (
	"fmt"
)

const (
	FALSE = 0
	TRUE  = 1
)

// B converts a boolean to an int.
func B(b bool) int {
	if b {
		return 1
	}
	return 0
}

// BoolValueFormatter is used by SetValueToStringFormatter
// to display an int as a bool. If the encountered value
// is not a 0 or 1, it will panic.
func BoolValueFormatter(v interface{}) string {
	if v == nil {
		return "NaN"
	}

	str := fmt.Sprintf("%v", v)
	switch str {
	case "0":
		return "false"
	case "1":
		return "true"
	default:
		_ = v.(bool) // Intentionally panic
		return ""
	}
}

// DefaultValueFormatter will return a string representation
// of the data in a particular row.
func DefaultValueFormatter(v interface{}) string {
	if v == nil {
		return "NaN"
	}
	return fmt.Sprintf("%v", v)
}
