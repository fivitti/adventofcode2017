// Advent of code 2017. Day 1.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../utils/parsers"
)

func main() {
	if argparse.ValidateLength(2) != nil {
		fmt.Println("Robot captcha. Using with arguments: hash(string) [shiftStep(uint)]")
		return
	}
	input, _ := argparse.ReadString(1)
	shiftStep := argparse.ReadDecimalOrDefault(2, 1)

	numbers := parsers.StringToNumbers(input)
	nextNumbers := intutils.Shift(numbers, shiftStep)

	pairs := intutils.Zip([][]int{numbers, nextNumbers})
	theSamePairs := intutils.FilterLists(pairs, func(pair []int) bool {
		if len(pair) == 0 {
			return true
		}
		first := pair[0]
		return intutils.All(pair, func(val int) bool {
			return val == first
		})
	})
	theSameNumbers := intutils.MapAndReduceListList(theSamePairs, func(pair []int) int {
		if len(pair) == 0 {
			return 0
		}
		return pair[0]
	})
	sum := intutils.Reduce(theSameNumbers, func(accumulator int, val int) int {
		return accumulator + val
	})

	fmt.Println(sum)
}


