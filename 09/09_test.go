package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRawInputToReport(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	expected := Report{
		Histories: []History{
			{Values: []int{0, 3, 6, 9, 12, 15}},
			{Values: []int{1, 3, 6, 10, 15, 21}},
			{Values: []int{10, 13, 16, 21, 30, 45}},
		},
	}

	assert.Equal(t, expected, ConvertRawInputToReport(input))
}

func TestComputeNextValue(t *testing.T) {
	type testCase struct {
		history           History
		expectedNextValue int
	}

	testCases := []testCase{
		{
			history: History{
				Values: []int{0, 3, 6, 9, 12, 15},
			},
			expectedNextValue: 18,
		},
		{
			history: History{
				Values: []int{1, 3, 6, 10, 15, 21},
			},
			expectedNextValue: 28,
		},
		{
			history: History{
				Values: []int{10, 13, 16, 21, 30, 45},
			},
			expectedNextValue: 68,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedNextValue, testCase.history.ComputeNextValue())
	}
}

func TestComputeSumOfNextValues(t *testing.T) {
	input := Report{
		Histories: []History{
			{Values: []int{0, 3, 6, 9, 12, 15}},
			{Values: []int{1, 3, 6, 10, 15, 21}},
			{Values: []int{10, 13, 16, 21, 30, 45}},
		},
	}

	expected := 18 + 28 + 68

	assert.Equal(t, expected, input.ComputeSumOfNextValues())
}
