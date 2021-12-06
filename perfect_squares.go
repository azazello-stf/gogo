package main

import "fmt"

func main() {
	fmt.Println(numSquares(48))
}

func numSquares(n int) int {
	calculated := make([]int, 0, n+1)
	calculated = append(calculated, 0)
	currentVal := 1
	devider := 1
	for currentVal <= n {
		for {
			square := devider * devider
			remain := currentVal % square
			nextSquare := (devider + 1) * (devider + 1)
			var counter, prevCounter, last int
			if remain > 0 {
				for i := devider - 1; i > 0; i-- {
					prevCounter = currentVal/(i*i) + calculated[currentVal%(i*i)]
					if i == devider-1 || prevCounter < last {
						last = prevCounter
					}
				}
				prevCounter = last
			}
			counter = currentVal/square + calculated[remain]
			if prevCounter > 0 && prevCounter < counter {
				counter = prevCounter
			}

			calculated = append(calculated, counter)
			if currentVal+1 == nextSquare || currentVal > n {
				currentVal++
				break
			}
			currentVal++
		}
		devider++
	}

	return calculated[n]
}
