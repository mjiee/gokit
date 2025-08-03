package slicex

// GroupBy groups the elements of a slice by a key selector function.
//
// Parameters:
//   - elements: the input slice to process ([]T)
//   - keySelector: function that extracts a comparable key from each element (func(T) K)
//
// Returns:
//   - map[K][]T: a map where the keys are the result of the key selector function,
//     and the values are slices of elements that share the same key
//
// Example:
//
//	// Group users by age
//	users := []User{
//	    {ID: 1, Name: "Alice", Age: 30},
//	    {ID: 2, Name: "Bob", Age: 25},
//	    {ID: 3, Name: "Charlie", Age: 30},
//	}
//	groupedUsers := GroupBy(users, func(u User) int { return u.Age })
//
//	// groupedUsers == map[int][]User{
//	//     25: []User{{ID: 2, Name: "Bob", Age: 25}},
//	//     30: []User{{ID: 1, Name: "Alice", Age: 30}, {ID: 3, Name: "Charlie", Age: 30}},
//	// }
func GroupBy[T any, K comparable](elements []T, keySelector func(item T) K) map[K][]T {
	groupedItems := make(map[K][]T)

	for _, item := range elements {
		key := keySelector(item)

		groupedItems[key] = append(groupedItems[key], item)
	}

	return groupedItems
}
