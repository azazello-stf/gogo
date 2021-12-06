package main

import "fmt"

func main() {
	fmt.Println(numSquares(43))
}

func numSquares(n int) int {
	perfectSquares := make([]int, 0, 100)
	counter := 0
	currentNumber := n

	var square int
	for square <= n {
		perfectSquares = append(perfectSquares, square)
		counter++
		square = counter * counter
	}

	var perfectSquaresResults int
	try := 0
	for try < len(perfectSquares)-1 {
		currentNumber = n
		var perfectSquaresCounter int
		currentDivider := len(perfectSquares) - try - 1
		for currentNumber >= 1 && currentDivider >= 1 {
			ceilPerfectNumbers := currentNumber / perfectSquares[currentDivider]
			currentNumber -= perfectSquares[currentDivider] * ceilPerfectNumbers
			perfectSquaresCounter += ceilPerfectNumbers
			currentDivider--
		}
		if try == 0 || perfectSquaresCounter != 0 && perfectSquaresCounter < perfectSquaresResults {
			perfectSquaresResults = perfectSquaresCounter
		}
		try++
	}

	return perfectSquaresResults
}
