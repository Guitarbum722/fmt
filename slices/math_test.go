package slices

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

var uint8Cases = []struct {
	input       [][]uint8
	expectedLen int
	expected    []uint8
}{
	{
		[][]uint8{{2, 3, 4}, {5, 6, 10}},
		3,
		[]uint8{7, 9, 14},
	},
	{
		[][]uint8{{2, 3, 4, 5, 6, 7}, {5, 6, 10, 11, 40}},
		5,
		[]uint8{7, 9, 14, 16, 46},
	},
}

func TestAddUint8s(t *testing.T) {
	for _, tt := range uint8Cases {
		got := AddUint8s(tt.input[0], tt.input[1])
		if len(got) != tt.expectedLen {
			t.Fatalf("Len of AddUint8s(%q, %q) = %q, want %q", tt.input[0], tt.input[1], len(got), tt.expectedLen)
		}
		for i := range got {
			if tt.expected[i] != got[i] {
				t.Fatalf("AddUint8s(%q, %q) = %q, want %q", tt.input[0], tt.input[1], got, tt.expected)
			}
		}
	}
}
