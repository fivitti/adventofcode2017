package main

import (
	"../utils/intutils"
	"../utils/argparse"
	"fmt"
)

func main() {
	if argparse.ValidateLength(2) != nil {
		fmt.Printf("Day 6: Memory Reallocation. Arguments: path/to/file(string) [separator(string)]")
		return
	}

	separator := argparse.ReadStringOrDefault(2, "\t")
	banks, err := argparse.ReadIntRow(1, separator)
	if err != nil {
		fmt.Println("Invalid file.")
		print(err)
		return
	}

	steps, cycleSize := stepsToCycleAndCycleSize(intutils.Clone(banks))

	fmt.Printf("Steps to cycle: %d. Cycle size: %d.", steps, cycleSize)
}

func stepsToCycleAndCycleSize(banks []int) (int, int) {
	bankStates := make([][]int, 0)
	steps := 0

	firstOccursStateIdx	 := -1
	for firstOccursStateIdx == -1 {
		bankStates = append(bankStates, intutils.Clone(banks))
		nextBankState(banks)
		steps += 1
		firstOccursStateIdx = intutils.IndexOfListList(bankStates, banks)
	}

	return steps, steps - firstOccursStateIdx
}

func nextBankState(banks []int) {
	toRedistributeIdx := intutils.IndexOfMaximum(banks)
	toRedistribute := banks[toRedistributeIdx]

	if toRedistribute == 0 {
		return
	}

	banks[toRedistributeIdx] = 0
	startIdx := toRedistributeIdx + 1
	if startIdx == len(banks) {
		startIdx = 0
	}

	intutils.CycleIterate(banks, startIdx, func(banks []int, idx int) bool {
		banks[idx] += 1
		toRedistribute -= 1
		return toRedistribute != 0
	})
}
