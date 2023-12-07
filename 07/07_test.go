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

	expected := InputV1{
		Hands: []HandV1{
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
	input := InputV1{
		Hands: []HandV1{
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
		assert.Equal(t, expected[i], hand.Occurrences())
	}
}

func TestSortHands(t *testing.T) {
	type test struct {
		input    InputV1
		expected InputV1
	}

	tests := []test{
		{
			input: InputV1{
				Hands: []HandV1{
					{Cards: "T55J5", Bid: 684},
					{Cards: "32T3K", Bid: 765},
					{Cards: "QQQJA", Bid: 483},
					{Cards: "KK677", Bid: 28},
					{Cards: "KTJJT", Bid: 220},
				},
			},
			expected: InputV1{
				Hands: []HandV1{
					{Cards: "32T3K", Bid: 765},
					{Cards: "KTJJT", Bid: 220},
					{Cards: "KK677", Bid: 28},
					{Cards: "T55J5", Bid: 684},
					{Cards: "QQQJA", Bid: 483},
				},
			},
		},
		{
			input: InputV1{
				Hands: []HandV1{
					{Cards: "T3T3J", Bid: 17},
					{Cards: "Q2KJJ", Bid: 13},
				},
			},
			expected: InputV1{
				Hands: []HandV1{
					{Cards: "Q2KJJ", Bid: 13},
					{Cards: "T3T3J", Bid: 17},
				},
			},
		},
	}

	for _, tc := range tests {
		tc.input.SortHands()
		assert.Equal(t, tc.expected, tc.input)
	}
}

func TestComputeTotalPoints(t *testing.T) {
	input := InputV1{
		Hands: []HandV1{
			{Cards: "32T3K", Bid: 765},
			{Cards: "KTJJT", Bid: 220},
			{Cards: "KK677", Bid: 28},
			{Cards: "T55J5", Bid: 684},
			{Cards: "QQQJA", Bid: 483},
		},
	}

	expected := 6440

	assert.Equal(t, expected, input.ComputeTotalPoints())
}

/*
==============================
	Part 2
==============================
*/

func TestConvertRawInputToInputV2(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	expected := InputV2{
		Hands: []HandV2{
			{Cards: "32T3K", Bid: 765},
			{Cards: "T55J5", Bid: 684},
			{Cards: "KK677", Bid: 28},
			{Cards: "KTJJT", Bid: 220},
			{Cards: "QQQJA", Bid: 483},
		},
	}

	assert.Equal(t, expected, ConvertRawInputToInputV2(input))
}

func TestOccurencesV2(t *testing.T) {
	input := InputV2{
		Hands: []HandV2{
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
		input    []HandV2
		expected []HandV2
	}

	tests := []test{
		{
			input: []HandV2{
				{Cards: "T55J5", Bid: 684},
				{Cards: "32T3K", Bid: 765},
				{Cards: "QQQJA", Bid: 483},
				{Cards: "KK677", Bid: 28},
				{Cards: "KTJJT", Bid: 220},
			},
			expected: []HandV2{
				{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
			},
		},
		{
			input: []HandV2{
				{Cards: "T3T3J", Bid: 17},
				{Cards: "Q2KJJ", Bid: 13},
			},
			expected: []HandV2{
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
		input    []HandV2
		expected []HandV2
	}

	tests := []test{
		{
			input: []HandV2{
				{Cards: "T55J5", Bid: 684, HandType: ThreeOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: ThreeOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: TwoPair},
			},
			expected: []HandV2{
				{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
				{Cards: "32T3K", Bid: 765, HandType: OnePair},
				{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
				{Cards: "KK677", Bid: 28, HandType: TwoPair},
				{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
			},
		},
		{
			input: []HandV2{
				{Cards: "T3T3J", Bid: 17, HandType: TwoPair},
				{Cards: "Q2KJJ", Bid: 13, HandType: OnePair},
				{Cards: "JJJQA", Bid: 483, HandType: ThreeOfAKind},
			},
			expected: []HandV2{
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

func TestSortHandsV2(t *testing.T) {
	type test struct {
		input    InputV2
		expected InputV2
	}

	tests := []test{
		{
			input: InputV2{
				Hands: []HandV2{
					{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
				},
			},
			expected: InputV2{
				Hands: []HandV2{
					{Cards: "32T3K", Bid: 765, HandType: OnePair},
					{Cards: "KK677", Bid: 28, HandType: TwoPair},
					{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
					{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
					{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
				},
			},
		},
		{
			input: InputV2{
				Hands: []HandV2{
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
			expected: InputV2{
				Hands: []HandV2{
					{Cards: "Q2KJJ", Bid: 13, HandType: ThreeOfAKind},
					{Cards: "T3T3J", Bid: 17, HandType: FullHouse},
					{Cards: "JJJQA", Bid: 483, HandType: FourOfAKind},
				},
			},
		},
	}

	for _, tc := range tests {
		tc.input.SortHands()
		assert.Equal(t, tc.expected, tc.input)
	}
}

func TestComputeTotalPointsV2(t *testing.T) {
	input := InputV2{
		Hands: []HandV2{
			{Cards: "T55J5", Bid: 684, HandType: FourOfAKind},
			{Cards: "32T3K", Bid: 765, HandType: OnePair},
			{Cards: "QQQJA", Bid: 483, HandType: FourOfAKind},
			{Cards: "KK677", Bid: 28, HandType: TwoPair},
			{Cards: "KTJJT", Bid: 220, HandType: FourOfAKind},
		},
	}

	expected := 5905

	assert.Equal(t, expected, input.ComputeTotalPoints())
}
