// Advent of code 2017. Day 2.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinUint uint = 0
	MaxUint uint = ^MinUint
	MaxInt  int  = int(MaxUint >> 1)
	MinInt  int  = ^MaxInt
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Checksum. Use: path\\to\\data(string) [valueSep=\\t(string)]")
		return
	}

	valueSeparator := "\t"
	if len(os.Args) > 2 {
		valueSeparator = os.Args[3]
	}

	path := os.Args[1]
	input, err := readAllLines(path)
	if err != nil {
		fmt.Println("Invalid input file.")
		return
	}

	matrix, err := convertToMatrix(input, valueSeparator)
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	// Part 1
	selected1 := mapAndReduceListList(matrix, func(row []int) int {
		min, max := minAndMax(row)
		return max - min
	})

	// Part 2
	selected2 := mapAndReduceListList(matrix, func(row []int) int {
		return maxEventlyDivide(row)
	})

	checksum1 := calculateChecksum(selected1)
	checksum2 := calculateChecksum(selected2)

	fmt.Printf("Checksum part 1: %d", checksum1)
	fmt.Printf("Checksum part 2: %d", checksum2)
}

func calculateChecksum(arr []int) int {
	return reduce(arr, func(acc, val int) int {
		return acc + val
	})
}

func readAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convertToMatrix(input []string, valueSeparator string) ([][]int, error) {
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

func minAndMax(arr []int) (int, int) {
	maximum := MinInt
	minimum := MaxInt

	for _, val := range arr {
		if val > maximum {
			maximum = val
		}
		if val < minimum {
			minimum = val
		}
	}

	return minimum, maximum
}

func maxEventlyDivide(arr []int) int {
	result := 1
	for _, dividend := range arr {
		if dividend <= result {
			continue
		}
		for _, divider := range arr {
			if dividend%divider == 0 {
				quotient := dividend / divider
				if result < quotient {
					result = quotient
				}
			}
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
