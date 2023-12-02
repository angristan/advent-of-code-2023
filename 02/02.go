package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
	For example, the record of a few games might look like this:

	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

type CubeColor string
type CubeSample map[CubeColor]int
type GameSet []CubeSample
type GameSetsInput []GameSet

func ConvertInput(input []string) GameSetsInput {
	gameSets := GameSetsInput{}

	for _, line := range input {
		// drop "Game X: " prefixs
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]

		// split by ";" to get samples
		gameSet := GameSet{}
		for _, game := range strings.Split(line, ";") {
			gameMap := CubeSample{}
			for _, colorAndCount := range strings.Split(game, ",") {
				colorAndCount = strings.TrimSpace(colorAndCount)

				// split by " " to get count and color
				spaceIndex := strings.Index(colorAndCount, " ")

				countAsString := colorAndCount[:spaceIndex]
				color := colorAndCount[spaceIndex+1:]

				count, err := strconv.Atoi(countAsString)
				if err != nil {
					panic(err)
				}
				gameMap[CubeColor(color)] = count
			}
			gameSet = append(gameSet, gameMap)
		}

		gameSets = append(gameSets, gameSet)
	}

	return gameSets
}

/*
	Determine which games would have been possible if the bag had been
	loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes.
	What is the sum of the IDs of those games?
*/

var (
	targetCubeAvailability = map[CubeColor]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func (gameSets GameSetsInput) ComputeIDSumOfPossibleGames() int {
	sum := 0

	for id, gameSet := range gameSets {
		isPossible := true
		for _, sampleGame := range gameSet {
			for color, count := range sampleGame {
				if count > targetCubeAvailability[color] {
					isPossible = false
					break
				}
			}

		}

		if isPossible {
			sum += id + 1
		}
	}

	return sum
}

func main() {
	input := parseInput("input.txt")

	gameSets := ConvertInput(input)
	sum := gameSets.ComputeIDSumOfPossibleGames()

	println(sum)
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
