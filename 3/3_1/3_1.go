// Advent of code 2017. Day 2.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"../../utils/intutils"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Distance. Require n(int)")
		return
	}

	n, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Println("Invalid argument.")
		return
	}

	x, y := FindCoordinate(int(n))
	distance := FindManhattanDistance(0, 0, x, y)
	fmt.Printf("Coordinate: (%d, %d). Distance: %d.", x, y, distance)
}

func FindManhattanDistance(startX int, startY int, endX int, endY int) int {
	distanceX := intutils.Abs(startX - endX)
	distanceY := intutils.Abs(startY - endY)
	return distanceX + distanceY
}

func FindSquare(n int) int {
	return intutils.Ceil((math.Sqrt(float64(n)) + 1) / 2)
}

func FindBoxSize(square int) int {
	return 1 + (square-1)*2
}

func FindMaximum(square int) int {
	return intutils.Pow(FindBoxSize(square), 2)
}

func FindMinimum(square int) int {
	if square == 1 {
		return 1
	}
	return FindMaximum(square-1) + 1
}

func FindStartCoordinates(square int) (int, int) {
	if square == 1 {
		return 0, 0
	} else if square == 2 {
		return 1, 0
	} else {
		return (square - 2) + 1, -1 * (square - 2)
	}
}

func FindPosition(n int) int {
	square := FindSquare(n)
	min_ := FindMinimum(square)
	return n - min_
}

func FindCoordinate(n int) (int, int) {
	square := FindSquare(n)
	position := FindPosition(n)
	startX, startY := FindStartCoordinates(square)
	boxSize := FindBoxSize(square)

	max_ := boxSize / 2

	toAddY := intutils.Min(max_-startY, position)
	position = intutils.Max(position-toAddY, 0)
	toSubstractX := intutils.Min(boxSize-1, position)
	position = intutils.Max(position-toSubstractX, 0)
	toSubstractY := intutils.Min(boxSize-1, position)
	position = intutils.Max(position-toSubstractY, 0)
	toAddX := intutils.Min(boxSize-1, position)
	position = intutils.Max(position-toAddX, position)

	return startX - toSubstractX + toAddX, startY - toSubstractY + toAddY
}
