package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeEngineSchematic(t *testing.T) {
	type test struct {
		input []string
		want  EngineSchematic
	}

	tests := []test{
		{
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			want: EngineSchematic{
				Numbers: []Number{
					{Value: "467", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}},
					{Value: "114", DigitsCoordinates: []Coordinates{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}}},
					{Value: "35", DigitsCoordinates: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}},
					{Value: "633", DigitsCoordinates: []Coordinates{{X: 6, Y: 2}, {X: 7, Y: 2}, {X: 8, Y: 2}}},
					{Value: "617", DigitsCoordinates: []Coordinates{{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}}},
					{Value: "58", DigitsCoordinates: []Coordinates{{X: 7, Y: 5}, {X: 8, Y: 5}}},
					{Value: "592", DigitsCoordinates: []Coordinates{{X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}}},
					{Value: "755", DigitsCoordinates: []Coordinates{{X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7}}},
					{Value: "664", DigitsCoordinates: []Coordinates{{X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}}},
					{Value: "598", DigitsCoordinates: []Coordinates{{X: 5, Y: 9}, {X: 6, Y: 9}, {X: 7, Y: 9}}},
				},
				Symbols: []Symbol{
					{Coordinates: Coordinates{X: 3, Y: 1}, Value: "*"},
					{Coordinates: Coordinates{X: 6, Y: 3}, Value: "#"},
					{Coordinates: Coordinates{X: 3, Y: 4}, Value: "*"},
					{Coordinates: Coordinates{X: 5, Y: 5}, Value: "+"},
					{Coordinates: Coordinates{X: 3, Y: 8}, Value: "$"},
					{Coordinates: Coordinates{X: 5, Y: 8}, Value: "*"},
				},
			},
		},

		{
			// https://www.reddit.com/r/adventofcode/comments/189q9wv/2023_day_3_another_sample_grid_to_use/
			// thanks bro 🤝
			input: []string{
				"12.......*..",
				"+.........34",
				".......-12..",
				"..78........",
				"..*....60...",
				"78..........",
				".......23...",
				"....90*12...",
				"............",
				"2.2......12.",
				".*.........*",
				"1.1.......56",
			},
			want: EngineSchematic{
				Numbers: []Number{
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}}},
					{Value: "34", DigitsCoordinates: []Coordinates{{X: 10, Y: 1}, {X: 11, Y: 1}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 8, Y: 2}, {X: 9, Y: 2}}},
					{Value: "78", DigitsCoordinates: []Coordinates{{X: 2, Y: 3}, {X: 3, Y: 3}}},
					{Value: "60", DigitsCoordinates: []Coordinates{{X: 7, Y: 4}, {X: 8, Y: 4}}},
					{Value: "78", DigitsCoordinates: []Coordinates{{X: 0, Y: 5}, {X: 1, Y: 5}}},
					{Value: "23", DigitsCoordinates: []Coordinates{{X: 7, Y: 6}, {X: 8, Y: 6}}},
					{Value: "90", DigitsCoordinates: []Coordinates{{X: 4, Y: 7}, {X: 5, Y: 7}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 7, Y: 7}, {X: 8, Y: 7}}},
					{Value: "2", DigitsCoordinates: []Coordinates{{X: 0, Y: 9}}},
					{Value: "2", DigitsCoordinates: []Coordinates{{X: 2, Y: 9}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 9, Y: 9}, {X: 10, Y: 9}}},
					{Value: "1", DigitsCoordinates: []Coordinates{{X: 0, Y: 11}}},
					{Value: "1", DigitsCoordinates: []Coordinates{{X: 2, Y: 11}}},
					{Value: "56", DigitsCoordinates: []Coordinates{{X: 10, Y: 11}, {X: 11, Y: 11}}},
				},
				Symbols: []Symbol{
					{Coordinates: Coordinates{X: 9, Y: 0}, Value: "*"},
					{Coordinates: Coordinates{X: 0, Y: 1}, Value: "+"},
					{Coordinates: Coordinates{X: 7, Y: 2}, Value: "-"},
					{Coordinates: Coordinates{X: 2, Y: 4}, Value: "*"},
					{Coordinates: Coordinates{X: 6, Y: 7}, Value: "*"},
					{Coordinates: Coordinates{X: 1, Y: 10}, Value: "*"},
					{Coordinates: Coordinates{X: 11, Y: 10}, Value: "*"},
				},
			},
		},
	}

	for _, tc := range tests {
		got := ConvertInputToEngineSchematic(tc.input)

		assert.Equal(t, tc.want, got)
	}
}

