package gokit

import "math/rand/v2"

// SliceMap applies a transformation function to each element of a slice and returns a new slice
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
//	lengths := SliceMap(names, func(s string) int { return len(s) })
//	// lengths == []int{5, 3, 7}
func SliceMap[T any, V any](elements []T, transform func(T) V) []V {
	result := make([]V, 0, len(elements))

	for _, item := range elements {
		result = append(result, transform(item))
	}

	return result
}

// SliceMapErr applies a transformation function to each element of a slice and returns a new slice
//
// Parameters:
//   - elements: The input slice to process ([]T)
//   - transform: function that transforms each element of the slice (func(T) (V, error))
//
// Returns:
//   - []V: a slice of transformed elements
//   - error: an error if any occurred during processing
func SliceMapErr[T any, V any](elements []T, transform func(T) (V, error)) ([]V, error) {
	result := make([]V, 0, len(elements))

	for _, item := range elements {
		resultItem, err := transform(item)
		if err != nil {
			return nil, err
		}

		result = append(result, resultItem)
	}

	return result, nil
}

// SliceFilterMap applies a transformation function to each element of a slice and returns a new slice
// that contains the transformed elements.
//
// Parameters:
//   - elements: The input slice to process ([]T)
//   - transform: function that transforms each element of the slice (func(T) (V, error))
//
// Returns:
//   - []V: a slice of transformed elements
//   - error: an error if any occurred during processing
//
// Example:
//
//	// Filter out even numbers
//	numbers := []int{1, 2, 3}
//	numbersByEvenOdd := SliceFilterMap(numbers, func(n int) (bool, string) {
//	    return n%2 == 0, fmt.Sprintf("Number %d is even", n)
//	})
//	// numbersByEvenOdd == []string{"Number 2 is even"}
func SliceFilterMap[T any, V any](elements []T, transform func(T) (bool, V)) []V {
	result := make([]V, 0, len(elements))

	for _, item := range elements {
		isValid, resultItem := transform(item)
		if isValid {
			result = append(result, resultItem)
		}
	}

	return result
}

// SliceGroupBy groups the elements of a slice by a key selector function.
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
//	groupedUsers := SliceGroupBy(users, func(u User) int { return u.Age })
//
//	// groupedUsers == map[int][]User{
//	//     25: []User{{ID: 2, Name: "Bob", Age: 25}},
//	//     30: []User{{ID: 1, Name: "Alice", Age: 30}, {ID: 3, Name: "Charlie", Age: 30}},
//	// }
func SliceGroupBy[T any, K comparable](elements []T, keySelector func(item T) K) map[K][]T {
	groupedItems := make(map[K][]T)

	for _, item := range elements {
		key := keySelector(item)

		groupedItems[key] = append(groupedItems[key], item)
	}

	return groupedItems
}

// SliceFind returns the first element in a slice that satisfies the given predicate function.
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
//	firstEven := SliceFind(numbers, func(n int) bool { return n%2 == 0 })
//	// firstEven == 2
func SliceFind[T any](elements []T, predicate func(T) bool) T {
	var result T

	for _, item := range elements {
		if predicate(item) {
			result = item
			break
		}
	}

	return result
}

// SliceFindLast returns the last element in a slice that satisfies the given predicate function.
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
//	lastEven := SliceFindLast(numbers, func(n int) bool { return n%2 == 0 })
//	// lastEven == 4
func SliceFindLast[T any](elements []T, predicate func(T) bool) T {
	var result T

	for _, item := range elements {
		if predicate(item) {
			result = item
		}
	}

	return result
}

// SliceFilter iterates over a slice and returns a new slice containing all elements
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
//	evens := SliceFilter(numbers, func(n int) bool { return n%2 == 0 })
//	// evens == []int{2, 4}
func SliceFilter[T any](elements []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range elements {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// SliceDistinct returns a new slice containing only the unique elements from the input slice.
//
// The function uses a key selector function to determine uniqueness. Elements are considered
// distinct if they produce different keys when passed to the selector function.
//
// Parameters:
//   - elements - the input slice to process ([]T)
//   - keySelector - function that extracts a comparable key from each element (func(T) K)
//
// Returns:
//   - []T - a new slice containing only the distinct elements from the input
//
// Example:
//
//	// Get distinct users by ID
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 1, Name: "Alice"},
//	}
//	distinctUsers := SliceDistinct(users, func(u User) int { return u.ID })
//
//	// distinctUsers == []User{
//	//     {ID: 1, Name: "Alice"},
//	//     {ID: 2, Name: "Bob"},
//	// }
func SliceDistinct[T any, K comparable](elements []T, keySelector func(item T) K) []T {
	var (
		seenKeys      = make(map[K]bool)
		distinctItems = make([]T, 0, len(elements))
	)

	for _, item := range elements {
		key := keySelector(item)
		if !seenKeys[key] {
			seenKeys[key] = true

			distinctItems = append(distinctItems, item)
		}
	}

	return distinctItems
}

// SliceShuffle randomly shuffles the elements of the given slice and returns it.
//
// The function modifies the order of elements in place using the Fisher–Yates shuffle algorithm
// provided by Go's standard library (`rand.Shuffle`). This ensures that each permutation of the slice
// has an equal probability of occurrence.
//
// Parameters:
//   - elements - the input slice to shuffle ([]T)
//
// Returns:
//   - []T - the same slice with its elements randomly shuffled
//
// Example:
//
//	// Shuffle a list of numbers
//	numbers := []int{1, 2, 3, 4, 5}
//	shuffled := SliceShuffle(numbers)
//
//	// shuffled might be [3, 5, 1, 4, 2]
func SliceShuffle[T any](elements []T) []T {
	if len(elements) == 0 {
		return nil
	}

	copied := make([]T, len(elements))
	copy(copied, elements)

	rand.Shuffle(len(copied), func(i, j int) {
		copied[i], copied[j] = copied[j], copied[i]
	})

	return copied
}

// SliceToMap converts a slice of elements into a map using a key selector function.
//
// Parameters:
//   - elements - the input slice to convert ([]T)
//   - keySelector - function that extracts a key from each element (func(T) K)
//
// Returns:
//   - map[K]T - a map where keys are produced by the keySelector function and values are the original elements
//
// Example:
//
//	// Convert a list of users to a map by their IDs
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	}
//	usersMap := SliceToMap(users, func(u User) int { return u.ID })
//
//	// usersMap == map[int]User{
//	//     1: {ID: 1, Name: "Alice"},
//	//     2: {ID: 2, Name: "Bob"},
//	// }
func SliceToMap[T any, K comparable](elements []T, keySelector func(item T) K) map[K]T {
	result := make(map[K]T, len(elements))

	for _, item := range elements {
		key := keySelector(item)
		result[key] = item
	}

	return result
}
