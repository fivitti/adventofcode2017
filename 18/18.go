package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"errors"
)

type THREAD_STATUS int

const (
	SEND_MESSAGE THREAD_STATUS = iota
	WAIT_FOR_RECEIVE THREAD_STATUS = iota
	NOTHING THREAD_STATUS = iota
)

const (
	SET = "set"
	ADD = "add"
	MULTIPLY = "mul"
	SEND = "snd"
	MODULO = "mod"
	RECEIVE = "rcv"
	JUMP = "jgz"

	NAME_POS = 0
	REGISTER_POS = 1
	ARGUMENT_POS = 2
)

type Thread struct {
	queue []int
	currentPosition int
	attemptsToReceive int
	registers map[string]int
	sendCounter int
}

func (t *Thread) isTerminated() bool {
	return t.attemptsToReceive >= 4
}

func (t *Thread) pop() (int, bool) {
	if len(t.queue) == 0 {
		return  0, false
	}
	var first int
	first, t.queue = t.queue[0], t.queue[1:]
	return first, true
}

func (t *Thread) push(value int) {
	t.queue = append(t.queue, value)
}

func (t *Thread) execute(instructions [][]string, send func (int)) (continue_ bool, err error) {
	instruction := instructions[t.currentPosition]
	instructionName := instruction[NAME_POS]
	register := instruction[REGISTER_POS]

	switch instructionName {
	case SEND:
		{
			t.sendCounter += 1
			value := t.registers[register]
			send(value)
		}
	case SET:
		{
			value, err := getValue(t.registers, instruction[ARGUMENT_POS])
			if err != nil {
				return false, err
			}
			t.registers[register] = value
		}
	case ADD:
		{
			value, err := getValue(t.registers, instruction[ARGUMENT_POS])
			if err != nil {
				return false, err
			}
			t.registers[register] += value
		}
	case MULTIPLY:
		{
			value, err := getValue(t.registers, instruction[ARGUMENT_POS])
			if err != nil {
				return false, err
			}
			t.registers[register] *= value
		}
	case MODULO:
		{
			value, err := getValue(t.registers, instruction[ARGUMENT_POS])
			if err != nil {
				return false, err
			}
			t.registers[register] %= value
		}
	case JUMP:
		{
			condition, err := getValue(t.registers, register)
			if err != nil {
				return false, err
			}
			if condition <= 0 {
				break;
			}

			value, err := getValue(t.registers, instruction[ARGUMENT_POS])
			if err != nil {
				return false, err
			}

			t.currentPosition += value - 1
		}
	case RECEIVE:
		{
			//if t.registers[register] != 0 {
				value, ok := t.pop()
				if !ok {
					t.attemptsToReceive += 1
					return false, nil
				} else {
					t.registers[register] = value
					t.attemptsToReceive = 0
				}
			//}
		}
	}

	t.currentPosition += 1

	return true, nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 18: Duet", err)
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

	t0 := Thread{attemptsToReceive:0, registers:map[string]int{"p": 0}, currentPosition:0, queue:make([]int, 0), sendCounter:0}
	t1 := Thread{attemptsToReceive:0, registers:map[string]int{"p": 1}, currentPosition:0, queue:make([]int, 0), sendCounter:0}

	threadSended := false
	continueThread := true

	for !t0.isTerminated() && !t1.isTerminated() {
		for !threadSended && continueThread {
			continueThread, err = t0.execute(matrix, func(value int) {
				t1.push(value)
				threadSended = true
			})
			if err != nil {
				return err
			}
		}

		threadSended = false
		continueThread = true

		for !threadSended && continueThread {
			continueThread, err = t1.execute(matrix, func (value int) {
				t0.push(value)
				threadSended = true
			})
			if err != nil {
				return err
			}
		}

		threadSended = false
		continueThread = true
	}

	fmt.Println("Send count:", t1.sendCounter)

	return nil
}

func getValue(registers map[string]int, argument string) (int, error) {
	value, err := intutils.ParseInt(argument)
	if err == nil {
		return value, nil
	}
	value, ok := registers[argument]
	if ok {
		return value, nil
	}

	return -1, errors.New("unknown register")
}
