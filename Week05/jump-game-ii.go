package main


func jump(nums []int) int {
	target := len(nums) - 1
	times := 0
	for target > 0 {
		for i := 0; i < target; i++ {
			if i + nums[i] >= target {
				target = i
				times++
				break
			}
		}
	}
	return times
}
