package box

import (
	"testing"
)

var primeTests = []struct {
	input    int
	expected bool
}{
	{
		3,
		true,
	},
	{
		28,
		false,
	},
	{
		140,
		false,
	},
	{
		2,
		true,
	},
}

func TestIsPrime(t *testing.T) {
	for _, v := range primeTests {
		if primeNum := IsPrime(v.input); primeNum != v.expected {
			t.Fatalf("IsPrime(%d) = %v; want %v", v.input, primeNum, v.expected)
		}
	}
}
