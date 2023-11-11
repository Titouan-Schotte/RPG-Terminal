package utils

import (
	"sort"
)

func ContainsKey(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

func GetKeys(mapIn map[string]int) []string {
	keys := make([]string, len(mapIn))

	i := 0
	for k := range mapIn {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func GetFirstKey(mapIn map[string]string) string {
	for key := range mapIn {
		return key
	}
	return ""
}
