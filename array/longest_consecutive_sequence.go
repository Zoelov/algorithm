package array

/*
给定一个未排序的整数数组nums,找出数字连续的最长序列（不要求序列元素在原数组中连续)的长度。
请设计并实现时间复杂度为O(n)的算法解决此问题。
example:
Input: nums = [100,4,200,1,3,2]
Output: 4
*/
func LongestConsecutiveSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	m := make(map[int]struct{})
	for i := range len(nums) {
		m[nums[i]] = struct{}{}
	}

	for i := range len(nums) {
		if _, ok := m[nums[i]-1]; ok {
			continue
		}

		cur := 1
		for {
			if _, ok := m[nums[i]+cur]; ok {
				cur++
			} else {
				if cur > max {
					max = cur
				}
				break
			}
		}
	}

	return max
}
