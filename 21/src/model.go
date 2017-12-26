package src

import (
	"../../utils/intutils"
	"strings"
)

type Tile [][]byte

func (t Tile) GetSize() int {
	return len(t)
}

func (t Tile) Count(condition byte) int {
	count := 0
	for _, row := range t {
		for _, value := range row {
			if value == condition {
				count += 1
			}
		}
	}
	return count
}

func (t Tile) Split() [][]Tile {
	size := t.GetSize()

	if size % 2 == 0 {
		return t.splitBySize(2)
	} else /*if size % 3 == 0 */ {
		return t.splitBySize(3)
	}
}

func (t Tile) splitBySize(targetSize int) [][]Tile {
	size := t.GetSize()
	result := make([][]Tile, size / targetSize)

	for startRow := 0; startRow < size; startRow += targetSize {
		result[startRow / targetSize] = make([]Tile, size / targetSize)
		for startColumn := 0; startColumn < size; startColumn += targetSize {
			tile := t.getSubTile(startRow, startColumn, targetSize)
			result[startRow/targetSize][startColumn/targetSize] = tile
		}
	}
	return result
}

func Merge(tiles [][]Tile) Tile {
	firstTile := tiles[0][0]
	tileSize := firstTile.GetSize()

	boxSize := len(tiles) * tileSize

	result := make(Tile, boxSize)
	for i := 0; i < boxSize; i++ {
		result[i] = make([]byte, boxSize)
	}

	for tilesRowIdx, tilesRow := range tiles {
		for tilesColumnIdx, tile := range tilesRow {
			for rowIdx, row := range tile {
				resultRow := tilesRowIdx * tileSize + rowIdx
				for columnIdx, value := range row {
					resultColumn := tilesColumnIdx * tileSize + columnIdx
					result[resultRow][resultColumn] = value
				}
			}
		}
	}
	return result
}

func (t Tile) getSubTile(x, y, size int) Tile {
	result := make(Tile, size)
	for row := 0; row < size; row++ {
		result[row] = make([]byte, size)
		for column := 0; column < size; column++ {
			result[row][column] = t[row + x][column + y]
		}
	}
	return result
}

func CreatePattern(input string) TilePattern {
	parts := strings.Split(input, " => ")
	patternPart, outputPart := parts[0], parts[1]
	return TilePattern{pattern:readInputArray(patternPart), output:readInputArray(outputPart)}
}

func readInputArray(input string) [][]byte {
	rows := strings.Split(input, "/")
	result := make([][]byte, len(rows))
	for rowIdx, row := range rows {
		result[rowIdx] = make([]byte, len(row))
		for columnIdx, value := range row {
			result[rowIdx][columnIdx] = byte(value)
		}
	}
	return result
}

type TilePattern struct {
	pattern pattern
	output Tile
}

func (tp TilePattern) IsMatch(tile Tile) bool {
	return tp.pattern.IsMatch(tile)
}

func (tp TilePattern) GetOut() Tile {
	return tp.output
}

type pattern [][]byte

func (p pattern) getSize() int {
	return len(p)
}

func (p pattern) turn() pattern {
	size := p.getSize()
	result := make([][]byte, size)
	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
	}

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			newColumn := row
			newRow := size - column - 1
			result[newRow][newColumn] = p[row][column]
		}
	}

	return result
}

func (p pattern) flip() pattern {
	result := make(pattern, len(p))
	copy(result, p)
	for _, row := range result {
		intutils.ReverseBytesInPlace(row)
	}
	return result
}

func (p pattern) isEquals(tile Tile) bool {
	for rowIdx, row := range p {
		for columnIdx, value := range row {
			if value != tile[rowIdx][columnIdx] {
				return false
			}
		}
	}
	return true
}

func (p pattern) matchAnyTurn(tile Tile) bool {
	toCheck := p
	for i := 0; i < 4; i++ {
		if toCheck.isEquals(tile) {
			return true
		}
		toCheck = toCheck.turn()
	}
	return false
}

func (p pattern) IsMatch(tile Tile) bool {
	if p.getSize() != tile.GetSize() {
		return false
	}

	if p.matchAnyTurn(tile) {
		return true
	}
	return p.flip().matchAnyTurn(tile)
}