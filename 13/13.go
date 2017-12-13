package main

import (
	"../utils/argparse"
	"../utils/intutils"
	"fmt"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 13: Packet Scanners: /path/to/file(string) |", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}

	input, err := argparse.ReadIntMatrix(1, ": ")
	if err != nil {
		return err
	}

	layers := make(map[int]int)

	for _, row := range input {
		depth := row[0]
		range_ := row[1]
		layers[depth] = range_
	}

	layerCount := input[len(input) - 1][0] + 1

	caughs := trip(layers, layerCount, 0)
	caughtResume := resume(layers, caughs)

	fmt.Println("Caught count:", caughtResume)

	delay := tripWithoutCaught(layers, layerCount)
	fmt.Println("Delay to no caught:", delay)

	return nil
}

func trip(input map[int]int, layerCount int, delay int) []int {
	result := make([]int, 0)

	for layer := 0; layer < layerCount; layer++ {
		time := layer + delay
		layerRange, isSecure := input[layer]
		if isSecure {
			firewallWindow := getPosition(layerRange, time)
			if firewallWindow == 0 {
				result = append(result, layer)
			}
		}
	}

	return result
}

func tripWithoutCaught(input map[int]int, layerCount int) int {
	delay := 0
	for len(trip(input, layerCount, delay)) != 0 {
		delay += 1
	}
	return delay
}

func resume(input map[int]int, caughtLayers []int) int {
	return intutils.Reduce(caughtLayers, func (acc int, layer int) int {
		return acc + layer * input[layer]
	} )
}

func getPosition(range_, time int) int {
	position := time % (range_ - 1)
	round := time / (range_ - 1)

	if round % 2 == 1 {
		position = (range_ - 1) - position
	}

	return position
}
