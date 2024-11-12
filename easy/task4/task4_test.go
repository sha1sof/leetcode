package task4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTask4(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
			[]int{2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
			[]int{9, 4},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums1), func(t *testing.T) {
			result := intersection(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
