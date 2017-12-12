package main

import "fmt"
import (
	"../utils/stringutils"
	"../utils/intutils"
	"../utils/argparse"
	"../utils/parsers"
	"strings"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 12: Digital Plumber: /path/to/file(string) |", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}

	rawInput, err := argparse.ReadStringMatrix(1, " ")
	if err != nil {
		return err
	}

	input, err := readInput(rawInput)
	if err != nil {
		return err
	}

	groups := calculateGroups(input)

	_, zeroGroup := findGroup(groups, 0)

	fmt.Println("Group with zero contains:", len(zeroGroup), "elements.")
	fmt.Println("Total group count:", len(groups))

	return nil
}

func readInput(input [][]string) (map[int][]int, error) {
	result := make(map[int][]int)

	for _, row := range input {
		program, err := intutils.ParseInt(row[0])
		if err != nil {
			return nil, err
		}
		neighbours, err := parsers.StringsToNumbers(stringutils.MapStrings(row[2:], func (val string) string {
			return strings.TrimRight(val, ",")
		}))

		if err != nil {
			return nil, err
		}

		result[program] = neighbours
	}

	return result, nil
}

func calculateGroups(input map[int][]int) [][]int {
	groups := make([][]int, 0)

	for program := range input {
		groups = append(groups, []int { program })
	}

	for program, neightbours := range input {
		programGroupIdx, programGroup := findGroup(groups, program)
		groups = intutils.DeleteFromListList(groups, programGroupIdx)

		for _, neightbour := range neightbours {
			if intutils.Contains(programGroup, neightbour) {
				continue
			}
			neightbourGroupIdx, neightbourGroup := findGroup(groups, neightbour)
			groups = intutils.DeleteFromListList(groups, neightbourGroupIdx)

			programGroup = intutils.Concat(programGroup, neightbourGroup)
		}

		groups = append(groups, programGroup)
	}

	return groups
}

func findGroup(groups [][]int, program int) (int, []int) {
	for idx, group := range groups {
		if intutils.Contains(group, program) {
			return idx, group
		}
	}
	return 0, nil
}
