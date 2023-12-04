package main

import (
	"testing"
)

func TestComputeCalibrationSum(t *testing.T) {
	type test struct {
		input []string
		want  int
	}

	tests := []test{
		{
			[]string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			142,
		},
		{
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			281,
		},
	}

	for _, tc := range tests {
		got := ComputeCalibrationSum(tc.input)

		if got != tc.want {
			t.Errorf("Expected sum to be %d, got %d", tc.want, got)
		}

	}
}

func BenchmarkComputeCalibrationSum(b *testing.B) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	for i := 0; i < b.N; i++ {
		ComputeCalibrationSum(input)
	}
}
