package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	cards := ConvertInputToListOfCards(input)
	part1Score := cards.ComputeTotalPoints()
	fmt.Printf("Part 1: %d\n", part1Score)

	part2Score := cards.ComputeTotalCardsCount()
	fmt.Printf("Part 2: %d\n", part2Score)
}

/*
The Elf leads you over to the pile of colorful cards. There, you discover
dozens of scratchcards, all with their opaque covering already scratched off.
Picking one up, it looks like each card has two lists of numbers
separated by a vertical bar (|): a list of winning numbers and then
a list of numbers you have.
You organize the information into a table (your puzzle input).

For example:

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
*/
type CardNumber int

type Card struct {
	ID             CardNumber
	WinningNumbers []CardNumber
	MyNumbers      []CardNumber
}

type ElfStack []Card

func ConvertInputToListOfCards(input []string) ElfStack {
	cards := make([]Card, 0)

	for i, line := range input {
		card := Card{}
		card.ID = CardNumber(i + 1)

		// Remove the "Card X: " part
		numbers := strings.Split(line, ":")[1]

		// Split by "|" to get winning numbers and my numbers
		rawWinningNumbers := strings.Split(numbers, "|")[0]
		rawMyNumbers := strings.Split(numbers, "|")[1]

		// Split by " " to get each number
		for _, rawWinningNumber := range strings.Split(rawWinningNumbers, " ") {
			// Space at beginning of line, or double space before single digit
			if rawWinningNumber == " " || rawWinningNumber == "" {
				continue
			}
			rawWinningNumber = strings.TrimSpace(rawWinningNumber)
			number, err := strconv.Atoi(rawWinningNumber)
			if err != nil {
				panic(err)
			}
			card.WinningNumbers = append(card.WinningNumbers, CardNumber(number))
		}

		for _, rawMyNumber := range strings.Split(rawMyNumbers, " ") {
			// Space at beginning of line, or double space before single digit
			if rawMyNumber == " " || rawMyNumber == "" {
				continue
			}
			rawMyNumber = strings.TrimSpace(rawMyNumber)
			number, err := strconv.Atoi(rawMyNumber)
			if err != nil {
				panic(err)
			}
			card.MyNumbers = append(card.MyNumbers, CardNumber(number))
		}

		cards = append(cards, card)
	}

	return cards
}

/*
As far as the Elf has been able to figure out, you have to figure out which
of the numbers you have appear in the list of winning numbers. The first match
makes the card worth one point and each match after the first doubles
the point value of that card.
*/

func (card Card) ComputePoints() int {
	score := 0

	for _, myNumber := range card.MyNumbers {
		for _, winningNumber := range card.WinningNumbers {
			if myNumber == winningNumber {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}

	return score
}

func (elfStack ElfStack) ComputeTotalPoints() int {
	totalPoints := 0

	for _, card := range elfStack {
		totalPoints += card.ComputePoints()
	}

	return totalPoints
}

/*
There's no such thing as "points". Instead, scratchcards only cause you to win
more scratchcards equal to the number of winning numbers you have.
*/

func (stack ElfStack) ComputeTotalCardsCount() int {
	cardMatchCount := make([]int, len(stack))
	cardCount := make([]int, len(stack))

	for _, card := range stack {
		cardMatchCount[card.ID-1] = card.ComputeMatchCount()
		cardCount[card.ID-1] = 1
	}

	totalCardsCount := 0

	for card, count := range cardCount {
		totalCardsCount += count

		for i := 0; i < cardMatchCount[card]; i++ {
			cardNumberToIncrease := card + 1 + i

			if int(cardNumberToIncrease) >= len(cardMatchCount) {
				break
			}

			cardCount[cardNumberToIncrease] += count
		}
	}

	return totalCardsCount
}

func (card Card) ComputeMatchCount() int {
	count := 0

	for _, myNumber := range card.MyNumbers {
		for _, winningNumber := range card.WinningNumbers {
			if myNumber == winningNumber {
				count++
			}
		}
	}

	return count
}
