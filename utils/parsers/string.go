package parsers

import (
	"strings"
	"strconv"
	"../stringutils"
	"../intutils"
)

func StringToNumbers(str string) []int {
	chars := stringutils.StringToChars(str)
	byteNumbers := intutils.MapByte(chars, func (b byte) byte {
		return b - '0'
	})
	return intutils.BytesToInts(byteNumbers)
}

func StringsToNumbers(arr []string) ([]int, error) {
	values := make([]int, len(arr))
	for idxVal, rawVal := range arr {
		val, err := strconv.ParseInt(rawVal, 10, 0)
		if err != nil {
			return nil, err
		}
		values[idxVal] = int(val)
	}
	return values, nil
}

func ConvertToMatrix(input []string, valueSeparator string) ([][]int, error) {
	result := make([][]int, len(input))

	for idxRow, row := range input {
		rawValues := strings.Split(row, valueSeparator)
		values, err := StringsToNumbers(rawValues)
		if err != nil {
			return nil, err
		}
		result[idxRow] = values
	}

	return result, nil
}

