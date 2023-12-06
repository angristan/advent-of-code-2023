package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	parsedInput := ConvertRawInputToInput(input)
	part1Score := parsedInput.ComputeAllPossibleRecordCount()
	fmt.Printf("Part 1: %d\n", part1Score)
}

type Race struct {
	timeDurationMs   int
	distanceRecordMm int
}

type Input struct {
	Races []Race
}

var numberRegex = regexp.MustCompile(`\d+`)

func ConvertRawInputToInput(rawInput []string) Input {
	rawDurations := numberRegex.FindAllString(rawInput[0], -1)

	durations := []int{}
	for _, rawDuration := range rawDurations {
		duration, err := strconv.Atoi(rawDuration)
		if err != nil {
			panic(err)
		}

		durations = append(durations, duration)
	}

	rawDistances := numberRegex.FindAllString(rawInput[1], -1)

	input := Input{}
	for i, rawDistance := range rawDistances {
		distance, err := strconv.Atoi(rawDistance)
		if err != nil {
			panic(err)
		}

		race := Race{
			timeDurationMs:   durations[i],
			distanceRecordMm: distance,
		}

		input.Races = append(input.Races, race)
	}

	return input
}

func (race Race) ComputePossibleRecordsCount() int {
	possiblesRecordsCount := 0

	buttonPressedDuration := 0
	boatSpeed := 0

	for i := 0; i < race.timeDurationMs; i++ {
		buttonPressedDuration = i
		boatSpeed = i

		raceDuration := race.timeDurationMs - buttonPressedDuration

		distance := boatSpeed * raceDuration

		if distance > race.distanceRecordMm {
			possiblesRecordsCount += 1
		}
	}

	return possiblesRecordsCount
}

func (input Input) ComputeAllPossibleRecordCount() int {
	total := 0
	for _, race := range input.Races {
		if total == 0 {
			total = 1
		}
		total *= race.ComputePossibleRecordsCount()
	}

	return total
}
