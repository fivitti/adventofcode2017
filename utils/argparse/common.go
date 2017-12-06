package argparse

import (
	"strings"
	"../fileutils"
	"../parsers"
)

func ReadPath(position int) (string, error) {
	value, err := getValue(position)
	if err != nil {
		return "", err
	}
	err = fileutils.FileExists(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func ReadStringMatrix(position int, separator string) ([][]string, error) {
	path, err := ReadPath(1)
	if err != nil {
		return nil, err
	}

	lines, err := fileutils.ReadAllLines(path)
	if err != nil {
		return nil, err
	}

	matrix := make([][]string, len(lines))
	for lineIdx, line := range lines {
		splited := strings.Split(line, separator)
		matrix[lineIdx] = splited
	}

	return matrix, nil
}

func ReadIntMatrix(position int, separator string) ([][]int, error) {
	rawMatrix, err := ReadStringMatrix(position, separator)
	if err != nil {
		return nil, err
	}
	matrix := make([][]int, len(rawMatrix))
	for rowIdx, row := range rawMatrix {
		matrix[rowIdx], err = parsers.StringsToNumbers(row)
		if err != nil {
			return nil, err
		}
	}
	return matrix, nil
}

func ReadIntColumn(position int) ([]int, error) {
	path, err := ReadPath(1)
	if err != nil {
		return nil, err
	}

	lines, err := fileutils.ReadAllLines(path)
	if err != nil {
		return nil, err
	}

	return parsers.StringsToNumbers(lines)
}

func ReadIntRow(position int, separator string) ([]int, error) {
	path, err := ReadPath(1)
	if err != nil {
		return nil, err
	}

	lines, err := fileutils.ReadAllLines(path)
	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return make([]int, 0), nil
	}

	line := lines[0]
	rawValues := strings.Split(line, separator)
	return parsers.StringsToNumbers(rawValues)
}
