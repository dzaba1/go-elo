package main

func GetKeys[K comparable, V any](dict map[K]V) []K {
	keys := make([]K, len(dict))

	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}

	return keys
}

func GetValues[K comparable, V any](dict map[K]V) []V {
	values := make([]V, len(dict))

	i := 0
	for _, v := range dict {
		values[i] = v
		i++
	}

	return values
}
