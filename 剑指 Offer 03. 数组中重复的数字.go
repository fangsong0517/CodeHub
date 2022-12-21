package main

func findRepeatNumber(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return nums[i]
		}
		m[nums[i]] = true
	}
	return 0
}
