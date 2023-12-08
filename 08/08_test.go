package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRawInputToMap(t *testing.T) {
	type test struct {
		input    []string
		expected Map
	}

	tests := []test{
		{
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},

			expected: Map{
				Directions: []Direction{Right, Left},
				Nodes: map[string]Node{
					"AAA": {Value: "AAA", Left: "BBB", Right: "CCC"},
					"BBB": {Value: "BBB", Left: "DDD", Right: "EEE"},
					"CCC": {Value: "CCC", Left: "ZZZ", Right: "GGG"},
					"DDD": {Value: "DDD", Left: "DDD", Right: "DDD"},
					"EEE": {Value: "EEE", Left: "EEE", Right: "EEE"},
					"GGG": {Value: "GGG", Left: "GGG", Right: "GGG"},
					"ZZZ": {Value: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
				},
				EndingANodesKeys: []string{"AAA"},
			},
		},
		{
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},

			expected: Map{
				Directions: []Direction{Left, Left, Right},
				Nodes: map[string]Node{
					"AAA": {Value: "AAA", Left: "BBB", Right: "BBB"},
					"BBB": {Value: "BBB", Left: "AAA", Right: "ZZZ"},
					"ZZZ": {Value: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
				},
				EndingANodesKeys: []string{"AAA"},
			},
		},
		{
			input: []string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},

			expected: Map{
				Directions: []Direction{Left, Right},
				Nodes: map[string]Node{
					"11A": {Value: "11A", Left: "11B", Right: "XXX"},
					"11B": {Value: "11B", Left: "XXX", Right: "11Z"},
					"11Z": {Value: "11Z", Left: "11B", Right: "XXX"},
					"22A": {Value: "22A", Left: "22B", Right: "XXX"},
					"22B": {Value: "22B", Left: "22C", Right: "22C"},
					"22C": {Value: "22C", Left: "22Z", Right: "22Z"},
					"22Z": {Value: "22Z", Left: "22B", Right: "22B"},
					"XXX": {Value: "XXX", Left: "XXX", Right: "XXX"},
				},
				EndingANodesKeys: []string{"11A", "22A"},
			},
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.expected, ConvertRawInputToMap(v.input))
	}
}

func TestStepsCountToZZZ(t *testing.T) {
	type test struct {
		input    Map
		expected int
	}

	tests := []test{
		{
			input: Map{
				Directions: []Direction{Right, Left},
				Nodes: map[string]Node{
					"AAA": {Value: "AAA", Left: "BBB", Right: "CCC"},
					"BBB": {Value: "BBB", Left: "DDD", Right: "EEE"},
					"CCC": {Value: "CCC", Left: "ZZZ", Right: "GGG"},
					"DDD": {Value: "DDD", Left: "DDD", Right: "DDD"},
					"EEE": {Value: "EEE", Left: "EEE", Right: "EEE"},
					"GGG": {Value: "GGG", Left: "GGG", Right: "GGG"},
					"ZZZ": {Value: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
				},
				EndingANodesKeys: []string{"AAA"},
			},
			expected: 2,
		},
		{
			input: Map{
				Directions: []Direction{Left, Left, Right},
				Nodes: map[string]Node{
					"AAA": {Value: "AAA", Left: "BBB", Right: "BBB"},
					"BBB": {Value: "BBB", Left: "AAA", Right: "ZZZ"},
					"ZZZ": {Value: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
				},
				EndingANodesKeys: []string{"AAA"},
			},
			expected: 6,
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.expected, v.input.StepsCountToZZZ())
	}
}

func TestStepsCountToEndingZGhostMode(t *testing.T) {
	type test struct {
		input    Map
		expected int
	}

	tests := []test{
		{
			input: Map{
				Directions: []Direction{Left, Right},
				Nodes: map[string]Node{
					"11A": {Value: "11A", Left: "11B", Right: "XXX"},
					"11B": {Value: "11B", Left: "XXX", Right: "11Z"},
					"11Z": {Value: "11Z", Left: "11B", Right: "XXX"},
					"22A": {Value: "22A", Left: "22B", Right: "XXX"},
					"22B": {Value: "22B", Left: "22C", Right: "22C"},
					"22C": {Value: "22C", Left: "22Z", Right: "22Z"},
					"22Z": {Value: "22Z", Left: "22B", Right: "22B"},
					"XXX": {Value: "XXX", Left: "XXX", Right: "XXX"},
				},
				EndingANodesKeys: []string{"11A", "22A"},
			},
			expected: 6,
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.expected, v.input.StepsCountToEndingZGhostMode())
	}
}
