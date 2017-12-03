package parsers

import (
	"strings"
	"strconv"
)

func StringToNumbers(str string) []int {
	result := make([]int, len(str))

	for idx := 0; idx < len(str); idx++ {
		result[idx] = int(str[idx] - '0')
	}

	return result
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