package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"strings"
	"errors"
)

type Direction int

const (
	LEFT Direction= -1
	RIGHT Direction = 1
)

type Behaviour struct {
	direction Direction
	value int
	nextState string
	applyValue int
}

type State struct {
	name string
	behaviours []Behaviour
}

func (s State) getBehaviour(currentValue int) (Behaviour, error) {
	for _, behaviour := range s.behaviours {
		if behaviour.applyValue == currentValue {
			return behaviour, nil
		}
	}
	return Behaviour{}, errors.New("unknown behaviour")
}

type Tape struct {
	current int
	values map[int]int
}

func (t *Tape) move(direction Direction) {
	t.current += int(direction)
}

func (t Tape) getCurrentValue() int {
	value, ok := t.values[t.current]
	if !ok {
		return 0
	}
	return value
}

func (t *Tape) setCurrentValue(value int) {
	t.values[t.current] = value
}

func (t Tape) getChecksum(condition int) int {
	count := 0
	for _, value := range t.values {
		if value == condition {
			count += 1
		}
	}
	return count
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 25: The Halting Problem", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadStringMatrix(1, " ")
	if err != nil {
		return err
	}

	tape := Tape{values:make(map[int]int), current:0}
	states, currentState, steps, err := readInput(matrix)
	if err != nil {
		return err
	}

	for i := 0; i < steps; i++ {
		state, ok := states[currentState]
		if !ok {
			return errors.New("unknown state")
		}
		currentValue := tape.getCurrentValue()
		behaviour, err := state.getBehaviour(currentValue)
		if err != nil {
			return err
		}
		tape.setCurrentValue(behaviour.value)
		tape.move(behaviour.direction)
		currentState = behaviour.nextState
	}

	checksum := tape.getChecksum(1)
	fmt.Println("Checksum:", checksum)

	return nil
}

func readInput(matrix [][]string) (states map[string]State, beginState string, steps int, err error) {
	beginStateRaw := matrix[0][3]
	beginState = strings.TrimRight(beginStateRaw, ".")

	stepsRaw := matrix[1][5]
	steps, err = intutils.ParseInt(stepsRaw)
	if err != nil {
		return
	}

	states = make(map[string]State)

	for i := 3; i < len(matrix); i += 10 {
		var state State
		state, err = readInputState(matrix[i:i+10])
		if err != nil {
			return
		}
		states[state.name] = state
	}

	return
}

func readInputState(matrix [][]string) (state State, err error) {
	nameRow := matrix[0]
	name := nameRow[2]
	name = strings.TrimRight(name, ":")

	behaviours := make([]Behaviour, 0)
	for i := 1; i < len(matrix) - 1; i += 4 {
		var behaviour Behaviour
		behaviour, err = readInputBehaviour(matrix[i:i+4])
		if err != nil {
			return
		}
		behaviours = append(behaviours, behaviour)
	}

	state = State{behaviours:behaviours, name:name}
	return
}

func readInputBehaviour(matrix [][]string) (behaviour Behaviour, err error) {
	applyValueRow := matrix[0]
	offset := 2
	applyValueRaw := applyValueRow[5+offset]
	applyValueRaw = strings.TrimRight(applyValueRaw, ":")
	applyValue, err := intutils.ParseInt(applyValueRaw)
	if err != nil {
		return
	}

	offset += 2
	rawValue := matrix[1][4+offset]
	rawValue = strings.TrimRight(rawValue, ".")
	value, err := intutils.ParseInt(rawValue)
	if err != nil {
		return
	}

	rawDirection := matrix[2][6+offset]
	rawDirection = strings.TrimRight(rawDirection, ".")
	var direction Direction
	if rawDirection == "left" {
		direction = LEFT
	} else if rawDirection == "right" {
		direction = RIGHT
	} else {
		err = errors.New("unknown direction")
		return
	}

	nextState := matrix[3][4+offset]
	nextState = strings.TrimRight(nextState, ".")

	behaviour = Behaviour{direction:direction, value:value, nextState:nextState, applyValue:applyValue}
	return
}