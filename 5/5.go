package main

import (
	"../utils/argparse"
	"../utils/parsers"
	"../utils/fileutils"
	"../utils/intutils"
	"fmt"
)

func main() {
	if argparse.ValidateLength(2) != nil {
		fmt.Printf("Day 5: A Maze of Twisty Trampolines, All Alike. Arguments: path/to/file(string)")
		return
	}

	path, err := argparse.ReadPath(1)
	if err != nil {
		fmt.Printf("Invalid file: %s.", err.Error())
		return
	}

	rawColumn, err := fileutils.ReadAllLines(path)
	if err != nil {
		fmt.Println("Invalid file.")
		print(err)
		return
	}

	input, err := parsers.StringsToNumbers(rawColumn)
	if err != nil {
		fmt.Println("Invalid file.")
		print(err)
		return
	}

	// Part 1
	steps := calculateSteps(intutils.Clone(input), func (currentValue int) int {
		return currentValue + 1
	})

	fmt.Printf("Part 1. Steps: %d.\n", steps)

	// Part 2
	steps = calculateSteps(intutils.Clone(input), func (currentValue int) int {
		if currentValue >= 3 {
			return currentValue - 1
		}
		return currentValue + 1
	})

	fmt.Printf("Part 2. Steps: %d.", steps)
}

func calculateSteps(input []int, changeCell func (int) int) int {
	steps := 0
	currentIdx := 0
	for currentIdx < len(input) {
		toJump := input[currentIdx]
		input[currentIdx] = changeCell(input[currentIdx])
		currentIdx += toJump
		steps += 1
	}
	return steps
}


