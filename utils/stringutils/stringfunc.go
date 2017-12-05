package stringutils

func FilterListList(arr [][]string, f func ([]string) bool) [][]string {
	result := make([][]string, 0)

	for _, item := range arr {
		if f(item) {
			result = append(result, item)
		}
	}

	return result
}

func MapStrings(arr []string, f func (string) string) []string {
	result := make([]string, len(arr))
	for idx, val := range arr {
		result[idx] = f(val)
	}
	return result
}

func MapListList(arr [][]string, f func ([]string) []string) [][]string {
	result := make([][]string, 0)

	for _, item := range arr {
		result = append(result, f(item))
	}

	return result
}

func MapAndReduceListList(arr [][]string, f func ([]string) string) []string {
	result := make([]string, len(arr))
	for idx, val := range arr {
		result[idx] = f(val)
	}
	return result
}