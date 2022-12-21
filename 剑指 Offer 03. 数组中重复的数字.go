package main

// func findRepeatNumber(nums []int) int {
// 	m := map[int]bool{}
// 	for i := 0; i < len(nums); i++ {
// 		if m[nums[i]] {
// 			return nums[i]
// 		}
// 		m[nums[i]] = true
// 	}
// 	return 0
// }

func findRepeatNumber(nums []int) int {
	len := len(nums)
	array := make([]int, len)
	for i := 0; i < len; i++ {
		array[nums[i]]++
		if array[nums[i]] >= 2 {
			return array[nums[i]]
		}
	}
	return -1
}
