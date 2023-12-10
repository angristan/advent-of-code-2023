package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Grid1 = TheGrid{
		{
			{Type: PipeHorizontal, CoordX: 0, CoordY: 0},
			{Type: PipeBendL, CoordX: 1, CoordY: 0},
			{Type: PipeVertical, CoordX: 2, CoordY: 0},
			{Type: PipeBendF, CoordX: 3, CoordY: 0},
			{Type: PipeBend7, CoordX: 4, CoordY: 0},
		},
		{
			{Type: PipeBend7, CoordX: 0, CoordY: 1},
			{Type: Start, CoordX: 1, CoordY: 1},
			{Type: PipeHorizontal, CoordX: 2, CoordY: 1},
			{Type: PipeBend7, CoordX: 3, CoordY: 1},
			{Type: PipeVertical, CoordX: 4, CoordY: 1},
		},
		{
			{Type: PipeBendL, CoordX: 0, CoordY: 2},
			{Type: PipeVertical, CoordX: 1, CoordY: 2},
			{Type: PipeBend7, CoordX: 2, CoordY: 2},
			{Type: PipeVertical, CoordX: 3, CoordY: 2},
			{Type: PipeVertical, CoordX: 4, CoordY: 2},
		},
		{
			{Type: PipeHorizontal, CoordX: 0, CoordY: 3},
			{Type: PipeBendL, CoordX: 1, CoordY: 3},
			{Type: PipeHorizontal, CoordX: 2, CoordY: 3},
			{Type: PipeBendJ, CoordX: 3, CoordY: 3},
			{Type: PipeVertical, CoordX: 4, CoordY: 3},
		},
		{
			{Type: PipeBendL, CoordX: 0, CoordY: 4},
			{Type: PipeVertical, CoordX: 1, CoordY: 4},
			{Type: PipeHorizontal, CoordX: 2, CoordY: 4},
			{Type: PipeBendJ, CoordX: 3, CoordY: 4},
			{Type: PipeBendF, CoordX: 4, CoordY: 4},
		},
	}

	Grid2 = TheGrid{
		{
			{Type: PipeBend7, CoordX: 0, CoordY: 0},
			{Type: PipeHorizontal, CoordX: 1, CoordY: 0},
			{Type: PipeBendF, CoordX: 2, CoordY: 0},
			{Type: PipeBend7, CoordX: 3, CoordY: 0},
			{Type: PipeHorizontal, CoordX: 4, CoordY: 0},
		},
		{
			{Type: PipeGround, CoordX: 0, CoordY: 1},
			{Type: PipeBendF, CoordX: 1, CoordY: 1},
			{Type: PipeBendJ, CoordX: 2, CoordY: 1},
			{Type: PipeVertical, CoordX: 3, CoordY: 1},
			{Type: PipeBend7, CoordX: 4, CoordY: 1},
		},
		{
			{Type: Start, CoordX: 0, CoordY: 2},
			{Type: PipeBendJ, CoordX: 1, CoordY: 2},
			{Type: PipeBendL, CoordX: 2, CoordY: 2},
			{Type: PipeBendL, CoordX: 3, CoordY: 2},
			{Type: PipeBend7, CoordX: 4, CoordY: 2},
		},
		{
			{Type: PipeVertical, CoordX: 0, CoordY: 3},
			{Type: PipeBendF, CoordX: 1, CoordY: 3},
			{Type: PipeHorizontal, CoordX: 2, CoordY: 3},
			{Type: PipeHorizontal, CoordX: 3, CoordY: 3},
			{Type: PipeBendJ, CoordX: 4, CoordY: 3},
		},
		{
			{Type: PipeBendL, CoordX: 0, CoordY: 4},
			{Type: PipeBendJ, CoordX: 1, CoordY: 4},
			{Type: PipeGround, CoordX: 2, CoordY: 4},
			{Type: PipeBendL, CoordX: 3, CoordY: 4},
			{Type: PipeBendJ, CoordX: 4, CoordY: 4},
		},
	}
)

