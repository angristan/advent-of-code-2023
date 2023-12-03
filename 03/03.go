package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	Value             string
	DigitsCoordinates []Coordinates
}

type Symbol struct {
	Value       string
	Coordinates Coordinates
}

type Coordinates struct {
	X int
	Y int
}

type EngineSchematic struct {
	Numbers []Number
	Symbols []Symbol
}

/*
The engine schematic (your puzzle input) consists of a visual representation
of the engine. There are lots of numbers and symbols you don't really understand,
but apparently any number adjacent to a symbol, even diagonally, is a "part number"
and should be included in your sum. (Periods (.) do not count as a symbol.)
*/

func ComputeEngineSchematic(input []string) EngineSchematic {
	numbers := make([]Number, 0)

	for y, line := range input {
		tempNumber := Number{}
		for x, char := range line {
			if isRuneADigit(char) {
				if len(tempNumber.DigitsCoordinates) == 0 { // New number
					tempNumber.Value = fmt.Sprintf("%d", int(char-'0'))
					tempNumber.DigitsCoordinates = append(tempNumber.DigitsCoordinates, Coordinates{x, y})
				} else { // existing number
					tempNumber.Value += fmt.Sprintf("%d", int(char-'0'))
					tempNumber.DigitsCoordinates = append(tempNumber.DigitsCoordinates, Coordinates{x, y})
				}
			} else {
				if len(tempNumber.DigitsCoordinates) > 0 { //end of number
					numbers = append(numbers, tempNumber)
					tempNumber = Number{}
				}
			}
		}
		if len(tempNumber.DigitsCoordinates) > 0 { //end of line
			numbers = append(numbers, tempNumber)
			tempNumber = Number{}
		}
	}

	symbols := make([]Symbol, 0)

	for y, line := range input {
		for x, char := range line {
			if !isRuneADigit(char) && char != '.' {
				symbols = append(symbols, Symbol{
					Coordinates: Coordinates{x, y},
					Value:       string(char),
				})
			}
		}
	}

	return EngineSchematic{numbers, symbols}
}

func (es EngineSchematic) GetPartNumbersValues() []int {
	partNumbers := make([]int, 0)
	found := false
	for _, number := range es.Numbers {
		found = false
		for _, adjCoordsOfDigit := range number.GetAllAdjacentCoordinates() {
			for _, symbol := range es.Symbols {
				if adjCoordsOfDigit == symbol.Coordinates {
					intValue, err := strconv.Atoi(number.Value)
					if err != nil {
						panic(err)
					}
					partNumbers = append(partNumbers, intValue)
					found = true
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
	}

	return partNumbers
}

func (nb Number) GetAllAdjacentCoordinates() []Coordinates {
	adjacentCoordinates := make([]Coordinates, 0)

	for _, digits := range nb.DigitsCoordinates {
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X - 1, digits.Y - 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X, digits.Y - 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X + 1, digits.Y - 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X + 1, digits.Y})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X + 1, digits.Y + 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X, digits.Y + 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X - 1, digits.Y + 1})
		adjacentCoordinates = append(adjacentCoordinates, Coordinates{digits.X - 1, digits.Y})
	}

	return adjacentCoordinates
}

func (es EngineSchematic) ComputeSumOfPartNumbers() int {
	partNumbers := es.GetPartNumbersValues()

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	return sum
}

type Gear struct {
	Values []int
}

/*
A gear is any * symbol that is adjacent to exactly two part numbers.
*/
func (es EngineSchematic) GetGears() []Gear {
	asteriskSymbolToNumbers := make(map[Coordinates][]Number)

	for _, symbol := range es.Symbols {
		if symbol.Value == "*" {
			asteriskSymbolToNumbers[symbol.Coordinates] = make([]Number, 0)

			for _, number := range es.Numbers {
				for _, adjCoordsOfDigit := range number.GetAllAdjacentCoordinates() {
					if adjCoordsOfDigit == symbol.Coordinates {
						asteriskSymbolToNumbers[symbol.Coordinates] = append(asteriskSymbolToNumbers[symbol.Coordinates], number)
						break
					}
				}
			}
		}
	}

	gears := make([]Gear, 0)

	for _, numbers := range asteriskSymbolToNumbers {
		if len(numbers) == 2 {
			valueInt1, err := strconv.Atoi(numbers[0].Value)
			if err != nil {
				panic(err)
			}

			valueInt2, err := strconv.Atoi(numbers[1].Value)
			if err != nil {
				panic(err)
			}

			gears = append(gears, Gear{Values: []int{valueInt1, valueInt2}})
		}
	}

	return gears
}

/*
	Its gear ratio is the result of multiplying those two numbers together.
*/

func (g Gear) GetGearRatio() int {
	return g.Values[0] * g.Values[1]
}

func (es EngineSchematic) SumOfAllGearRatios() int {
	gears := es.GetGears()

	sum := 0

	for _, gear := range gears {
		sum += gear.GetGearRatio()
	}

	return sum
}

func isRuneADigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func parseInput(filename string) []string {
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

func main() {
	input := parseInput("input.txt")

	engineSchematic := ComputeEngineSchematic(input)
	part1Sum := engineSchematic.ComputeSumOfPartNumbers()
	fmt.Printf("Part 1: %d\n", part1Sum)

	part2Sum := engineSchematic.SumOfAllGearRatios()
	fmt.Printf("Part 2: %d\n", part2Sum)
}