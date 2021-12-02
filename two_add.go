// https://leetcode.com/problems/two-sum/submissions/
package main

import "fmt"

func main() {
	nums := []int{-1, -2, -3, -4, -5}
	fmt.Println(twoSum(nums, -8))
}

func twoSum(nums []int, target int) []int {
	var firstIndex, secondIndex int
	for k, v := range nums {
		firstIndex = k
		searchTarget := target - v
		for sk, sv := range nums {
			if sk > k && sv == searchTarget {
				secondIndex = sk
				break
			}
		}

		if secondIndex != 0 {
			break
		}
	}
	return []int{firstIndex, secondIndex}
}
