package main

import (
	"fmt"
	"testing"
)

func TestN053MaxSubarray(t *testing.T) {
	var n053 N053MaxSubarray
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{-2, -1}
	fmt.Print("N053MaxSubarray:\t", n053.maxSubArray(nums1), " ", n053.maxSubArray(nums2), "\n")
}
