package intutils

import (
	"sort"
)

func SortBytes(arr []byte) []byte {
	ints := BytesToInts(arr)
	sort.Ints(ints)
	return IntsToBytes(ints)
}