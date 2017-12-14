package knothash

import (
	"../../utils/intutils"
	"../../utils/stringutils"
)

var FIXED_SUFFIX = []int{17, 31, 73, 47, 23}
const (
	CHAIN_LENGTH = 256
	ROUNDS = 64
	GROUP_SIZE = 16
)

func Hash(input string) []byte {
	chars := stringutils.StringToChars(input)
	values := intutils.BytesToInts(chars)
	values = append(values, FIXED_SUFFIX...)

	chain := intutils.Range(0, CHAIN_LENGTH, 1)
	skipSize := 0
	currentPosition := 0

	for i := 0; i < ROUNDS; i++ {
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

	return hashBytes
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

	groups := intutils.Group(chain, GROUP_SIZE)

	for _, group := range groups {
		hashValue := intutils.Reduce(group, func (acc, val int) int {
			return acc ^ val
		})
		result = append(result, hashValue)
	}

	return result
}
