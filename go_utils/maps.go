package go_utils

func GetMapValue[K1 comparable, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2, defaultVal V) V {
	if inner, ok := m[k1]; ok {
		if v, ok := inner[k2]; ok {
			return v
		}
	}
	return defaultVal
}

func SetMapValue[K1 comparable, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2, val V) {
	if m[k1] == nil {
		m[k1] = make(map[K2]V)
	}
	m[k1][k2] = val
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	targetMap := make(map[K]V)
	for key, val := range m {
		targetMap[key] = val
	}

	return targetMap
}
