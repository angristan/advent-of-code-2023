package main

import (
	"log"
	"strconv"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	sum := ComputeCalibrationSum(input)

	log.Printf("Sum: %d", sum)
}

var digitsSpelledOut = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type CalibrationInput []string

func ComputeCalibrationSum(input CalibrationInput) int {
	sum := 0

	for _, line := range input {
		digitsOnTheLine := []rune{}
		for i, char := range line {
			if utils.IsRuneADigit(char) {
				if len(digitsOnTheLine) < 2 {
					digitsOnTheLine = append(digitsOnTheLine, char)
				} else {
					digitsOnTheLine[1] = char
				}
			} else {
				for spelledOutDigit, intValue := range digitsSpelledOut {
					if i+len(spelledOutDigit) > len(line) {
						continue
					}
					possibleDigit := line[i : i+len(spelledOutDigit)]
					if possibleDigit == spelledOutDigit {
						if len(digitsOnTheLine) < 2 {
							digitsOnTheLine = append(digitsOnTheLine, rune(intValue+'0'))
						} else {
							digitsOnTheLine[1] = rune(intValue + '0')
						}
					}
				}
			}
		}

		if len(digitsOnTheLine) == 1 {
			digitsOnTheLine = append(digitsOnTheLine, digitsOnTheLine[0])
		}

		number, err := strconv.Atoi(string(digitsOnTheLine))
		if err != nil {
			panic(err)
		}

		sum += number
	}

	return sum
}