func TestGetPartNumbers(t *testing.T) {
	type test struct {
		input EngineSchematic
		want  []int
	}

	tests := []test{
		{
			input: EngineSchematic{
				Numbers: []Number{
					{Value: "467", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}},
					{Value: "114", DigitsCoordinates: []Coordinates{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}}},
					{Value: "35", DigitsCoordinates: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}},
					{Value: "633", DigitsCoordinates: []Coordinates{{X: 6, Y: 2}, {X: 7, Y: 2}, {X: 8, Y: 2}}},
					{Value: "617", DigitsCoordinates: []Coordinates{{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}}},
					{Value: "58", DigitsCoordinates: []Coordinates{{X: 7, Y: 5}, {X: 8, Y: 5}}},
					{Value: "592", DigitsCoordinates: []Coordinates{{X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}}},
					{Value: "755", DigitsCoordinates: []Coordinates{{X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7}}},
					{Value: "664", DigitsCoordinates: []Coordinates{{X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}}},
					{Value: "598", DigitsCoordinates: []Coordinates{{X: 5, Y: 9}, {X: 6, Y: 9}, {X: 7, Y: 9}}},
				},
				Symbols: []Symbol{
					{Coordinates: Coordinates{X: 3, Y: 1}, Value: "*"},
					{Coordinates: Coordinates{X: 6, Y: 3}, Value: "#"},
					{Coordinates: Coordinates{X: 3, Y: 4}, Value: "*"},
					{Coordinates: Coordinates{X: 5, Y: 5}, Value: "+"},
					{Coordinates: Coordinates{X: 3, Y: 8}, Value: "$"},
					{Coordinates: Coordinates{X: 5, Y: 8}, Value: "*"},
				},
			},
			want: []int{467, 35, 633, 617, 592, 755, 664, 598},
		},
		{
			input: EngineSchematic{
				Numbers: []Number{
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}}},
					{Value: "34", DigitsCoordinates: []Coordinates{{X: 10, Y: 1}, {X: 11, Y: 1}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 8, Y: 2}, {X: 9, Y: 2}}},
					{Value: "78", DigitsCoordinates: []Coordinates{{X: 2, Y: 3}, {X: 3, Y: 3}}},
					{Value: "60", DigitsCoordinates: []Coordinates{{X: 7, Y: 4}, {X: 8, Y: 4}}},
					{Value: "78", DigitsCoordinates: []Coordinates{{X: 0, Y: 5}, {X: 1, Y: 5}}},
					{Value: "23", DigitsCoordinates: []Coordinates{{X: 7, Y: 6}, {X: 8, Y: 6}}},
					{Value: "90", DigitsCoordinates: []Coordinates{{X: 4, Y: 7}, {X: 5, Y: 7}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 7, Y: 7}, {X: 8, Y: 7}}},
					{Value: "2", DigitsCoordinates: []Coordinates{{X: 0, Y: 9}}},
					{Value: "2", DigitsCoordinates: []Coordinates{{X: 2, Y: 9}}},
					{Value: "12", DigitsCoordinates: []Coordinates{{X: 9, Y: 9}, {X: 10, Y: 9}}},
					{Value: "1", DigitsCoordinates: []Coordinates{{X: 0, Y: 11}}},
					{Value: "1", DigitsCoordinates: []Coordinates{{X: 2, Y: 11}}},
					{Value: "56", DigitsCoordinates: []Coordinates{{X: 10, Y: 11}, {X: 11, Y: 11}}},
				},
				Symbols: []Symbol{
					{Coordinates: Coordinates{X: 9, Y: 0}, Value: "*"},
					{Coordinates: Coordinates{X: 0, Y: 1}, Value: "+"},
					{Coordinates: Coordinates{X: 7, Y: 2}, Value: "-"},
					{Coordinates: Coordinates{X: 2, Y: 4}, Value: "*"},
					{Coordinates: Coordinates{X: 6, Y: 7}, Value: "*"},
					{Coordinates: Coordinates{X: 1, Y: 10}, Value: "*"},
					{Coordinates: Coordinates{X: 11, Y: 10}, Value: "*"},
				},
			},
			want: []int{12, 34, 12, 78, 78, 23, 90, 12, 2, 2, 12, 1, 1, 56},
		},
	}

	for _, tc := range tests {
		got := tc.input.GetPartNumbersValues()

		assert.Equal(t, tc.want, got)
	}
}

