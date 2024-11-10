package task2

import (
	"fmt"
	"testing"
)

func TestTask1(t *testing.T) {
	tests := []struct {
		prices   []int
		money    int
		expected int
	}{
		{
			[]int{1, 2, 2},
			3,
			0,
		},
		{
			[]int{3, 2, 3},
			3,
			3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.money), func(t *testing.T) {
			result := buyChoco(tt.prices, tt.money)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
