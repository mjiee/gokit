package slicex

// Distinct returns a new slice containing only the unique elements from the input slice.
//
// The function uses a key selector function to determine uniqueness. Elements are considered
// distinct if they produce different keys when passed to the selector function.
//
// Parameters:
//
//	sourceSlice - the input slice to process ([]T)
//	keySelector - function that extracts a comparable key from each element (func(T) K)
//
// Returns:
//
//	[]T - a new slice containing only the distinct elements from the input
//
// Example:
//
//	// Get distinct users by ID
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 1, Name: "Alice"},
//	}
//	distinctUsers := Distinct(users, func(u User) int { return u.ID })
func Distinct[T any, K comparable](sourceSlice []T, keySelector func(item T) K) []T {
	var (
		seenKeys      = make(map[K]bool)
		distinctItems = make([]T, 0, len(sourceSlice))
	)

	for _, item := range sourceSlice {
		key := keySelector(item)
		if !seenKeys[key] {
			seenKeys[key] = true

			distinctItems = append(distinctItems, item)
		}
	}

	return distinctItems
}
