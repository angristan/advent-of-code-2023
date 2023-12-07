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
		assert.Equal(t, expected[i], hand.Occurrences())
	}
}

func TestSortHands(t *testing.T) {
	type test struct {
		input    Input
		expected Input
	}

	tests := []test{
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T55J5", Bid: 684},
					{Cards: "32T3K", Bid: 765},
					{Cards: "QQQJA", Bid: 483},
					{Cards: "KK677", Bid: 28},
					{Cards: "KTJJT", Bid: 220},
				},
			},
			expected: Input{
				Hands: []Hand{
					{Cards: "32T3K", Bid: 765},
					{Cards: "KTJJT", Bid: 220},
					{Cards: "KK677", Bid: 28},
					{Cards: "T55J5", Bid: 684},
					{Cards: "QQQJA", Bid: 483},
				},
			},
		},
		{
			input: Input{
				Hands: []Hand{
					{Cards: "T3T3J", Bid: 17},
					{Cards: "Q2KJJ", Bid: 13},
				},
			},
			expected: Input{
				Hands: []Hand{
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
	input := Input{
		Hands: []Hand{
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