func TestComputeSumOfPartNumbers(t *testing.T) {
	engineSchematic := EngineSchematic{
		Numbers: []Number{
			{Value: "467", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}},
			{Value: "114", DigitsCoordinates: []Coordinates{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}}},
			{Value: "35", DigitsCoordinates: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}},
			{Value: "633", DigitsCoordinates: []Coordinates{{X: 6, Y: 2}, {X: 7, Y: 2}, {X: 8, Y: 2}}},
			{Value: "617", DigitsCoordinates: []Coordinates{{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}}},
			{Value: "58", DigitsCoordinates: []Coordinates{{X: 7, Y: 5}, {X: 8, Y: 5}}},
			{Value: "592", DigitsCoordinates: []Coordinates{{X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}}},
			{Value: "755", DigitsCoordinates: []Coordinates{{X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7}}},
			{Value: "664", DigitsCoordinates: []Coordinates{{X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}}},
			{Value: "598", DigitsCoordinates: []Coordinates{{X: 5, Y: 9}, {X: 6, Y: 9}, {X: 7, Y: 9}}},
		},
		Symbols: []Symbol{
			{Coordinates: Coordinates{X: 3, Y: 1}, Value: "*"},
			{Coordinates: Coordinates{X: 6, Y: 3}, Value: "#"},
			{Coordinates: Coordinates{X: 3, Y: 4}, Value: "*"},
			{Coordinates: Coordinates{X: 5, Y: 5}, Value: "+"},
			{Coordinates: Coordinates{X: 3, Y: 8}, Value: "$"},
			{Coordinates: Coordinates{X: 5, Y: 8}, Value: "*"},
		},
	}

	expectedSum := 4361
	assert.Equal(t, expectedSum, engineSchematic.ComputeSumOfPartNumbers())
}

func TestGetGears(t *testing.T) {
	engineSchematic := EngineSchematic{
		Numbers: []Number{
			{Value: "467", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}},
			{Value: "114", DigitsCoordinates: []Coordinates{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}}},
			{Value: "35", DigitsCoordinates: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}},
			{Value: "633", DigitsCoordinates: []Coordinates{{X: 6, Y: 2}, {X: 7, Y: 2}, {X: 8, Y: 2}}},
			{Value: "617", DigitsCoordinates: []Coordinates{{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}}},
			{Value: "58", DigitsCoordinates: []Coordinates{{X: 7, Y: 5}, {X: 8, Y: 5}}},
			{Value: "592", DigitsCoordinates: []Coordinates{{X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}}},
			{Value: "755", DigitsCoordinates: []Coordinates{{X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7}}},
			{Value: "664", DigitsCoordinates: []Coordinates{{X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}}},
			{Value: "598", DigitsCoordinates: []Coordinates{{X: 5, Y: 9}, {X: 6, Y: 9}, {X: 7, Y: 9}}},
		},
		Symbols: []Symbol{
			{Coordinates: Coordinates{X: 3, Y: 1}, Value: "*"},
			{Coordinates: Coordinates{X: 6, Y: 3}, Value: "#"},
			{Coordinates: Coordinates{X: 3, Y: 4}, Value: "*"},
			{Coordinates: Coordinates{X: 5, Y: 5}, Value: "+"},
			{Coordinates: Coordinates{X: 3, Y: 8}, Value: "$"},
			{Coordinates: Coordinates{X: 5, Y: 8}, Value: "*"},
		},
	}

	expectedGears := []Gear{
		{Values: []int{467, 35}},
		{Values: []int{755, 598}},
	}
	assert.Equal(t, expectedGears, engineSchematic.GetGears())
}

func TestSumOfAllGearRatios(t *testing.T) {
	engineSchematic := EngineSchematic{
		Numbers: []Number{
			{Value: "467", DigitsCoordinates: []Coordinates{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}},
			{Value: "114", DigitsCoordinates: []Coordinates{{X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}}},
			{Value: "35", DigitsCoordinates: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}},
			{Value: "633", DigitsCoordinates: []Coordinates{{X: 6, Y: 2}, {X: 7, Y: 2}, {X: 8, Y: 2}}},
			{Value: "617", DigitsCoordinates: []Coordinates{{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}}},
			{Value: "58", DigitsCoordinates: []Coordinates{{X: 7, Y: 5}, {X: 8, Y: 5}}},
			{Value: "592", DigitsCoordinates: []Coordinates{{X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}}},
			{Value: "755", DigitsCoordinates: []Coordinates{{X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7}}},
			{Value: "664", DigitsCoordinates: []Coordinates{{X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}}},
			{Value: "598", DigitsCoordinates: []Coordinates{{X: 5, Y: 9}, {X: 6, Y: 9}, {X: 7, Y: 9}}},
		},
		Symbols: []Symbol{
			{Coordinates: Coordinates{X: 3, Y: 1}, Value: "*"},
			{Coordinates: Coordinates{X: 6, Y: 3}, Value: "#"},
			{Coordinates: Coordinates{X: 3, Y: 4}, Value: "*"},
			{Coordinates: Coordinates{X: 5, Y: 5}, Value: "+"},
			{Coordinates: Coordinates{X: 3, Y: 8}, Value: "$"},
			{Coordinates: Coordinates{X: 5, Y: 8}, Value: "*"},
		},
	}

	expectedGearsRatioSum := 467835

	assert.Equal(t, expectedGearsRatioSum, engineSchematic.SumOfAllGearRatios())
}
