// Advent of code 2017. Day 2.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"../utils/fileutils"
	"../utils/parsers"
	"../utils/intutils"
	"../utils/argparse"
)

func main() {
	if argparse.ValidateLength(2) != nil {
		fmt.Println("Checksum. Use: path\\to\\data(string) [valueSep=\\t(string)]")
		return
	}

	valueSeparator := argparse.ReadStringOrDefault(2, "\t")
	path, err := argparse.ReadPath(1)
	if err != nil {
		fmt.Println("Invalid file")
		return
	}
	input, err := fileutils.ReadAllLines(path)
	if err != nil {
		fmt.Println("Invalid input file.")
		return
	}

	matrix, err := parsers.ConvertToMatrix(input, valueSeparator)
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	// Part 1
	selected1 := intutils.MapAndReduceListList(matrix, func(row []int) int {
		min, max := minAndMax(row)
		return max - min
	})

	// Part 2
	selected2 := intutils.MapAndReduceListList(matrix, func(row []int) int {
		return maxEventlyDivide(row)
	})

	checksum1 := calculateChecksum(selected1)
	checksum2 := calculateChecksum(selected2)

	fmt.Printf("Checksum part 1: %d", checksum1)
	fmt.Printf("Checksum part 2: %d", checksum2)
}

func calculateChecksum(arr []int) int {
	return intutils.Reduce(arr, func(acc, val int) int {
		return acc + val
	})
}

func minAndMax(arr []int) (int, int) {
	maximum := intutils.MinInt
	minimum := intutils.MaxInt

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
