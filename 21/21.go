package main

import (
	"fmt"
	"../utils/argparse"
	. "./src"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 21: Fractal Art. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	rawPatterns, err := argparse.ReadStringColumn(1)
	if err != nil {
		return err
	}

	grid := CreateDefaultGrid()
	patterns := make([]TilePattern, len(rawPatterns))
	for idx, rawPatten := range rawPatterns {
		patterns[idx] = CreatePattern(rawPatten)
	}

	const ITERATIONS = 18

	for i := 0; i < ITERATIONS; i++ {
		err := grid.Transform(patterns)
		if err != nil {
			return err
		}
		fmt.Println("Iteration:", i)
		//grid.Print()
		//fmt.Println()
	}


	fmt.Println("Grid size:", grid.GetSize())
	fmt.Println("On pixels:", grid.GetOnPixelCount())

	return nil
}