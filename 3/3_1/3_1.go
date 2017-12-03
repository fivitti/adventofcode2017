// Advent of code 2017. Day 2.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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
	distanceX := Abs(startX - endX)
	distanceY := Abs(startY - endY)
	return distanceX + distanceY
}

func FindSquare(n int) int {
	return Ceil((math.Sqrt(float64(n)) + 1) / 2)
}

func FindBoxSize(square int) int {
	return 1 + (square-1)*2
}

func FindMaximum(square int) int {
	return Pow(FindBoxSize(square), 2)
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

	toAddY := Min(max_-startY, position)
	position = Max(position-toAddY, 0)
	toSubstractX := Min(boxSize-1, position)
	position = Max(position-toSubstractX, 0)
	toSubstractY := Min(boxSize-1, position)
	position = Max(position-toSubstractY, 0)
	toAddX := Min(boxSize-1, position)
	position = Max(position-toAddX, position)

	return startX - toSubstractX + toAddX, startY - toSubstractY + toAddY
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func Ceil(x float64) int {
	return int(math.Ceil(x))
}

func Pow(x int, f int) int {
	return int(math.Pow(float64(x), float64(f)))
}
