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
	return intutils.ByteToInt(byteNumbers)
}

func ConvertToMatrix(input []string, valueSeparator string) ([][]int, error) {
	result := make([][]int, len(input))

	for idxRow, row := range input {
		rawValues := strings.Split(row, valueSeparator)
		values := make([]int, len(rawValues))
		for idxVal, rawVal := range rawValues {
			val, err := strconv.ParseInt(rawVal, 10, 0)
			if err != nil {
				return make([][]int, 0), err
			}
			values[idxVal] = int(val)
		}
		result[idxRow] = values
	}

	return result, nil
}