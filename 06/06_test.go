package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRawInputToInput(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	expected := Input{
		Races: []Race{
			{
				timeDurationMs:   7,
				distanceRecordMm: 9,
			},
			{
				timeDurationMs:   15,
				distanceRecordMm: 40,
			},
			{
				timeDurationMs:   30,
				distanceRecordMm: 200,
			},
		},
	}

	assert.Equal(t, expected, ConvertRawInputToInput(input))
}

func TestComputePossibleRecordsCount(t *testing.T) {

	type test struct {
		race Race
		want int
	}

	tests := []test{
		{
			race: Race{
				timeDurationMs:   7,
				distanceRecordMm: 9,
			},
			want: 4,
		},
		{
			race: Race{
				timeDurationMs:   15,
				distanceRecordMm: 40,
			},
			want: 8,
		},
		{
			race: Race{
				timeDurationMs:   30,
				distanceRecordMm: 200,
			},
			want: 9,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, tc.race.ComputePossibleRecordsCount())
	}
}

func TestComputeAllPossibleRecordCount(t *testing.T) {
	input := Input{
		Races: []Race{
			{
				timeDurationMs:   7,
				distanceRecordMm: 9,
			},
			{
				timeDurationMs:   15,
				distanceRecordMm: 40,
			},
			{
				timeDurationMs:   30,
				distanceRecordMm: 200,
			},
		},
	}

	expected := 288

	assert.Equal(t, expected, input.ComputeAllPossibleRecordCount())
}
