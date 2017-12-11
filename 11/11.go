package main

import (
	"../utils/intutils"
	"../utils/argparse"
	"strings"
	"fmt"
)

const (
	X = 0
	Y = 1
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 11: Hex Ed: /path/to/file(string) |", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}

	inputSymbolsRaw, err := argparse.ReadStringRow(1)
	if err != nil {
		return err
	}

	inputSymbols := strings.Split(inputSymbolsRaw, ",")

	moves := map[string][2]int {
		"n" : { 0,  2},
		"ne": { 1,  1},
		"se": { 1, -1},
		"s" : { 0, -2},
		"sw": {-1, -1},
		"nw": {-1,  1},
	}

	sumX, sumY := sumVectors(inputSymbols, moves)
	outputSymbols := splitToVectors(sumX, sumY, moves)

	//fmt.Println("Shortest way:", outputSymbols)
	fmt.Println("Needed steps:", len(outputSymbols))

	stepsFromFurthest := furthestSteps(inputSymbols, moves)
	fmt.Println("Steps to futhest position:", stepsFromFurthest)

	return nil
}

func furthestSteps(arr []string, moves map[string][2]int) int {
	x := 0
	y := 0

	furthestSteps := 0

	for _, symbol := range arr {
		symbolMove := moves[symbol]
		moveX := symbolMove[X]
		moveY := symbolMove[Y]
		x += moveX
		y += moveY

		steps := len(splitToVectors(x, y, moves))

		if furthestSteps < steps {
			furthestSteps = steps
		}
	}

	return furthestSteps
}

func sumVectors(arr []string, moves map[string][2]int) (int, int) {
	x := 0
	y := 0

	for _, symbol := range arr {
		symbolMove := moves[symbol]
		moveX := symbolMove[X]
		moveY := symbolMove[Y]
		x += moveX
		y += moveY
	}

	return x, y
}

func splitToVectors(x, y int, moves map[string][2]int) []string {
	result := make([]string, 0)

	for x != 0 && y != 0 {
		bestMoveSymbol := getBestMove(x, y, moves)
		symbolMove := moves[bestMoveSymbol]
		moveX, moveY := symbolMove[X], symbolMove[Y]
		x -= moveX
		y -= moveY

		result = append(result, bestMoveSymbol)
	}

	return result
}

func getBestMove(x, y int, moves map[string][2]int) string {
	minScore := intutils.MaxInt
	var minSymbol string

	for symbol, moves := range moves {
		moveX := moves[X]
		moveY := moves[Y]

		movedX := x - moveX
		movedY := y - moveY

		score := vectorScore(movedX, movedY)
		if score < minScore {
			minScore = score
			minSymbol = symbol
		}
	}

	return minSymbol
}

func vectorScore(x, y int) int {
	return intutils.Pow(x, 2) + intutils.Pow(y, 2)
}