package gokit

import "encoding/json"

// UnmarshalSafe is a safe version of json.Unmarshal
//
// Parameters:
//   - data: string
//
// Returns:
//   - T: T
func UnmarshalSafe[T any](data string) T {
	var result T

	if data == "" {
		return result
	}

	_ = json.Unmarshal([]byte(data), &result)

	return result
}

// Unmarshal decodes a string input into a Go value according to the provided unmarshal and decode options.
//
// Parameters:
//   - data: string
//
// Returns:
//   - T: T
//   - error: error
func Unmarshal[T any](data string) (T, error) {
	var result T

	if data == "" {
		return result, nil
	}

	err := json.Unmarshal([]byte(data), &result)

	return result, err
}
