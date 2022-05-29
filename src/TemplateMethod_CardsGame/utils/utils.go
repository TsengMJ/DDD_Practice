package utils

func GetMapKeys(table map[string]int) []string {
	keys := make([]string, 0)

	for key, _ := range table {
		keys = append(keys, key)
	}

	return keys
}
