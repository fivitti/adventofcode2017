package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 15: Dueling Generators. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadStringMatrix(1, " ")
	if err != nil {
		return err
	}

	generatorAInitialValue, err := intutils.ParseInt(matrix[0][4])
	if err != nil {
		return err
	}
	generatorBInitialValue, err := intutils.ParseInt(matrix[1][4])
	if err != nil {
		return err
	}

	const FACTOR_A = 16807
	const FACTOR_B = 48271
	const CYCLES = 5000000
	const MASK = 65535

	aValue := generatorAInitialValue
	bValue := generatorBInitialValue

	judgeCount := 0

	for i := 0; i < CYCLES; i++ {
		aValue = generateNextValue(aValue, FACTOR_A, func (num int) bool {
			return num % 4 == 0
		})
		bValue = generateNextValue(bValue, FACTOR_B, func (num int) bool {
			return num % 8 == 0
		})

		if compareBits(aValue, bValue, MASK) {
			judgeCount += 1
		}
	}

	fmt.Println("Judge count:", judgeCount)

	return nil
}

func generateNextValue(previousValue int, factor int, criteria func(num int) bool) int {
	const REDUCTOR_FACTOR = 2147483647
	valueCandidate := (previousValue * factor) % REDUCTOR_FACTOR
	if criteria(valueCandidate) {
		return valueCandidate
	} else {
		return generateNextValue(valueCandidate, factor, criteria)
	}
}

func compareBits(first int, second int, mask int) bool {
	firstComparePart := first & mask
	secondComparePart := second & mask

	//fmt.Println(intutils.BitsToString(intutils.IntToBits(firstComparePart), 16))
	//fmt.Println(intutils.BitsToString(intutils.IntToBits(secondComparePart), 16))
	//fmt.Println()

	return firstComparePart == secondComparePart
}