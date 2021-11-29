// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"errors"
	"fmt"
)

type length struct {
	start int
	end   int
}

func main() {
	testSlice := []int{0, 14, 6, 2, 10, 9, 4, 1, 10, 3}
	fmt.Println(maxArea(testSlice))
}

func maxArea(heights []int) int {
	n := len(heights)
	max := maxOf(heights)
	min := minOf(heights)
	heightMap := make(map[int]length, n)
	err := error(nil)

	if n >= 2 && n <= 10e5 {
		err = errors.New("Max N size constrain exceeded")
	}

	if max >= 10e4 || min <= 0 {
		err = errors.New("Max height size constrain exceeded")
	}

	if err != nil {
		panic(err.Error())
	}

	startHeight := max

	for startHeight >= 1 {
		potentioalStart := 0
		potentioalStartSet := false
		for k, v := range heights {
			heightLength, ok := heightMap[startHeight]

			if !potentioalStartSet && v >= startHeight {
				potentioalStart = k
				potentioalStartSet = true
			}

			if !ok && v == startHeight {
				heightMap[startHeight] = length{start: k, end: 0}
			}

			if ok && v >= startHeight {
				heightLength.end = k
				heightMap[startHeight] = heightLength
			}

		}

		if potentioalStartSet && heightMap[startHeight].end-heightMap[startHeight].start < heightMap[startHeight].end-potentioalStart {
			modHeight := heightMap[startHeight]
			if modHeight.end == 0 {
				modHeight.end = modHeight.start
				modHeight.start = potentioalStart
			} else {
				modHeight.start = potentioalStart
			}
			heightMap[startHeight] = modHeight
		}

		startHeight--
	}

	for k, height := range heights {
		heights[k] = heightMap[height].getSquare(height)
	}

	fmt.Println(heightMap)
	fmt.Println(heights)

	return maxOf(heights)
}

func (interval length) getSquare(heigth int) int {
	if interval.end == 0 {
		return 0
	}

	return heigth * (interval.end - interval.start)
}

func maxOf(vars []int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func minOf(vars []int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
