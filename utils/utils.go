package utils

import (
	"bufio"
	"os"
)

func IsRuneADigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func ParseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return input
}
