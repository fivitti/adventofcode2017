package intutils

func BytesToInts(arr []byte) []int {
	result := make([]int, len(arr))
	for idx, val := range arr {
		result[idx] = int(val)
	}
	return result
}

func IntsToBytes(arr []int) []byte {
	result := make([]byte, len(arr))
	for idx, val := range arr {
		result[idx] = byte(val)
	}
	return result
}

func Clone(arr []int) []int {
	result := make([]int, len(arr))
	for idx, val := range arr {
		result[idx] = val
	}
	return result
}
