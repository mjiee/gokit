package slicex

// Map applies a transformation function to each element of a slice and returns a new slice
// containing the transformed values. This is a functional programming style operation
// commonly known as "map" or "transform".
//
// Parameters:
//   - elements: The input slice to transform (type []T where T is any type)
//   - transform: A function that takes an element of type T and returns a value of type V.
//     This function is applied to each element in the slice.
//
// Returns:
//   - A new slice of type []V containing the transformed values.
//
// Type Parameters:
//   - T: The type of elements in the input slice
//   - V: The type of elements in the output slice
//
// Example:
//
//	// Convert strings to their lengths
//	names := []string{"Alice", "Bob", "Charlie"}
//	lengths := Map(names, func(s string) int { return len(s) })
//	// lengths == []int{5, 3, 7}
//
//	// Convert numbers to strings
//	numbers := []int{1, 2, 3}
//	strings := Map(numbers, func(n int) string { return fmt.Sprintf("%d", n) })
//	// strings == []string{"1", "2", "3"}
func Map[T any, V any](elements []T, transform func(T) V) []V {
	result := make([]V, 0, len(elements))

	for _, item := range elements {
		result = append(result, transform(item))
	}

	return result
}
