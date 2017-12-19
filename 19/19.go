package main

import (
	"fmt"
	"../utils/argparse"
	"errors"
	"strings"
)

type Matrix []string

type Point struct {
	row int
	column int
}

func (p *Point) getValue(matrix Matrix) byte {
	return matrix[p.row][p.column]
}

func (p *Point) in(matrix Matrix) bool {
	if p.row < 0 {
		return false
	}
	if p.row >= len(matrix) {
		return false
	}
	if (p.column < 0) {
		return false
	}
	if p.column >= len(matrix[p.row]) {
		return false
	}
	return true
}

func (p *Point) up() Point {
	return Point{p.row - 1, p.column}
}

func (p *Point) down() Point {
	return Point{p.row + 1, p.column}
}

func (p *Point) left() Point {
	return Point{p.row, p.column - 1}
}

func (p *Point) right() Point {
	return Point{p.row, p.column + 1}
}

func (p *Point) isEmpty(matrix Matrix) bool {
	return p.getValue(matrix) == EMPTY
}

func (p *Point) isWay(matrix Matrix) bool {
	return p.getValue(matrix) == VERTICAL || p.getValue(matrix) == HORIZONTAL
}

func (p *Point) isCross(matrix Matrix) bool {
	return p.getValue(matrix) == CROSS
}

func (p *Point) isLetter(matrix Matrix) bool {
	return !p.isEmpty(matrix) && !p.isWay(matrix) && !p.isCross(matrix)
}

func (p *Point) substract(other Point) Point {
	return Point{p.row - other.row, p.column-other.column}
}

func (p *Point) add(other Point) Point {
	return Point{p.row + other.row, p.column + other.column}
}

func (p *Point) next(matrix Matrix, prev Point) (Point, bool, error) {
	if p.isWay(matrix) {
		move := p.substract(prev)
		return p.add(move), true, nil
	}
	if p.isLetter(matrix) {
		move := p.substract(prev)
		ret := p.add(move)
		return ret, ret.in(matrix) && !ret.isEmpty(matrix), nil
	}
	if p.isCross(matrix) {
		candidates := []Point{p.up(), p.down(), p.left(), p.right()}
		for _, candidate := range candidates {
			if candidate == prev {
				continue
			} else if candidate.in(matrix) && !candidate.isEmpty(matrix) {
				return candidate, true, nil
			}
		}
		return Point{-1, -1}, false, nil
	}
	return Point{-1,-1}, false, errors.New("Invalid position.")
}

const (
	VERTICAL = '|'
	HORIZONTAL = '-'
	CROSS = '+'
	EMPTY = ' '
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 19: A Series of Tubes. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadStringColumn(1)
	if err != nil {
		return err
	}

	letters, steps, err := trip(matrix)
	if err != nil {
		return err
	}

	fmt.Println("Path:", string(letters))
	fmt.Println("Steps:", steps)

	return nil
}

func trip(matrix Matrix) ([]byte, int, error) {
	letters := make([]byte, 0)

	current, err := findStart(matrix)
	if err != nil {
		return nil, -1, err
	}

	previous := current.up()
	steps := 1

	for next, founded, err := current.next(matrix, previous); founded; next, founded, err = current.next(matrix, previous) {
		if err != nil {
			return nil, -1, err
		}

		steps += 1
		previous = current
		current = next
		if current.isLetter(matrix) {
			letters = append(letters, current.getValue(matrix))
		}
	}

	return letters, steps, err
}

func findStart(matrix Matrix) (point Point, err error) {
	point.row = 0
	top := matrix[point.row]
	point.column = strings.IndexByte(top, VERTICAL)
	if point.column < 0 {
		err = errors.New("start point not found")
	}
	return
}
