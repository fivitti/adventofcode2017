package stringutils

import (
	"../intutils"
)

func SortLetters(str string) string {
	chars := StringToChars(str)
	sortedChars := intutils.SortBytes(chars)
	return string(sortedChars)
}
