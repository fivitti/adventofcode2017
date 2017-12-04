package main

import (
	"../utils/fileutils"
	"../utils/stringutils"
	"../utils/argparse"
	"strings"
	"fmt"
)

func main() {
	path, err := argparse.ReadPath(1)
	if err != nil {
		fmt.Println("Day 4: High-Entropy Passphrases. Arguments: path/to/file(string) [separator](string)=\" \"",
			"Invalid input file.")
		return
	}

	separator := argparse.ReadStringOrDefault(2, " ")

	matrix, err := readFile(path, separator)
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

func readFile(path, sep string) ([][]string, error) {
	lines, err := fileutils.ReadAllLines(path)
	if err != nil {
		return nil, err
	}
	return splitLines(lines, sep), nil
}

func splitLines(lines []string, sep string) [][]string {
	matrix := make([][]string, len(lines))
	for lineIdx, line := range lines {
		splited := strings.Split(line, sep)
		matrix[lineIdx] = splited
	}
	return matrix
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