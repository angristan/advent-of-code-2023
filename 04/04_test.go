package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInputToListOfCards(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	want := []Card{
		{WinningNumbers: []CardNumber{41, 48, 83, 86, 17},
			MyNumbers: []CardNumber{83, 86, 6, 31, 17, 9, 48, 53}},
		{WinningNumbers: []CardNumber{13, 32, 20, 16, 61},
			MyNumbers: []CardNumber{61, 30, 68, 82, 17, 32, 24, 19}},
		{WinningNumbers: []CardNumber{1, 21, 53, 59, 44},
			MyNumbers: []CardNumber{69, 82, 63, 72, 16, 21, 14, 1}},
		{WinningNumbers: []CardNumber{41, 92, 73, 84, 69},
			MyNumbers: []CardNumber{59, 84, 76, 51, 58, 5, 54, 83}},
		{WinningNumbers: []CardNumber{87, 83, 26, 28, 32},
			MyNumbers: []CardNumber{88, 30, 70, 12, 93, 22, 82, 36}},
		{WinningNumbers: []CardNumber{31, 18, 13, 56, 72},
			MyNumbers: []CardNumber{74, 77, 10, 23, 35, 67, 36, 11}},
	}

	cards := ConvertInputToListOfCards(input)

	assert.Equal(t, want, cards)
}

func TestComputePoints(t *testing.T) {
	type testInput struct {
		card      Card
		wantScore int
	}

	input := []testInput{
		{card: Card{WinningNumbers: []CardNumber{41, 48, 83, 86, 17},
			MyNumbers: []CardNumber{83, 86, 6, 31, 17, 9, 48, 53}}, wantScore: 8},
		{card: Card{WinningNumbers: []CardNumber{13, 32, 20, 16, 61},
			MyNumbers: []CardNumber{61, 30, 68, 82, 17, 32, 24, 19}}, wantScore: 2},
		{card: Card{WinningNumbers: []CardNumber{1, 21, 53, 59, 44},
			MyNumbers: []CardNumber{69, 82, 63, 72, 16, 21, 14, 1}}, wantScore: 2},
		{card: Card{WinningNumbers: []CardNumber{41, 92, 73, 84, 69},
			MyNumbers: []CardNumber{59, 84, 76, 51, 58, 5, 54, 83}}, wantScore: 1},
		{card: Card{WinningNumbers: []CardNumber{87, 83, 26, 28, 32},
			MyNumbers: []CardNumber{88, 30, 70, 12, 93, 22, 82, 36}}, wantScore: 0},
		{card: Card{WinningNumbers: []CardNumber{31, 18, 13, 56, 72},
			MyNumbers: []CardNumber{74, 77, 10, 23, 35, 67, 36, 11}}, wantScore: 0},
	}

	for _, ti := range input {
		assert.Equal(t, ti.wantScore, ti.card.ComputePoints())
	}
}

func TestComputeTotalPoints(t *testing.T) {
	type testInput struct {
		elfStack      ElfStack
		wantTotalWins int
	}

	input := []testInput{
		{elfStack: ElfStack{
			{WinningNumbers: []CardNumber{41, 48, 83, 86, 17},
				MyNumbers: []CardNumber{83, 86, 6, 31, 17, 9, 48, 53}},
			{WinningNumbers: []CardNumber{13, 32, 20, 16, 61},
				MyNumbers: []CardNumber{61, 30, 68, 82, 17, 32, 24, 19}},
			{WinningNumbers: []CardNumber{1, 21, 53, 59, 44},
				MyNumbers: []CardNumber{69, 82, 63, 72, 16, 21, 14, 1}},
			{WinningNumbers: []CardNumber{41, 92, 73, 84, 69},
				MyNumbers: []CardNumber{59, 84, 76, 51, 58, 5, 54, 83}},
			{WinningNumbers: []CardNumber{87, 83, 26, 28, 32},
				MyNumbers: []CardNumber{88, 30, 70, 12, 93, 22, 82, 36}},
			{WinningNumbers: []CardNumber{31, 18, 13, 56, 72},
				MyNumbers: []CardNumber{74, 77, 10, 23, 35, 67, 36, 11}},
		}, wantTotalWins: 13},
	}

	for _, ti := range input {
		assert.Equal(t, ti.wantTotalWins, ti.elfStack.ComputeTotalPoints())
	}
}
