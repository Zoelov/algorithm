package array

import (
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{10},
			want: 1,
		},
		{
			name: "normal sequence",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4, // [1, 2, 3, 4]
		},
		{
			name: "sequence with duplicates",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9, // [0, 1, 2, 3, 4, 5, 6, 7, 8]
		},
		{
			name: "unsorted sequence",
			nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
			want: 7, // [-1, 0, 1, ... but wait: -1, 0, 1 is 3. 3,4,5,6,7,8,9 is 7. ]
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3},
			want: 3,
		},
		{
			name: "no consecutive sequence",
			nums: []int{1, 3, 5, 7},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutiveSequence(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutiveSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
