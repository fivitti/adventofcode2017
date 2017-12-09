package main

import "../utils/intutils"
import "../utils/stringutils"
import "../utils/argparse"
import "fmt"

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 9: Stream Processing. Parameters: /path/to/file(string) ", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}

	input, err := argparse.ReadStringRow(1)
	if err != nil {
		return err
	}

	rowGroupScores, collectedCount := processRow(input)
	sumScores := intutils.Reduce(rowGroupScores, func(acc int, val int) int {
		return acc + val
	})

	fmt.Println("Sum all scores:", sumScores)
	fmt.Println("Garbage collected elements: ", collectedCount)
	return nil
}


func processRow(row string) ([]int, int) {
	chars := stringutils.StringToChars(row)
	chars = performCancelations(chars)
	chars, collectedCount := performGarbageColletion(chars)
	scores := findGroupScores(chars)
	return scores, collectedCount
}

func performCancelations(stream []byte) []byte {
	chars := make([]byte, 0)

	for i := 0; i < len(stream); i++ {
		char := stream[i]
		if char == '!' {
			// Skip next
			i += 1
			continue
		}
		chars = append(chars, char)
	}

	return chars
}

func performGarbageColletion(stream []byte) ([]byte, int) {
	isCollectionInProgress := false
	chars := make([]byte, 0)
	collectedCount := 0

	for _, char := range stream {
		if isCollectionInProgress && char == '>' {
			isCollectionInProgress = false
			continue
		} else if !isCollectionInProgress && char == '<' {
			isCollectionInProgress = true
			continue
		}
		if !isCollectionInProgress {
			chars = append(chars, char)
		} else {
			collectedCount += 1
		}
	}

	return chars, collectedCount
}

func findGroupScores(stream []byte) []int {
	groupScores := make([]int, 0)
	currentScore := 0

	for _, char := range stream {
		if char == '{' {
			currentScore += 1
			groupScores = append(groupScores, currentScore)
		} else if char == '}' {
			currentScore -= 1
		}
	}

	return groupScores
}