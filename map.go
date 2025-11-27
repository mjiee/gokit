package gokit

// MapToSlice converts a map to a slice
//
// Parameters:
//   - elements - The map to convert
//   - transform - The function to apply to each element
//
// Returns:
//   - []T - The slice of transformed elements
//
// Example:
//
//	map := map[string]int{"a": 1, "b": 2, "c": 3}
//	slice := MapToSlice(map, func(k string, v int) int {
//	    return v * 2
//	})
//
// slice := []int{2, 4, 6}
func MapToSlice[K comparable, V, T any](elements map[K]V, transform func(K, V) T) []T {
	result := make([]T, 0, len(elements))

	for k, v := range elements {
		result = append(result, transform(k, v))
	}

	return result
}
