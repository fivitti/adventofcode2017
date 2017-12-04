package intutils

func MapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for idx, val := range arr {
		result[idx] = val
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
