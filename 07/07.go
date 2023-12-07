package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	parsedInput := ConvertRawInputToInput(input)
	part1Score := parsedInput.ComputeTotalPoints()
	fmt.Printf("Part 1: %d\n", part1Score)

	parsedInputV2 := ConvertRawInputToInputV2(input)
	for i := range parsedInputV2.Hands {
		parsedInputV2.Hands[i].ComputeAndAssignHandType()
		parsedInputV2.Hands[i].JokerMode()
	}
	parsedInputV2.SortHands()
	part2Score := parsedInputV2.ComputeTotalPoints()
	fmt.Printf("Part 2: %d\n", part2Score)
}

type HandV1 struct {
	Cards string
	Bid   int
}

type InputV1 struct {
	Hands []HandV1
}

func ConvertRawInputToInput(rawInput []string) InputV1 {
	input := InputV1{}

	for _, line := range rawInput {
		splitLine := strings.Split(line, " ")
		cards := splitLine[0]
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}

		hand := HandV1{
			Cards: cards,
			Bid:   bid,
		}

		input.Hands = append(input.Hands, hand)
	}

	return input
}

func (input InputV1) SortHands() {
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

func (hand HandV1) Occurrences() []int {
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

func (input InputV1) ComputeTotalPoints() int {
	input.SortHands()

	totalPoints := 0
	for i, hand := range input.Hands {
		totalPoints += hand.Bid * (i + 1)
	}

	return totalPoints
}

/*
==============================
	Part 2
==============================
*/

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type HandV2 struct {
	Cards    string
	Bid      int
	HandType HandType
}

type InputV2 struct {
	Hands []HandV2
}

func ConvertRawInputToInputV2(rawInput []string) InputV2 {
	input := InputV2{}

	for _, line := range rawInput {
		splitLine := strings.Split(line, " ")
		cards := splitLine[0]
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}

		hand := HandV2{
			Cards: cards,
			Bid:   bid,
		}

		input.Hands = append(input.Hands, hand)
	}

	return input
}

func (hand HandV2) ComputeOcurrences() []int {
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

func (hand *HandV2) ComputeAndAssignHandType() {
	occurrences := hand.ComputeOcurrences()

	if occurrences[5-1] == 1 {
		hand.HandType = FiveOfAKind
		return
	}

	if occurrences[4-1] == 1 {
		hand.HandType = FourOfAKind
		return
	}

	if occurrences[3-1] == 1 && occurrences[2-1] == 1 {
		hand.HandType = FullHouse
		return
	}

	if occurrences[3-1] == 1 {
		hand.HandType = ThreeOfAKind
		return
	}

	if occurrences[2-1] == 2 {
		hand.HandType = TwoPair
		return
	}

	if occurrences[2-1] == 1 {
		hand.HandType = OnePair
		return
	}

	hand.HandType = HighCard
}

func (hand *HandV2) JokerMode() {
	JCount := strings.Count(hand.Cards, "J")

	if JCount == 0 {
		return
	}

	for i := 0; i < JCount; i++ {
		switch hand.HandType {
		case FiveOfAKind:
			return
		case FourOfAKind:
			hand.HandType = FiveOfAKind
			if strings.Count(hand.Cards, "J") == 4 {
				return
			}
		case FullHouse:
			hand.HandType = FourOfAKind
		case ThreeOfAKind:
			hand.HandType = FourOfAKind
			if strings.Count(hand.Cards, "J") == 3 {
				return
			}
		case TwoPair:
			hand.HandType = FullHouse
		case OnePair:
			hand.HandType = ThreeOfAKind
			if strings.Count(hand.Cards, "J") == 2 {
				return
			}
		case HighCard:
			hand.HandType = OnePair
		}
	}
}

func (input InputV2) SortHands() {
	sort.Slice(input.Hands, func(i, j int) bool {
		if input.Hands[i].HandType != input.Hands[j].HandType {
			return input.Hands[i].HandType < input.Hands[j].HandType
		}

		strengths := map[string]int{
			"A": 13,
			"K": 12,
			"Q": 11,
			"T": 10,
			"9": 9,
			"8": 8,
			"7": 7,
			"6": 6,
			"5": 5,
			"4": 4,
			"3": 3,
			"2": 2,
			"J": 1,
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

func (input InputV2) ComputeTotalPoints() int {
	input.SortHands()

	totalPoints := 0
	for i, hand := range input.Hands {
		totalPoints += hand.Bid * (i + 1)
	}

	return totalPoints
}
