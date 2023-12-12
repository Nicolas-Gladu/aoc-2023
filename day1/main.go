package main

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/Nicolas-Gladu/aoc-2023/shared"
)

const ERR_NOT_FOUND = "character not found"

func main() {
	challenge, err := shared.NewDayChallenge(1)
	if err != nil {
		panic(err)
	}
	var total int
	for _, input := range challenge.Inputs {
		value := findCalibration(input)
		total += value
	}
	println(total)
}

func findCalibration(input string) int {
	firstChar, err := findFirstChar(input)
	if err != nil {
		panic(err)
	}

	secondChar, err := findLastChar(input)
	if err != nil {
		panic(err)
	}

	calibration, err := strconv.ParseUint(firstChar+secondChar, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(calibration)
}

func findFirstChar(input string) (string, error) {
	for _, c := range input {
		if unicode.IsDigit(c) {
			return string(c), nil
		}
	}
	return "", errors.New(ERR_NOT_FOUND)
}

func findLastChar(input string) (string, error) {
	chars := []rune(input)
	for i := len(chars) - 1; i >= 0; i-- {
		if unicode.IsDigit(chars[i]) {
			return string(chars[i]), nil
		}
	}
	return "", errors.New(ERR_NOT_FOUND)
}
