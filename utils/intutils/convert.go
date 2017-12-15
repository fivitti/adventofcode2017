package intutils

import (
	"strconv"
	"strings"
	"fmt"
)

const BITS_IN_BYTE = 8
const MOST_SIGNIFICANT_BIT_IN_BYTE_VALUE = 128

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

func BytesToBits(bytes []byte) []bool {
	result := make([]bool, 0)

	for _, byte_ := range bytes {
		bits := ByteToBits(byte_)
		result = append(result, bits...)
	}

	return result
}

func ByteToBits(byte_ byte) []bool {
	result := make([]bool, BITS_IN_BYTE)

	for bit := uint(0); bit < BITS_IN_BYTE; bit++ {
		result[bit] = byte_ & (MOST_SIGNIFICANT_BIT_IN_BYTE_VALUE >> bit) != 0
	}

	return result
}

func IntToBits(num int) []bool {
	result := make([]bool, 0)

	for num > 0 {
		if num & 1 == 1 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
		num = num >> 1
	}

	return ReverseBools(result)
}

func BitsToString(bits []bool, group int) string {
	const SPACE = ' '
	chars := make([]byte, 0)
	for idx, bit := range bits {
		if idx != 0 && group > 1 && idx % group == 0 {
			chars = append(chars, SPACE)
		}
		if bit {
			chars = append(chars, '1')
		} else {
			chars = append(chars, '0')
		}
	}

	return string(chars)
}