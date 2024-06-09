package decimal

import "testing"

func TestPlus(t *testing.T) {
	tests := [][]string{
		{"123", "456"},
		{"0", "456"},
		{"456", "456"},
		{"1111111111", "4444444"},
		{"1111111111", "99999999999"},
	}
	for _, test := range tests {
		t.Log(test[0], test[1], Plus(test[0], test[1]))
	}
}
