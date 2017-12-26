package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
)

const (
	Clean = "."
	Weakend = "W"
	Infected = "#"
	Flagged = "F"
)

type World struct {
	state map[int]map[int]string
}

func (w World) get(row, column int) string {
	stateRow, exist := w.state[row]
	if !exist {
		return Clean
	}
	stateValue, existsValue := stateRow[column]
	if !existsValue {
		return Clean
	}
	return stateValue
}

func (w *World) set(row, column int, value string) {
	if _, rowExists := w.state[row]; !rowExists {
		w.state[row] = make(map[int]string)
	}
	w.state[row][column] = value
}

func (w World) infectCount() int {
	count := 0
	for _, row := range w.state {
		for _, value := range row {
			if value == Infected  {
				count += 1
			}
		}
	}
	return count
}

func (w World) print(currentRow, currentColumn int) {
	minimum := intutils.MaxInt
	maximum := intutils.MinInt

	for rowNum, row := range w.state {
		minimum = intutils.Min(minimum, rowNum)
		maximum = intutils.Max(maximum, rowNum)
		for columnNum, _ := range row {
			minimum = intutils.Min(minimum, columnNum)
			maximum = intutils.Max(maximum, columnNum)
		}
	}

	for row := minimum; row <= maximum; row++ {
		for column := minimum; column <= maximum; column++ {
			if currentRow == row && currentColumn == column {
				fmt.Print("[")
			} else {
				fmt.Print(" ")
			}

			fmt.Print(w.get(row, column))

			if currentRow == row && currentColumn == column {
				fmt.Print("]")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type Move struct {
	row int
	column int
}

func (m *Move) left() {
	m.row, m.column = -m.column, m.row
}

func (m *Move) right() {
	m.row, m.column = m.column, -m.row
}

func (m *Move) reverse() {
	m.row, m.column = -m.row, -m.column
}

func (m Move) next(row int, column int) (int, int) {
	return row + m.row, column + m.column
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 22: Sporifica Virus. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadStringMatrix(1, "")
	if err != nil {
		return err
	}

	world := World{make(map[int]map[int]string)}
	readInput(matrix, world)

	currentRow := 0
	currentColumn := 0
	ROUNDS := 10000000
	move := Move{row:-1, column:0}

	infectBurst := 0

	for i := 0; i < ROUNDS; i++ {
		current := world.get(currentRow, currentColumn)
		nextState := Clean
		if current == Infected {
			move.right()
			nextState = Flagged
		} else if current == Clean {
			move.left()
			nextState = Weakend
		} else if current == Flagged {
			move.reverse()
			nextState = Clean
		} else if current == Weakend {
			nextState = Infected
		}
		world.set(currentRow, currentColumn, nextState)
		currentRow, currentColumn = move.next(currentRow, currentColumn)

		if nextState == Infected {
			infectBurst += 1
		}
		//world.print(currentRow, currentColumn)
		//fmt.Println()
	}

	infected := world.infectCount()
	fmt.Println("Infected:", infected)
	fmt.Print("Infect burst:", infectBurst)

	return nil
}

func readInput(matrix [][]string, world World) {
	rowCount := len(matrix)
	columnCount := len(matrix[0])
	offsetRow := rowCount / 2
	offsetColumn := columnCount / 2
	for rowId, row := range matrix {
		for columnId, cell := range row {
			world.set(rowId - offsetRow, columnId - offsetColumn, cell)
		}
	}
}