package task5

import (
	"fmt"
	"testing"
)

func TestTask5(t *testing.T) {
	tests := []struct {
		price    []int
		expected int
	}{
		{
			[]int{1, 2, 3, 0, 2},
			3,
		},
		{
			[]int{1},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.price), func(t *testing.T) {
			result := maxProfit(tt.price)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
