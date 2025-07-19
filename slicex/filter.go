package slicex

// Filter iterates over a slice and returns a new slice containing all elements
// that satisfy the given predicate function.
//
// Parameters:
//   - elements: The input slice to filter (type []T where T is any type)
//   - predicate: A function that takes an element of type T and returns a bool.
//     Elements that evaluate to true are included in the result.
//
// Returns:
//   - A new slice containing only elements that satisfy the predicate.
//
// Type Parameters:
//   - T: The type of elements in the slice (can be any type)
//
// Example:
//
//	// Filter even numbers
//	numbers := []int{1, 2, 3, 4, 5}
//	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
//	// evens == []int{2, 4}
func Filter[T any](elements []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range elements {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
