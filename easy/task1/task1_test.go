package task1

import (
	"fmt"
	"testing"
)

func TestTask1(t *testing.T) {
	tests := []struct {
		money    int
		children int
		expected int
	}{
		{
			20,
			3,
			1,
		},
		{
			16,
			2,
			2,
		},
		{
			12,
			2,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.money), func(t *testing.T) {
			result := distMoney(tt.money, tt.children)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
