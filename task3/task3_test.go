package task3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTask3(t *testing.T) {
	tests := []struct {
		nums     []int
		queries  []int
		expected []int
	}{
		{
			[]int{4, 5, 2, 1},
			[]int{3, 10, 21},
			[]int{2, 3, 4},
		},
		{
			[]int{2, 3, 4, 5},
			[]int{1},
			[]int{0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			result := answerQueries(tt.nums, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
