package array

/*
 给你一个整数数组，问其中有多少个连续的子数组，它们的和正好等于K
 [-1,0,1, 2,-1, 3] k = 2
*/

func SubArraySum(arr []int, k int) int {
	m := make(map[int]int, len(arr))
	m[0] = 1

	count := 0
	sum := 0
	for i := range len(arr) {
		sum += arr[i]
		value, ok := m[sum-k]
		if ok {
			count += value
		}

		m[sum]++
	}

	return count
}
