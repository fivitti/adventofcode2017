// Advent of code 2017. Day 1.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Robot captcha. Using with arguments: hash(string) [shiftStep(uint)]")
		return
	}
	input := args[1]
	shiftStep := 1
	if len(args) > 2 {
		if step, err := strconv.ParseInt(args[2], 10, 0); err != nil {
			shiftStep = 1
		} else {
			shiftStep = int(step)
		}
	}

	numbers := stringToNumbers(input)
	nextNumbers := shift(numbers, shiftStep)

	pairs := zip([][]int{numbers, nextNumbers})
	theSamePairs := filterLists(pairs, func(pair []int) bool {
		if len(pair) == 0 {
			return true
		}
		first := pair[0]
		return all(pair, func(val int) bool {
			return val == first
		})
	})
	theSameNumbers := mapAndReduceListList(theSamePairs, func(pair []int) int {
		if len(pair) == 0 {
			return 0
		}
		return pair[0]
	})
	sum := reduce(theSameNumbers, func(accumulator int, val int) int {
		return accumulator + val
	})

	fmt.Println(sum)
}

func stringToNumbers(str string) []int {
	result := make([]int, len(str))

	for idx := 0; idx < len(str); idx++ {
		result[idx] = int(str[idx] - '0')
	}

	return result
}

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for idx, val := range arr {
		result[idx] = val
	}
	return result
}

func shift(arr []int, step int) []int {
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

func zip(arrs [][]int) [][]int {
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

func filterLists(arrs [][]int, filter func([]int) bool) [][]int {
	result := make([][]int, 0)
	for _, arr := range arrs {
		if filter(arr) {
			result = append(result, arr)
		}
	}
	return result
}

func mapAndReduceListList(arrs [][]int, f func([]int) int) []int {
	result := make([]int, len(arrs))
	for idx := 0; idx < len(arrs); idx++ {
		result[idx] = f(arrs[idx])
	}
	return result
}

func reduce(arr []int, reducer func(int, int) int) int {
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

func all(arr []int, f func(int) bool) bool {
	for _, val := range arr {
		if !f(val) {
			return false
		}
	}
	return true
}
