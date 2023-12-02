package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInput(t *testing.T) {
	type test struct {
		input []string
		want  GameSetsInput
	}

	tests := []test{
		{
			[]string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			GameSetsInput{
				{
					{"blue": 3, "red": 4},
					{"red": 1, "green": 2, "blue": 6},
					{"green": 2},
				},
				{
					{"blue": 1, "green": 2},
					{"green": 3, "blue": 4, "red": 1},
					{"green": 1, "blue": 1},
				},
				{
					{"green": 8, "blue": 6, "red": 20},
					{"blue": 5, "red": 4, "green": 13},
					{"green": 5, "red": 1},
				},
				{
					{"green": 1, "red": 3, "blue": 6},
					{"green": 3, "red": 6},
					{"green": 3, "blue": 15, "red": 14},
				},
				{
					{"red": 6, "blue": 1, "green": 3},
					{"blue": 2, "red": 1, "green": 2},
				},
			},
		},
	}

	for _, tc := range tests {
		got := ConvertInput(tc.input)

		assert.Equal(t, tc.want, got)
	}
}

func TestComputeIDSumOfPossibleGames(t *testing.T) {
	type test struct {
		input GameSetsInput
		want  int
	}

	tests := []test{
		{
			GameSetsInput{
				{
					{"blue": 3, "red": 4},
					{"red": 1, "green": 2, "blue": 6},
					{"green": 2},
				},
				{
					{"blue": 1, "green": 2},
					{"green": 3, "blue": 4, "red": 1},
					{"green": 1, "blue": 1},
				},
				{
					{"green": 8, "blue": 6, "red": 20},
					{"blue": 5, "red": 4, "green": 13},
					{"green": 5, "red": 1},
				},
				{
					{"green": 1, "red": 3, "blue": 6},
					{"green": 3, "red": 6},
					{"green": 3, "blue": 15, "red": 14},
				},
				{
					{"red": 6, "blue": 1, "green": 3},
					{"blue": 2, "red": 1, "green": 2},
				},
			},
			8,
		},
	}

	for _, tc := range tests {
		got := tc.input.ComputeIDSumOfPossibleGames()

		assert.Equal(t, tc.want, got)
	}
}

func TestComputeMinimumGameSet(t *testing.T) {
	type test struct {
		input GameSet
		want  MinimumGameSet
	}

	tests := []test{
		{
			GameSet{
				{"blue": 3, "red": 4},
				{"red": 1, "green": 2, "blue": 6},
				{"green": 2},
			},
			MinimumGameSet{
				"red":   4,
				"green": 2,
				"blue":  6,
			},
		},
		{
			GameSet{
				{"blue": 1, "green": 2},
				{"green": 3, "blue": 4, "red": 1},
				{"green": 1, "blue": 1},
			},
			MinimumGameSet{
				"red":   1,
				"green": 3,
				"blue":  4,
			},
		},
		{
			GameSet{
				{"green": 8, "blue": 6, "red": 20},
				{"blue": 5, "red": 4, "green": 13},
				{"green": 5, "red": 1},
			},
			MinimumGameSet{
				"red":   20,
				"green": 13,
				"blue":  6,
			},
		},
		{
			GameSet{
				{"green": 1, "red": 3, "blue": 6},
				{"green": 3, "red": 6},
				{"green": 3, "blue": 15, "red": 14},
			},
			MinimumGameSet{
				"red":   14,
				"green": 3,
				"blue":  15,
			},
		},
		{
			GameSet{
				{"red": 6, "blue": 1, "green": 3},
				{"blue": 2, "red": 1, "green": 2},
			},
			MinimumGameSet{
				"red":   6,
				"green": 3,
				"blue":  2,
			},
		},
	}

	for _, tc := range tests {
		got := tc.input.ComputeMinimumGameSet()

		assert.Equal(t, tc.want, got)
	}
}

func TestComputeSumOfPowerOfMinimalGameSets(t *testing.T) {
	type test struct {
		input GameSetsInput
		want  int
	}

	tests := []test{
		{
			GameSetsInput{
				{
					{"blue": 3, "red": 4},
					{"red": 1, "green": 2, "blue": 6},
					{"green": 2},
				},
				{
					{"blue": 1, "green": 2},
					{"green": 3, "blue": 4, "red": 1},
					{"green": 1, "blue": 1},
				},
				{
					{"green": 8, "blue": 6, "red": 20},
					{"blue": 5, "red": 4, "green": 13},
					{"green": 5, "red": 1},
				},
				{
					{"green": 1, "red": 3, "blue": 6},
					{"green": 3, "red": 6},
					{"green": 3, "blue": 15, "red": 14},
				},
				{
					{"red": 6, "blue": 1, "green": 3},
					{"blue": 2, "red": 1, "green": 2},
				},
			},
			2286,
		},
	}

	for _, tc := range tests {
		got := tc.input.ComputeSumOfPowerOfMinimalGameSets()

		assert.Equal(t, tc.want, got)
	}
}
