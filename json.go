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

// Marshal encodes a Go value into a string according to the provided marshal and encode options.
//
// Parameters:
//   - data: T
//
// Returns:
//   - string: string
//   - error: error
func Marshal(data any) (string, error) {
	bytes, err := json.Marshal(data)

	return string(bytes), err
}

// MarshalSafe encodes a Go value into a string according to the provided marshal and encode options.
//
// Parameters:
//   - data: T
//
// Returns:
//   - string: string
func MarshalSafe(data any) string {
	if data == nil {
		return ""
	}

	result, err := Marshal(data)

	if err != nil {
		return ""
	}

	return result
}
