package main

type N053MaxSubarray struct {
}

func (this *N053MaxSubarray) maxSubArray(nums []int) int {
	count := len(nums)
	start, end := 0, 1
	sum := nums[0]
	maxSum := sum
	current := 0
	for start <= end && start < count && end < count {
		current = nums[end]
		if current >= 0 {
			if sum > 0 {
				sum += current
			} else { // If sum <= 0, need not add current, re-calculate
				start = end
				sum = current
			}
		} else {
			if current > sum { // Handle all negative numbers
				start = end
				sum = current
			} else {
				sum += current
			}
		}
		end++
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
