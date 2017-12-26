package src

import (
	"errors"
	"fmt"
)

func CreateDefaultGrid() Grid {
	tile := [][]byte{
		{ '.', '#', '.'},
		{ '.', '.', '#'},
		{ '#', '#', '#'},
	}
	return Grid{tiles:[][]Tile{{tile}}}
}

type Grid struct {
	tiles [][]Tile
}

func (g *Grid) Transform(patterns []TilePattern) error {
	for tileRowIdx, tileRow := range g.tiles {
		for tileColumnIdx, tile := range tileRow {
			newTile, correct := transformTile(tile, patterns)
			if !correct {
				return errors.New("Cannot transform tile.")
			}
			g.tiles[tileRowIdx][tileColumnIdx] = newTile
		}
	}

	g.tiles = Merge(g.tiles).Split()

	return nil
}

func (g Grid) GetSize() int {
	gridSize := len(g.tiles)
	if gridSize == 0 {
		return 0
	}
	tileSize := g.tiles[0][0].GetSize()
	return gridSize * tileSize
}

func (g Grid) GetOnPixelCount() int {
	count := 0
	for _, tileRow := range g.tiles {
		for _, tile := range tileRow {
			count += tile.Count('#')
		}
	}
	return count
}

func transformTile(tile Tile, patterns []TilePattern) (Tile, bool) {
	for _, pattern := range patterns {
		if pattern.IsMatch(tile) {
			return pattern.GetOut(), true
		}
	}
	return Tile{}, false
}

func (g Grid) Print() {
	tilesInRow := len(g.tiles)

	gridSize := g.GetSize()

	for tileRowIdx, tileRow := range g.tiles {
		tileSize := tileRow[0].GetSize()
		for row := 0; row < tileSize; row++ {
			for tileColumnIdx, tile := range tileRow {
				for column := 0; column < tileSize; column++ {
					value := tile[row][column]
					fmt.Print(string(value))
				}

				if tileColumnIdx != tilesInRow - 1 {
					fmt.Print("|")
				}
			}
			fmt.Println()
		}

		if tileRowIdx != tilesInRow - 1 {
			for i := 0; i < gridSize; i++ {
				if i%tileSize == 0 && i != 0 {
					fmt.Print("+")
				}
				fmt.Print("-")
			}
			fmt.Println()
		}
	}
}