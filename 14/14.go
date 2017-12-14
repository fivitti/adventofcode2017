package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../10/knothash"
)

type Partition [][]bool
type RegionMatrix [][]int

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 14: Disk Defragmentation. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	hashBase, err := argparse.ReadStringRow(1)
	if err != nil {
		return err
	}

	const GRID_SIZE = 128

	partition := make(Partition, GRID_SIZE)

	for i := 0; i < GRID_SIZE; i++ {
		hashSeed := fmt.Sprintf("%s-%d", hashBase, i)
		row := calculateRowUsage(hashSeed)
		partition[i] = row
	}

	usedSquares := countUsed(partition)

	fmt.Println("Used squares:", usedSquares)

	regionMatrix := makeRegionMatrix(partition)
	findRegions(regionMatrix)
	regionCount := intutils.MaximumListList(regionMatrix)

	fmt.Println("Region count:", regionCount)

	return nil
}

func calculateRowUsage(hashSeed string) []bool {
	hash := knothash.Hash(hashSeed)
	bits := intutils.BytesToBits(hash)
	return bits
}

func countUsed(partition Partition) int {
	count := 0

	for _, row := range partition {
		for _, square := range row {
			if square {
				count += 1
			}
		}
	}

	return count
}

const UNUSED_SQUARE = -1
const UNKNOWN_REGION_SQUARE = 0

func makeRegionMatrix(partition Partition) RegionMatrix {
	result := make(RegionMatrix, len(partition))

	for rowIdx, row := range partition {
		resultRow := make([]int, len(row))

		for squareId, squareUsed := range row {
			if !squareUsed {
				resultRow[squareId] = UNUSED_SQUARE
			}
		}

		result[rowIdx] = resultRow
	}

	return result
}

func findRegions(regionMatrix RegionMatrix) {
	region := 1

	for x, row := range regionMatrix {
		for  y, square := range row {
			if square == UNKNOWN_REGION_SQUARE {
				markRegion(regionMatrix, x, y, region)
				region += 1
			}
		}
	}
}

func markRegion(regionMatrix RegionMatrix, x int, y int, region int) {
	if regionMatrix[x][y] != UNKNOWN_REGION_SQUARE {
		return
	}

	regionMatrix[x][y] = region

	// Left
	if y > 0 {
		markRegion(regionMatrix, x, y - 1, region)
	}
	// Right
	if y < len(regionMatrix[x]) - 1 {
		markRegion(regionMatrix, x, y + 1, region)
	}
	// Up
	if x > 0 {
		markRegion(regionMatrix, x - 1, y, region)
	}
	//Down
	if x < len(regionMatrix) - 1 {
		markRegion(regionMatrix, x + 1, y, region)
	}
}
