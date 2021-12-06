package main

import "fmt"

func main() {
	input := []int{1, 3}
	target := 3
	fmt.Println(searchInsert(input, target))
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	end := length - 1
	position := end / 2
	step := 1
	if step > 1 {
		step = position
	}
	prevPosition := length + 1
	for {
		if position == end && nums[position] < target {
			return end + 1
		} else if position+1 <= length-1 && nums[position] < target && nums[position+1] > target {
			return position + 1
		} else if position-1 >= 0 && nums[position] > target && nums[position-1] < target {
			return position
		} else if prevPosition == position && step <= 1 {
			return position
		}

		prevPosition = position

		if position > 0 && nums[position] > target {
			position -= step
		} else if position+step <= length-1 && nums[position] < target {
			position += step
		}
		if step > 1 {
			step /= 2
		}
	}
}
