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
			},
			expected: 6,
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.expected, v.input.StepsCountToZZZ())
	}
}
