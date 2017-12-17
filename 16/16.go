package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../utils/parsers"
	"strings"
	"errors"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 16: Permutation Promenade. Parameters: /path/to/file(string)", err)
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

	const DANCERS = 16
	const ROUNDS = 1000000000

	dancerState := makeDancers(DANCERS)
	commands := strings.Split(row, ",")
	findFinishState(dancerState, commands, ROUNDS)

	fmt.Println(string(intutils.IntsToBytes(dancerState)))

	return nil
}

func findFinishState(dancerState []int, commands []string, rounds int) error {
	cycleSize, err := findCycle(dancerState, commands)
	if err != nil {
		return err
	}
	stateNum := rounds % cycleSize
	for i := 0; i < stateNum; i++ {
		err := dance(dancerState, commands)
		if err != nil {
			return err
		}
	}

	return nil
}

func findCycle(initDancerState []int, commands []string) (int, error) {
	dancerState := make([]int, len(initDancerState))
	copy(dancerState, initDancerState)

	cycleSize := -1

	const SEARCH_DEPTH = 1000

	for i := 0; i < SEARCH_DEPTH; i++ {
		err := dance(dancerState, commands)
		if err != nil {
			return -1, err
		}

		if compareArrays(initDancerState, dancerState) {
			cycleSize = i + 1
		}
	}

	if cycleSize == -1 {
		return -1, errors.New("Cannot find cycle.")
	}
	return cycleSize, nil
}

func dance(dancerState []int, commands []string) error {
	for _, command := range commands {
		err := performCommand(dancerState, command)
		if err != nil {
			return err
		}
	}

	return nil
}

func compareArrays(first []int, second []int) bool {
	for idx, val := range first {
		if second[idx] != val {
			return  false
		}
	}
	return true
}

func makeDancers(dancerCount int) []int {
	dancers := make([]int, dancerCount)
	for dancerId := 0; dancerId < dancerCount; dancerId++ {
		dancers[dancerId] = 'a' + dancerId
	}
	return dancers
}

func performCommand(dancers []int, command string) error {
	commandSign := command[0]
	switch (commandSign) {
	case 's':
		spinSize, err := intutils.ParseInt(command[1:])
		if err != nil {
			return err
		}
		performSpin(dancers, spinSize)
	case 'x':
		positions, err :=  parsers.StringsToNumbers(strings.Split(command[1:], "/"))
		if err != nil {
			return err
		}
		performExchange(dancers, positions[0], positions[1])
	case 'p':
		programs := strings.Split(command[1:], "/")
		performPartner(dancers, int(programs[0][0]), int(programs[1][0]))
	default:
		return errors.New("Invalid command")
	}
	return nil
}

func performSpin(dancers []int, size int) {
	shiftedDancers := intutils.Shift(dancers, size)
	for idx, val := range shiftedDancers {
		dancers[idx] = val
	}
}

func performExchange(dancers []int, positionA int, positionB int) {
	dancerA := dancers[positionA]
	dancers[positionA] = dancers[positionB]
	dancers[positionB] = dancerA
}

func performPartner(dancers []int, first int, second int) {
	dancerA := intutils.IndexOf(dancers, first)
	dancerB := intutils.IndexOf(dancers, second)
	performExchange(dancers, dancerA, dancerB)
}