package bittracker_test

import (
	"github.com/snowmerak/bittracker"
	"testing"
)

func TestBitTracker_GetBit(t *testing.T) {
	bt := bittracker.NewBitTracker([]byte{0b01010101, 0b11111111, 0b00000000, 0b10101010})
	corrects := []bool{false, true, false, true, false, true, false, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	for i, answer := range corrects {
		if bt.GetBit(i) != answer {
			t.Errorf("Index %d Expected %t, got %t", i, answer, bt.GetBit(i))
		}
	}
}

func TestBitTracker_SetBit(t *testing.T) {
	bt := bittracker.NewBitTracker([]byte{0b01010101, 0b11111111, 0b00000000, 0b10101010})

	for i := 0; i < 32; i++ {
		if i&1 == 1 {
			bt.SetBit(i, true)
		} else {
			bt.SetBit(i, false)
		}
	}

	for i := 0; i < 32; i++ {
		if i&1 == 1 {
			if !bt.GetBit(i) {
				t.Errorf("Index %d Expected true, got false", i)
			}
		} else {
			if bt.GetBit(i) {
				t.Errorf("Index %d Expected false, got true", i)
			}
		}
	}
}

func TestBitTracker_GetRange(t *testing.T) {
	bt := bittracker.NewBitTracker([]byte{0b01010101, 0b11111111, 0b00000000, 0b10101010})

	tests := [][]int{
		{0, 8},
		{0, 16},
		{0, 24},
		{0, 32},
		{12, 20},
		{12, 28},
		{5, 26},
	}
	corrects := [][]byte{
		{0b01010101},
		{0b01010101, 0b11111111},
		{0b01010101, 0b11111111, 0b00000000},
		{0b01010101, 0b11111111, 0b00000000, 0b10101010},
		{0b11110000},
		{0b11110000, 0b00001010},
		{0b10111111, 0b11100000, 0b00010000},
	}

	for i, test := range tests {
		result := bt.GetRange(test[0], test[1])
		if len(result) != len(corrects[i]) {
			t.Errorf("Test %d Expected length %d, got %d, %v", i, len(corrects[i]), len(result), result)
			continue
		}

		for j, b := range result {
			if b != corrects[i][j] {
				t.Errorf("Test %d Index %d Expected %08b, got %08b", i, j, corrects[i][j], b)
			}
		}
	}
}
