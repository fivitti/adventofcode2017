package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
)

type Component struct {
	begin int
	end int
}

func (c Component) getStrenght() int {
	return c.begin + c.end
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 24: Electromagnetic Moat. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadIntMatrix(1, "/")
	if err != nil {
		return err
	}
	components := readInput(matrix)
	strongestBridge := findStrongestBridge(components, 0)
	strength := getStrength(strongestBridge)

	fmt.Println("Strongest bridge:", strength)

	return nil
}

func readInput(matrix [][]int) []Component {
	result := make([]Component, 0)
	for _, rawComponent := range matrix {
		result = append(result, Component{begin:rawComponent[0], end:rawComponent[1]})
	}
	return result
}

func findCompatible(components []Component, condition int) []Component {
	result := make([]Component, 0)
	for _, component := range components {
		if component.begin == condition || component.end == condition {
			result = append(result, component)
		}
	}
	return result
}

func getStrength(components []Component) int {
	strength := 0
	for _, component := range components {
		strength += component.getStrenght()
	}
	return strength
}

func without(components []Component, componentToRemove Component) []Component {
	for idx, component := range components {
		if component == componentToRemove {
			return copyWithout(components, idx)
		}
	}
	return components
}

func copyWithout(components []Component, toIgnore int) []Component {
	result := make([]Component, 0)
	for idx, component := range components {
		if idx != toIgnore {
			result = append(result, component)
		}
	}
	return result
}

func findStrongestBridge(components []Component, begin int) []Component {
	if len(components) == 0 {
		return nil
	}

	candidates := findCompatible(components, begin)

	strongestBridgeStrength := intutils.MinInt
	var strongestBridge []Component
	strongestBridgeLength := intutils.MinInt

	for _, candidate := range candidates {
		candidateStrength := candidate.getStrenght()
		candidateBridge := []Component{candidate}
		candidateLenght := 1

		endToConnect := candidate.end
		if candidate.end == begin {
			endToConnect = candidate.begin
		}

		strongestSubBridge := findStrongestBridge(without(components, candidate), endToConnect)
		if strongestSubBridge != nil {
			candidateBridge = append(candidateBridge, strongestSubBridge...)
			candidateStrength = getStrength(candidateBridge)
			candidateLenght = len(candidateBridge)
		}
		if candidateLenght > strongestBridgeLength ||
			(candidateLenght == strongestBridgeLength && candidateStrength > strongestBridgeStrength) {
			strongestBridgeStrength = candidateStrength
			strongestBridge = candidateBridge
			strongestBridgeLength = candidateLenght
		}
	}

	return strongestBridge
}