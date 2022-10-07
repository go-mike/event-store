package maps

func MapValue[K comparable, V any, R any](source map[K]V, selector func(key K, value V) R) map[K]R {
	result := make(map[K]R)
	for key, value := range source {
		result[key] = selector(key, value)
	}
	return result
}
