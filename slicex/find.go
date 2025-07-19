package slicex

// Find returns the first element in a slice that satisfies the given predicate function.
// If no element is found, it returns the zero value of type T.
//
// Parameters:
//   - elements: The input slice to search (type []T where T is any type)
//   - predicate: A function that takes an element of type T and returns a bool.
//     The first element that evaluates to true is returned.
//
// Returns:
//   - The first matching element, or zero value if none found.
//
// Type Parameters:
//   - T: The type of elements in the slice (can be any type)
//
// Example:
//
//	// Find first even number
//	numbers := []int{1, 3, 2, 4}
//	firstEven := Find(numbers, func(n int) bool { return n%2 == 0 })
//	// firstEven == 2
func Find[T any](elements []T, predicate func(T) bool) T {
	var result T

	for _, item := range elements {
		if predicate(item) {
			result = item
			break
		}
	}

	return result
}

// FindLast returns the last element in a slice that satisfies the given predicate function.
// If no element is found, it returns the zero value of type T.
//
// Parameters:
//   - elements: The input slice to search (type []T where T is any type)
//   - predicate: A function that takes an element of type T and returns a bool.
//     The last element that evaluates to true is returned.
//
// Returns:
//   - The last matching element, or zero value if none found.
//
// Type Parameters:
//   - T: The type of elements in the slice (can be any type)
//
// Example:
//
//	// Find last even number
//	numbers := []int{1, 2, 3, 4, 5}
//	lastEven := FindLast(numbers, func(n int) bool { return n%2 == 0 })
//	// lastEven == 4
func FindLast[T any](elements []T, predicate func(T) bool) T {
	var result T

	for _, item := range elements {
		if predicate(item) {
			result = item
		}
	}

	return result
}