func TestConvertRawInputToSurfacePipes(t *testing.T) {
	type testCase struct {
		rawInput []string
		expected TheGrid
	}

	testCases := []testCase{
		{
			rawInput: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			expected: Grid1,
		},
		{
			rawInput: []string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			expected: Grid2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, ConvertRawInputToSurfacePipes(testCase.rawInput))
	}
}

func TestGetStartPipe(t *testing.T) {
	type testCase struct {
		grid     TheGrid
		expected Tile
	}

	testCases := []testCase{
		{
			grid:     Grid1,
			expected: Tile{Type: Start, CoordX: 1, CoordY: 1},
		},
		{
			grid:     Grid2,
			expected: Tile{Type: Start, CoordX: 0, CoordY: 2},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.grid.GetStartPipe())
	}
}

func TestGetAdjacentTiles(t *testing.T) {
	type testCase struct {
		grid     TheGrid
		tile     Tile
		expected []Tile
	}

	testCases := []testCase{
		{
			grid: Grid1,
			tile: Tile{Type: Start, CoordX: 1, CoordY: 1},
			expected: []Tile{
				{Type: PipeBend7, CoordX: 0, CoordY: 1},
				{Type: PipeHorizontal, CoordX: 2, CoordY: 1},
				{Type: PipeBendL, CoordX: 1, CoordY: 0},
				{Type: PipeVertical, CoordX: 1, CoordY: 2},
			},
		},
		{
			grid: Grid2,
			tile: Tile{Type: Start, CoordX: 0, CoordY: 2},
			expected: []Tile{
				{Type: PipeBendJ, CoordX: 1, CoordY: 2},
				{Type: PipeGround, CoordX: 0, CoordY: 1},
				{Type: PipeVertical, CoordX: 0, CoordY: 3},
			},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.grid.GetAdjacentTiles(testCase.tile))
	}
}

func TestGetAdjacentPipes(t *testing.T) {
	type testCase struct {
		grid     TheGrid
		tile     Tile
		expected []Tile
	}

	testCases := []testCase{
		{
			grid: Grid1,
			tile: Tile{Type: Start, CoordX: 1, CoordY: 1},
			expected: []Tile{
				{Type: PipeBend7, CoordX: 0, CoordY: 1},
				{Type: PipeHorizontal, CoordX: 2, CoordY: 1},
				{Type: PipeBendL, CoordX: 1, CoordY: 0},
				{Type: PipeVertical, CoordX: 1, CoordY: 2},
			},
		},
		{
			grid: Grid2,
			tile: Tile{Type: Start, CoordX: 0, CoordY: 2},
			expected: []Tile{
				{Type: PipeBendJ, CoordX: 1, CoordY: 2},
				{Type: PipeVertical, CoordX: 0, CoordY: 3},
			},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.grid.GetAdjacentPipes(testCase.tile))
	}
}

func TestGetConnectedPipes(t *testing.T) {
	type testCase struct {
		grid     TheGrid
		tile     Tile
		expected []Tile
	}

	testCases := []testCase{
		{
			grid: Grid1,
			tile: Tile{Type: Start, CoordX: 1, CoordY: 1},
			expected: []Tile{
				{Type: PipeHorizontal, CoordX: 2, CoordY: 1},
				{Type: PipeVertical, CoordX: 1, CoordY: 2},
			},
		},
		{
			grid: Grid2,
			tile: Tile{Type: Start, CoordX: 0, CoordY: 2},
			expected: []Tile{
				{Type: PipeBendJ, CoordX: 1, CoordY: 2},
				{Type: PipeVertical, CoordX: 0, CoordY: 3},
			},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.grid.GetConnectedPipes(testCase.tile))
	}
}

func TestGetFurthestPipeFromStartStepsCount(t *testing.T) {
	type testCase struct {
		grid     TheGrid
		expected int
	}

	testCases := []testCase{
		{
			grid:     Grid1,
			expected: 4,
		},
		{
			grid:     Grid2,
			expected: 8,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.grid.GetFurthestPipeFromStartStepsCount())
	}
}
