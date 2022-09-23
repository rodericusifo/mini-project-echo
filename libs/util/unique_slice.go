package util

func UniqueSlice[T comparable](slices ...[]T) []T {
	uniqueMap := make(map[T]bool)

	for _, slice := range slices {
		for _, val := range slice {
			uniqueMap[val] = true
		}
	}

	result := make([]T, 0, len(uniqueMap))

	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}
