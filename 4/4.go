package main

import (
	"../utils/fileutils"
	"../utils/stringutils"
	"../utils/argparse"
	"strings"
	"fmt"
)

func main() {
	err := argparse.ValidateLength(2)
	if err != nil {
		fmt.Println("Day 4: High-Entropy Passphrases. Arguments: path/to/file(string) [separator](string)=\" \"",
			"Invalid input file.")
		return
	}

	separator := argparse.ReadStringOrDefault(2, " ")

	matrix, err := argparse.ReadStringMatrix(1, separator)
	if err != nil {
		fmt.Println("Invalid input file")
		return
	}

	// Part 1
	linesWithDuplicates := stringutils.FilterListList(matrix, existDuplicates)
	linesWithoutDuplicates := len(matrix) - len(linesWithDuplicates)

	fmt.Printf("Lines without duplicates: %d.\n", linesWithoutDuplicates)

	// Part 2
	matrix = stringutils.MapListList(matrix, func (row []string) []string {
		return stringutils.MapStrings(row, stringutils.SortLetters)
	})
	linesWithDuplicates = stringutils.FilterListList(matrix, existDuplicates)
	linesWithoutDuplicates = len(matrix) - len(linesWithDuplicates)

	fmt.Printf("Lines without anagram duplicates: %d.", linesWithoutDuplicates)
}

func existDuplicates(entry []string) bool {
	set := make(map[string]bool)
	for _, item := range entry {
		if _, exists := set[item]; exists {
			return true
		}
		set[item] = true
	}
	return false
}