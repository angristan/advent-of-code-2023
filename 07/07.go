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
	for i := range parsedInput.Hands {
		parsedInput.Hands[i].ComputeAndAssignHandType()
	}
	parsedInput.SortHands(StrengthsPart1)
	part1Score := parsedInput.ComputeTotalPoints()
	fmt.Printf("Part 1: %d\n", part1Score)

	parsedInputPart2 := ConvertRawInputToInput(input)
	for i := range parsedInputPart2.Hands {
		parsedInputPart2.Hands[i].ComputeAndAssignHandType()
		parsedInputPart2.Hands[i].JokerMode()
	}
	parsedInputPart2.SortHands(StrengthsPart2)
	part2Score := parsedInputPart2.ComputeTotalPoints()
	fmt.Printf("Part 2: %d\n", part2Score)
}

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

type Hand struct {
	Cards    string
	Bid      int
	HandType HandType
}

type Input struct {
	Hands []Hand
}

var (
	StrengthsPart1 = map[string]int{
		"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1,
	}

	StrengthsPart2 = map[string]int{
		"A": 13, "K": 12, "Q": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1,
	}
)

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

func (hand Hand) ComputeOcurrences() []int {
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

func (hand *Hand) ComputeAndAssignHandType() {
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

func (hand *Hand) JokerMode() {
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

func (input Input) SortHands(strengthsMap map[string]int) {
	sort.Slice(input.Hands, func(i, j int) bool {
		if input.Hands[i].HandType != input.Hands[j].HandType {
			return input.Hands[i].HandType < input.Hands[j].HandType
		}

		// If both hands have the same type, compare the cards by strength
		for cardIndex := 0; cardIndex < 5; cardIndex++ {
			cardJ := string(input.Hands[j].Cards[cardIndex])
			cardI := string(input.Hands[i].Cards[cardIndex])

			if strengthsMap[cardI] != strengthsMap[cardJ] {
				return strengthsMap[cardI] < strengthsMap[cardJ]
			}
		}

		return true
	})
}

func (input Input) ComputeTotalPoints() int {
	totalPoints := 0
	for i, hand := range input.Hands {
		totalPoints += hand.Bid * (i + 1)
	}

	return totalPoints
}
