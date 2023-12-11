package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRawInputToImage(t *testing.T) {
	rawInput := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	expectedImage := Image{
		{s, s, s, G, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, G, s, s},
		{G, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, G, s, s, s},
		{s, G, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, G},
		{s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, G, s, s},
		{G, s, s, s, G, s, s, s, s, s},
	}

	assert.Equal(t, expectedImage, ConvertRawInputToImage(rawInput))
}

func TestComputeShortestBetween(t *testing.T) {
	type testCase struct {
		start    Coords
		end      Coords
		expected int
	}

	testCases := []testCase{
		{Coords{0, 0}, Coords{0, 1}, 1},
		{Coords{0, 0}, Coords{1, 0}, 1},
		{Coords{0, 0}, Coords{1, 1}, 2},
		{Coords{4, 0}, Coords{9, 10}, 15},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, ComputeShortestBetween(testCase.start, testCase.end))
	}
}

func TestSumShortestPathBetweenAllGalaxies(t *testing.T) {
	image := Image{
		{s, s, s, s, G, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, G, s, s, s},
		{G, s, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, G, s, s, s, s},
		{s, G, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, s, s, s, G},
		{s, s, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, s, s, s, s},
		{s, s, s, s, s, s, s, s, s, G, s, s, s},
		{G, s, s, s, s, G, s, s, s, s, s, s, s},
	}

	assert.Equal(t, 374, image.SumShortestPathBetweenAllGalaxies(1))
}
