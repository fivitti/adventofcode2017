package intutils

import (
	"strconv"
	"strings"
	"fmt"
)

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

func ParseInt(str string) (int, error) {
	temp, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(temp), err
}

func BytesToHexes(bytes []byte, separator string) string {
	result := make([]string, len(bytes))

	for idx, val := range bytes {
		result[idx] = fmt.Sprintf("%02x", val)
	}

	return strings.Join(result, separator)
}