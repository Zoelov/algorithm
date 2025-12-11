package array

import "testing"

func TestKSubarray(t *testing.T) {
	data := []struct {
		arr      []int
		k        int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			3,
			2,
		},

		{
			[]int{1, 2, 3, 4, 5},
			5,
			2,
		},
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			[]int{-1, 0, 1, 2, -1, 3},
			2,
			5,
		},
	}

	for _, d := range data {
		actual := SubArraySum(d.arr, d.k)
		if actual != d.expected {
			t.Errorf("acutal:%d not equal:%d", actual, d.expected)
		}
	}
}
