package decimal

import "testing"

func TestMinus(t *testing.T) {
	tests := [][]string{
		{"1", "2"},
		{"0", "2"},
		{"10", "288"},
		{"100000", "288"},
		{"289", "288"},
		{"38", "288"},
		{"3", "3"},
	}
	for _, test := range tests {
		t.Log(test[0], test[1], Minus(test[0], test[1]))
	}
}
