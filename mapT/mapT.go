package mapT

func MapContains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func MapKeys[K comparable, V any](m map[K]V) []K {//return all keys in the map
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func MapGet[K comparable, V any](m map[K]V, key K) (V, bool) {	
	v, ok := m[key]
	return v, ok
}

func MapPut[K comparable, V any](m map[K]V, key K, value V) {
	m[key] = value
}

func MapRemove[K comparable, V any](m map[K]V, key K) {
	delete(m, key)
}

func MapClear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

func MapSize[K comparable, V any](m map[K]V) int {
	return len(m)
}

func MapIsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

func MapForEach[K comparable, V any](m map[K]V, consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}

func MapFilter[K comparable, V any](m map[K]V, filter func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if filter(k, v) {
			result[k] = v
		}
	}
	return result
}


func MapMap[K comparable, V any, U any](m map[K]V, mapper func(K, V) U) map[K]U {
	result := make(map[K]U)
	for k, v := range m {
		result[k] = mapper(k, v)
	}
	return result
}

func MapMerge[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

func MapMergeAll[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func MapClone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func MapEquals[K comparable, V comparable](m1 map[K]V, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

