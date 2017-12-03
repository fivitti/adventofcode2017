// Advent of code 2017. Day 2.
// Author: Slawomir 'Fivitti' Figiel
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const matrixSize = 10000

var matrix [matrixSize][matrixSize]int

func c2m(x, y int) (int, int) {
	center := matrixSize / 2
	return x + center, y + center
}

func get(x, y int) int {
	matrixX, matrixY := c2m(x, y)
	return matrix[matrixX][matrixY]
}

func set(x, y, value int) {
	matrixX, matrixY := c2m(x, y)
	matrix[matrixX][matrixY] = value
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Find first greater than. Require n(int)")
		return
	}

	rawExpected, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Println("Invalid argument.")
		return
	}

	expected := int(rawExpected)

	n, err := genMatrix(func(val int) bool {
		return val > expected
	})
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Printf("Finded: %d", n)
}

func nextCoords(boxSize, x, y int) (int, int, error) {
	halfBoxSize := boxSize / 2

	if x == halfBoxSize && y == -halfBoxSize {
		return x + 1, y, nil
	}

	// Up
	if x == halfBoxSize && y < halfBoxSize {
		return x, y + 1, nil
	}
	// Left
	if x > -halfBoxSize && y == halfBoxSize {
		return x - 1, y, nil
	}
	// Down
	if x == -halfBoxSize && y > -halfBoxSize {
		return x, y - 1, nil
	}
	// Right
	if x < halfBoxSize && y == -halfBoxSize {
		return x + 1, y, nil
	}

	return 0, 0, errors.New("Unsupported operation.")
}

func genMatrix(stopCond func(int) bool) (int, error) {
	set(0, 0, 1)

	var err error
	boxSize := 3
	x, y := 1, 0

	for boxSize < matrixSize {
		halfBoxSize := boxSize / 2

		for x >= -halfBoxSize && x <= halfBoxSize && y >= -halfBoxSize && y <= halfBoxSize {
			neightbourValue := 0
			for idxX := Max(x-1, -halfBoxSize); idxX <= Min(x+1, halfBoxSize); idxX++ {
				for idxY := Max(y-1, -halfBoxSize); idxY <= Min(y+1, halfBoxSize); idxY++ {
					if idxX == x && idxY == y {
						continue
					}
					neightbourValue += get(idxX, idxY)
				}
			}
			set(x, y, neightbourValue)

			if stopCond(neightbourValue) {
				return neightbourValue, nil
			}

			x, y, err = nextCoords(boxSize, x, y)
			if err != nil {
				return 0, err
			}
		}

		boxSize += 2
	}

	return get(matrixSize/2-1, -matrixSize/2-1), nil
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
