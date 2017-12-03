// Advent of code 2017. Day 1.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"os"
	"strconv"
	"../utils/intutils"
	"../utils/parsers"
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


