package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRawInputToInput(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	expected := Input{
		Hands: []Hand{
			{Cards: "32T3K", Bid: 765},
			{Cards: "T55J5", Bid: 684},
			{Cards: "KK677", Bid: 28},
			{Cards: "KTJJT", Bid: 220},
			{Cards: "QQQJA", Bid: 483},
		},
	}

	assert.Equal(t, expected, ConvertRawInputToInput(input))
}

func TestOccurences(t *testing.T) {
	input := Input{
		Hands: []Hand{
			{Cards: "T55J5", Bid: 684},
			{Cards: "32T3K", Bid: 765},
			{Cards: "QQQJA", Bid: 483},
			{Cards: "KK677", Bid: 28},
			{Cards: "KTJJT", Bid: 220},
			{Cards: "T3T3J", Bid: 17},
			{Cards: "Q2KJJ", Bid: 13},
		},
	}

	expected := [][]int{
		{2, 0, 1, 0, 0},
		{3, 1, 0, 0, 0},
		{2, 0, 1, 0, 0},
		{1, 2, 0, 0, 0},
		{1, 2, 0, 0, 0},
		{1, 2, 0, 0, 0},
		{3, 1, 0, 0, 0},
	}

	for i, hand := range input.Hands {
		assert.Equal(t, expected[i], hand.ComputeOcurrences())
	}
}

func TestComputeAndAssignHandType(t *testing.T) {
	type test struct {
		input    []Hand
		expected []Hand
	}

	tests := []test{
		{
			input: []Hand{
				{Cards: "T55J5", Bid: 684},
				{Cards: "32T3K", Bid: 765},
				{Cards: "QQQJA", Bid: 483},
				{Cards: "KK677", Bid: 28},
				{Cards: "KTJJT", Bid: 220},
			},
			expected: []Hand{
				{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
			},
		},
		{
			input: []Hand{
				{Cards: "T3T3J", Bid: 17},
				{Cards: "Q2KJJ", Bid: 13},
			},
			expected: []Hand{
				{Cards: "T3T3J", Bid: 17, HandType: TwoPair},
				{Cards: "Q2KJJ", Bid: 13, HandType: OnePair},
			},
		},
	}

	for _, tc := range tests {
		for i := range tc.input {
			tc.input[i].ComputeAndAssignHandType()
			assert.Equal(t, tc.expected[i], tc.input[i])
		}
	}
}

func TestJokerMode(t *testing.T) {
	type test struct {
		input    []Hand
		expected []Hand
	}

	tests := []test{
		{
			input: []Hand{
				{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
			},
			expected: []Hand{
				{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
			},
		},
		{
			input: []Hand{
				{Cards: "T3T3J", Bid: 17, HandType: TwoPair},
				{Cards: "Q2KJJ", Bid: 13, HandType: OnePair},
				{Cards: "JJJQA", Bid: 483, HandType: ThreeOfAKind},
			},
			expected: []Hand{
				{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
				{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
				{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
			},
		},
	}

	for _, tc := range tests {
		for i := range tc.input {
			tc.input[i].JokerMode()
			assert.Equal(t, tc.expected[i], tc.input[i])
		}
	}
}

func TestSortHandsPart1(t *testing.T) {
	type test struct {
		input    Input
		expected Input
	}

	tests := []test{
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
				},
			},
			expected: Input{
				Hands: []Hand{
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
					{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				},
			},
		},
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
			expected: Input{
				Hands: []Hand{
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
		},
	}

	for _, tc := range tests {
		tc.input.SortHands(StrengthsPart1)
		assert.Equal(t, tc.expected, tc.input)
	}
}

func TestSortHandsPart2(t *testing.T) {
	type test struct {
		input    Input
		expected Input
	}

	tests := []test{
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
				},
			},
			expected: Input{
				Hands: []Hand{
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
					{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
					{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
				},
			},
		},
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
			expected: Input{
				Hands: []Hand{
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
		},
	}

	for _, tc := range tests {
		tc.input.SortHands(StrengthsPart2)
		assert.Equal(t, tc.expected, tc.input)
	}
}

func TestComputeTotalPoints(t *testing.T) {
	type test struct {
		input         Input
		expectedScore int
	}

	tests := []test{
		{
			input: Input{
				Hands: []Hand{
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
					{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				},
			},
			expectedScore: 6440,
		},
		{
			input: Input{
				Hands: []Hand{
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
					{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
					{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
				},
			},
			expectedScore: 5905,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedScore, tc.input.ComputeTotalPoints())
	}
}
