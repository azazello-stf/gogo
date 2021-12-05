package main

import "fmt"

func main() {
	input := []int{1, 3}
	target := 2
	fmt.Println(searchInsert(input, target))
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	start, end := 0, length
	key := end / 2
	for {
		if key-1 >= 0 && key+1 < length && nums[key-1] < target && nums[key+1] >= target {
			return key
		} else if key == length-1 && nums[key] < target {
			return key + 1
		} else if key == 0 {
			return key
		}

		if nums[key] > target {
			end = key
			key = (end - start) / 2
		} else if nums[key] < target {
			start = key
			key = key + (end-start)/2
		}
	}
}
