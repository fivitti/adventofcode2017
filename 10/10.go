package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"./knothash"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 10: Knot Hash. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	row, err := argparse.ReadStringRow(1)
	if err != nil {
		return err
	}

	hashHex := intutils.BytesToHexes(knothash.Hash(row), "")

	fmt.Println(hashHex)

	return nil
}