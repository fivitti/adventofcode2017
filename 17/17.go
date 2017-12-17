package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../utils/commonutils"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 17: Spinlock. Parameters: step(int)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	step, err := argparse.ReadDecimalInt(1)
	if err != nil {
		return err
	}

	const VALUE_COUNT = 2018

	currentPosition := 0
	buffer := []int{0}

	for value := 1; value < VALUE_COUNT; value++ {
		currentPosition = commonutils.NextCircularIndex(len(buffer), currentPosition, step) + 1
		buffer = intutils.Insert(buffer, currentPosition, value)
	}


	endPosition := intutils.IndexOf(buffer, 2017)
	outputPosition := commonutils.NextCircularIndex(len(buffer), endPosition, 1)
	outputValue := buffer[outputPosition]

	fmt.Println("Part one:", outputValue)

	outputValueTwo := executePart2(step)
	fmt.Println("Part two:", outputValueTwo)

	return nil
}

func executePart2(step int) int {
	const VALUE_COUNT = 50000000

	zeroPosition := 0
	nextZeroValue := 0
	currentPosition := 0
	bufferLength := 1

	for value := 1; value < VALUE_COUNT; value++ {
		currentPosition = commonutils.NextCircularIndex(bufferLength, currentPosition, step) + 1
		if currentPosition == zeroPosition + 1 {
			nextZeroValue = value
		} else if currentPosition <= zeroPosition {
			zeroPosition += 1
		}
		bufferLength += 1

		if value % 500000 == 0 {
			fmt.Println(value)
		}
	}

	return nextZeroValue
}

