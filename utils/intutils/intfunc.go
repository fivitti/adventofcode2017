package intutils

func MapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for idx, val := range arr {
		result[idx] = f(val)
	}
	return result
}

func MapByte(arr []byte, f func(byte) byte) []byte {
	result := make([]byte, len(arr))
	for idx, val := range arr {
		result[idx] = val
	}
	return result
}

func Shift(arr []int, step int) []int {
	arrLen := len(arr)
	result := make([]int, arrLen)
	var newIdx int
	for idx, val := range arr {
		newIdx = idx + step
		if newIdx >= arrLen {
			newIdx -= arrLen
		}
		result[newIdx] = val
	}
	return result
}

func Zip(arrs [][]int) [][]int {
	arrsCount := len(arrs)
	if arrsCount == 0 {
		return make([][]int, 0)
	}

	arrLen := len(arrs[0])
	result := make([][]int, arrLen)

	for idx := 0; idx < arrLen; idx++ {
		partResult := make([]int, arrsCount)
		for arrIdx := 0; arrIdx < arrsCount; arrIdx++ {
			partResult[arrIdx] = arrs[arrIdx][idx]
		}
		result[idx] = partResult
	}

	return result
}

func FilterLists(arrs [][]int, filter func([]int) bool) [][]int {
	result := make([][]int, 0)
	for _, arr := range arrs {
		if filter(arr) {
			result = append(result, arr)
		}
	}
	return result
}

func MapAndReduceListList(arrs [][]int, f func([]int) int) []int {
	result := make([]int, len(arrs))
	for idx := 0; idx < len(arrs); idx++ {
		result[idx] = f(arrs[idx])
	}
	return result
}

func Reduce(arr []int, reducer func(int, int) int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	accumulator := arr[0]
	for idx := 1; idx < len(arr); idx++ {
		accumulator = reducer(accumulator, arr[idx])
	}
	return accumulator
}

func All(arr []int, f func(int) bool) bool {
	for _, val := range arr {
		if !f(val) {
			return false
		}
	}
	return true
}

func Same(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	return All(arr, func(val int) bool {
		return val == arr[0]
	})
}

func Equals(arrs [][]int) bool {
	for _, group := range Zip(arrs) {
		if !Same(group) {
			return false
		}
	}
	return true
}

func IndexOf(arr []int, item int) int {
	for idx, val := range arr {
		if val == item {
			return idx
		}
	}
	return -1
}

func Contains(arr []int, item int) bool {
	return IndexOf(arr, item) != -1
}

func IndexOfMaximum(arr []int) int {
	maximumIdx := 0
	maximum := MinInt

	for idx, val := range arr {
		if val > maximum {
			maximum = val
			maximumIdx = idx
		}
	}
	return maximumIdx
}

func IndexOfListList(matrix [][]int, arr[]int) int {
	for idx, matrixEntry := range matrix {
		if Equals([][]int{matrixEntry, arr}) {
			return idx
		}
	}
	return -1
}

func CycleIterate(arr []int, startIdx int, f func (arr []int, idx int) bool) {
	arrLen := len(arr)
	for f(arr, startIdx) {
		startIdx += 1
		if startIdx == arrLen {
			startIdx = 0
		}
	}
}

func Reverse(arr []int) {
	for left, right := 0, len(arr) - 1; left < right; left, right = left + 1, right - 1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
}

func Range(start int, end int, step int) []int {
	result := make([]int, 0)
	for val := start; val < end; val += step {
		result = append(result, val)
	}
	return result
}

func Group(arr []int, groupSize int) [][]int {
	result := make([][]int, 0)
	var currentGroup []int

	for idx, val := range arr {
		// Begin next group
		if idx % groupSize == 0 {
			if currentGroup != nil {
				result = append(result, currentGroup)
			}
			currentGroup = make([]int, 0)
		}

		currentGroup = append(currentGroup, val)
	}

	result = append(result, currentGroup)
	return result
}

func Concat(first []int, second []int) []int {
	return append(first, second...)
}

func Delete(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func DeleteFromListList(arr [][]int, idx int) [][]int {
	return append(arr[:idx], arr[idx+1:]...)
}