package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../utils/stringutils"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 10: Knot Hash. Parameters: /path/to/file(string) [chainLength=256](int) [rounds=64](int)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	row, err := argparse.ReadStringRow(1)
	if err != nil {
		return err
	}

	chainLength := argparse.ReadDecimalOrDefault(2, 256)
	rounds := argparse.ReadDecimalOrDefault(3, 64)

	fixedSuffix := []int{17, 31, 73, 47, 23}
	chars := stringutils.StringToChars(row)
	values := intutils.BytesToInts(chars)
	values = append(values, fixedSuffix...)

	chain := intutils.Range(0, chainLength, 1)
	skipSize := 0
	currentPosition := 0

	for i := 0; i < rounds; i++ {
		for _, length := range values {
			for currentPosition >= len(chain) {
				currentPosition -= len(chain)
			}
			currentPosition = apply(chain, length, currentPosition)
			currentPosition += skipSize
			skipSize += 1
		}
	}

	hashInt := reduceChain(chain)
	hashBytes := intutils.IntsToBytes(hashInt)
	hash := intutils.BytesToHexes(hashBytes, "")

	fmt.Println("Hash:    ", hash)

	return nil
}

func apply(chain []int, length int, startPosition int) int {
	if length == 0 {
		return startPosition
	}

	endPosition := startPosition + length
	var partToReverse []int
	if endPosition > len(chain) {
		endPosition -= len(chain)
		partToReverse = append(chain[startPosition:len(chain)], chain[0:endPosition]...)
	} else {
		partToReverse = chain[startPosition:endPosition]
	}

	intutils.Reverse(partToReverse)

	reversedIdx := 0

	intutils.CycleIterate(chain, startPosition, func (arr []int, idx int) bool {
		arr[idx] = partToReverse[reversedIdx]
		reversedIdx += 1
		return idx + 1 != endPosition
	})

	return endPosition
}

func reduceChain(chain []int) []int {
	result := make([]int, 0)
	const GROUP_SIZE = 16

	groups := intutils.Group(chain, GROUP_SIZE)

	for _, group := range groups {
		hashValue := intutils.Reduce(group, func (acc, val int) int {
			return acc ^ val
		})
		result = append(result, hashValue)
	}

	return result
}