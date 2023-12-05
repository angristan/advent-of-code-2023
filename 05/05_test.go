package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInputToAlmanac(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	want := Almanac{
		Seeds: []Seed{79, 14, 55, 13},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	almanac := ConvertInputToAlmanac(input)

	assert.Equal(t, want, almanac)
}

func TestGetSeedsLocations(t *testing.T) {
	alamanac := Almanac{
		Seeds: []Seed{79, 14, 55, 13},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	expectedLocations := []int{82, 43, 86, 35}

	assert.Equal(t, alamanac.GetSeedsLocations(), expectedLocations)
}

func TestGetLowestLocationNumber(t *testing.T) {
	alamanac := Almanac{
		Seeds: []Seed{79, 14, 55, 13},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	expectedLowestLocationNumber := 35

	assert.Equal(t, alamanac.GetLowestLocationNumber(), expectedLowestLocationNumber)
}

func TestConvertInputToAlmanacV2(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	want := AlmanacV2{
		Seeds: []SeedV2{
			{Number: 79, Range: 14},
			{Number: 55, Range: 13},
		},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	almanac := ConvertInputToAlmanacV2(input)

	assert.Equal(t, want, almanac)
}

func TestGetLowestLocationNumberV2(t *testing.T) {
	alamanac := AlmanacV2{
		Seeds: []SeedV2{
			{Number: 79, Range: 14},
			{Number: 55, Range: 13},
		},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	expectedLowestLocationNumber := 46

	assert.Equal(t, alamanac.GetLowestLocationNumber(), expectedLowestLocationNumber)
}
func BenchmarkGetLowestLocationNumberV2(b *testing.B) {
	alamanac := AlmanacV2{
		Seeds: []SeedV2{
			{Number: 79, Range: 14},
			{Number: 55, Range: 13},
		},
		Maps: []Map{
			{Ranges: []Range{
				{DestinationIndex: 50, SourceIndex: 98, RangeLength: 2},
				{DestinationIndex: 52, SourceIndex: 50, RangeLength: 48}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 15, RangeLength: 37},
				{DestinationIndex: 37, SourceIndex: 52, RangeLength: 2},
				{DestinationIndex: 39, SourceIndex: 0, RangeLength: 15}},
			},
			{Ranges: []Range{
				{DestinationIndex: 49, SourceIndex: 53, RangeLength: 8},
				{DestinationIndex: 0, SourceIndex: 11, RangeLength: 42},
				{DestinationIndex: 42, SourceIndex: 0, RangeLength: 7},
				{DestinationIndex: 57, SourceIndex: 7, RangeLength: 4}},
			},
			{Ranges: []Range{
				{DestinationIndex: 88, SourceIndex: 18, RangeLength: 7},
				{DestinationIndex: 18, SourceIndex: 25, RangeLength: 70}},
			},
			{Ranges: []Range{
				{DestinationIndex: 45, SourceIndex: 77, RangeLength: 23},
				{DestinationIndex: 81, SourceIndex: 45, RangeLength: 19},
				{DestinationIndex: 68, SourceIndex: 64, RangeLength: 13}},
			},
			{Ranges: []Range{
				{DestinationIndex: 0, SourceIndex: 69, RangeLength: 1},
				{DestinationIndex: 1, SourceIndex: 0, RangeLength: 69}},
			},
			{Ranges: []Range{
				{DestinationIndex: 60, SourceIndex: 56, RangeLength: 37},
				{DestinationIndex: 56, SourceIndex: 93, RangeLength: 4}},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		alamanac.GetLowestLocationNumber()
	}

}
