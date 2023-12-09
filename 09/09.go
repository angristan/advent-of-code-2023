package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	report := ConvertRawInputToReport(input)
	fmt.Printf("Part 1: %d\n", report.ComputeSumOfNextValues())

	report = ConvertRawInputToReport(input)
	fmt.Printf("Part 2: %d\n", report.ComputeSumOfPreviousValues())
}

type History struct {
	Values []int
}

type Report struct {
	Histories []History
}

var numberRegex = regexp.MustCompile(`\-*\d+`)

func ConvertRawInputToReport(rawInput []string) Report {
	report := Report{}

	for _, rawHistory := range rawInput {
		rawValues := numberRegex.FindAllString(rawHistory, -1)

		values := []int{}
		for _, rawValue := range rawValues {
			value, err := strconv.Atoi(rawValue)
			if err != nil {
				panic(err)
			}

			values = append(values, value)
		}

		history := History{
			Values: values,
		}

		report.Histories = append(report.Histories, history)
	}

	return report
}

func (history History) ComputeNextValue() int {
	nextValue := 0

	currentSlice := history.Values
	for true {
		nextValue += currentSlice[len(currentSlice)-1]
		nextSlice := make([]int, len(currentSlice)-1)

		for k := range currentSlice {
			if k == 0 {
				continue
			}

			delta := currentSlice[k] - currentSlice[k-1]
			nextSlice[k-1] = delta

		}

		// next slice is all 0s
		if slices.Max(nextSlice) == 0 && slices.Min(nextSlice) == 0 {
			break
		}

		currentSlice = nextSlice
	}

	return nextValue
}

func (report Report) ComputeSumOfNextValues() int {
	sum := 0

	for _, history := range report.Histories {
		sum += history.ComputeNextValue()
	}

	return sum
}

func (history History) ComputePreviousValue() int {
	// reverse slice so that the predicated next value will actually be the previous value
	for i := len(history.Values)/2 - 1; i >= 0; i-- {
		opp := len(history.Values) - 1 - i
		history.Values[i], history.Values[opp] = history.Values[opp], history.Values[i]
	}

	return history.ComputeNextValue()
}

func (report Report) ComputeSumOfPreviousValues() int {
	sum := 0

	for _, history := range report.Histories {
		sum += history.ComputePreviousValue()
	}

	return sum
}
