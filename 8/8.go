package main

import (
	"strconv"
	"../utils/argparse"
	"../utils/intutils"
	"fmt"
)

func main() {
	if argparse.ValidateLength(2) != nil {
		fmt.Println("Day 8: I Heard You Like Registers. Parameter: path/to/file(string)")
		return
	}

	if err := execute(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func execute() error {
	input, err := argparse.ReadStringMatrix(1, " ")
	if err != nil {
		return err
	}

	variableMap := make(map[string]int)

	biggestAnytime := 0

	for _, row := range input {
		err := executeRow(variableMap, row)
		if err != nil {
			return nil
		}
		biggestNow := getBiggestValue(variableMap)
		if biggestNow > biggestAnytime {
			biggestAnytime = biggestNow
		}
	}

	biggestEnd := getBiggestValue(variableMap)

	fmt.Println("Biggest value: ", biggestEnd)
	fmt.Println("Biggest anytime value: ", biggestAnytime)
	return nil
}

func getBiggestValue(m map[string]int) int {
	values := getValues(m)
	return intutils.Reduce(values, func (acc int, val int) int {
		return intutils.Max(acc, val)
	})
}

func getValues(m map[string]int) []int {
	results := make([]int, 0)

	for _, value := range m {
		results = append(results, value)
	}

	return  results
}

func parseRow(row []string) (string, int, string, string, int, error) {
	variable := row[0]
	operation := row[1]
	valueRaw := row[2]

	conditionVariable := row[4]
	condition := row[5]
	conditionValueRaw := row[6]

	tempValue, err := strconv.ParseInt(valueRaw, 10, 0)
	if err != nil {
		return "", 0, "", "", 0, err
	}
	value := int(tempValue)

	tempConditionValue, err := strconv.ParseInt(conditionValueRaw, 10, 0)
	if err != nil {
		return "", 0, "", "", 0, err
	}
	conditionValue := int(tempConditionValue)

	if operation == "dec" {
		value *= -1
	}

	return variable, value, conditionVariable, condition, conditionValue, err
}

func needPerformOperation(value int, condition string, conditionValue int) bool {
	switch (condition) {
	case "==":
		return value == conditionValue
	case "!=":
		return value != conditionValue
	case ">":
		return value > conditionValue
	case ">=":
		return value >= conditionValue
	case "<":
		return value < conditionValue
	case "<=":
		return value <= conditionValue
	default:
		return false
	}
}

func executeRow(variableMap map[string]int, row []string) error {
	variable, value, conditionVariable, condition, conditionValue, err := parseRow(row)
	if err != nil {
		return err
	}

	if _, exists := variableMap[variable]; !exists {
		variableMap[variable] = 0
	}
	if _, exists := variableMap[conditionVariable]; !exists {
		variableMap[conditionVariable] = 0
	}

	conditionVariableValue := variableMap[conditionVariable]

	performOperation := needPerformOperation(conditionVariableValue, condition, conditionValue)

	if performOperation {
		variableMap[variable] += value
	}

	return nil
}
