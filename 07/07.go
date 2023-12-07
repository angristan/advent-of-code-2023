package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("/Users/stanislaslange/advent-of-code-2023/07/input.txt")

	parsedInput := ConvertRawInputToInput(input)
	part1Score := parsedInput.ComputeTotalPoints()
	fmt.Printf("Part 1: %d\n", part1Score)
}

type Hand struct {
	Cards string
	Bid   int
}

type Input struct {
	Hands []Hand
}

func ConvertRawInputToInput(rawInput []string) Input {
	input := Input{}

	for _, line := range rawInput {
		splitLine := strings.Split(line, " ")
		cards := splitLine[0]
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}

		hand := Hand{
			Cards: cards,
			Bid:   bid,
		}

		input.Hands = append(input.Hands, hand)
	}

	return input
}

func (input Input) SortHands() {
	sort.Slice(input.Hands, func(i, j int) bool {
		occurrencesI := input.Hands[i].Occurrences()
		occurrencesJ := input.Hands[j].Occurrences()

		maxOccurrenceCountI := 0
		for i, occurrence := range occurrencesI {
			if occurrence > 0 {
				maxOccurrenceCountI = i + 1
			}
		}

		maxOccurrenceCountJ := 0
		for i, occurrence := range occurrencesJ {
			if occurrence > 0 {
				maxOccurrenceCountJ = i + 1
			}
		}

		if maxOccurrenceCountI != maxOccurrenceCountJ {
			return maxOccurrenceCountI < maxOccurrenceCountJ
		}

		if maxOccurrenceCountI == 3 || maxOccurrenceCountI == 2 {
			if occurrencesI[2-1] != occurrencesJ[2-1] {
				return occurrencesI[2-1] < occurrencesJ[2-1]
			}
		}

		strengths := map[string]int{
			"A": 13,
			"K": 12,
			"Q": 11,
			"J": 10,
			"T": 9,
			"9": 8,
			"8": 7,
			"7": 6,
			"6": 5,
			"5": 4,
			"4": 3,
			"3": 2,
			"2": 1,
		}

		for cardIndex := 0; cardIndex < 5; cardIndex++ {
			cardJ := string(input.Hands[j].Cards[cardIndex])
			cardI := string(input.Hands[i].Cards[cardIndex])

			if strengths[cardI] != strengths[cardJ] {
				return strengths[cardI] < strengths[cardJ]
			}
		}

		return true
	})
}

func (hand Hand) Occurrences() []int {
	occurrences := make([]int, 5)

	for _, card := range hand.Cards {
		count := strings.Count(hand.Cards, string(card))

		if count == 0 {
			continue
		}
		indexCount := count - 1

		if ok := occurrences[indexCount]; ok == 0 {
			occurrences[indexCount] = 1
		} else {
			occurrences[indexCount]++
		}

		hand.Cards = strings.ReplaceAll(hand.Cards, string(card), "")
	}

	return occurrences
}

func (input Input) ComputeTotalPoints() int {
	input.SortHands()

	for _, hand := range input.Hands {
		fmt.Printf("%s %d\n", hand.Cards, hand.Bid)
	}

	totalPoints := 0
	for i, hand := range input.Hands {
		totalPoints += hand.Bid * (i + 1)
	}

	return totalPoints
}
