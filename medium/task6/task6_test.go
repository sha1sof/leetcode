package task6

import (
	"fmt"
	"testing"
)

func TestTask6(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{
			10,
			4,
		},
		{
			0,
			0,
		},
		{
			1,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			result := countPrimes(tt.n)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
