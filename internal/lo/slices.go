package lo

func Map[I, O any](items []I, fn func(i int, item I) O) []O {
	result := make([]O, 0, len(items))

	for i, item := range items {
		result = append(result, fn(i, item))
	}
	return result
}

func Filter[I any](items []I, fn func(i int, item I) bool) []I {
	var result []I

	for i, item := range items {
		if fn(i, item) {
			result = append(result, item)
		}
	}

	return result
}
