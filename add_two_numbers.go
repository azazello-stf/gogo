// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	numberOne := createListNode(567)
	numberTwo := createListNode(433)
	res := add(*numberOne, *numberTwo)

	fmt.Print(res.getSlice())
}

func createListNode(number int) *ListNode {
	devider := 10
	var prevNode *ListNode = nil
	var firstNode *ListNode = nil
	_ = prevNode

	for number != 0 {
		digit := number % devider

		newNode := ListNode{digit, nil}

		if firstNode == nil {
			firstNode = &newNode
		} else {
			prevNode.Next = &newNode
		}

		prevNode = &newNode
		number /= devider
	}

	return firstNode
}

func add(n1 ListNode, n2 ListNode) *ListNode {
	var prevNode *ListNode = nil
	var firstNode *ListNode = nil

	addition := 0
	_ = addition

	for n1.Next != nil || n2.Next != nil || addition != 0 {
		sum := n1.Val + n2.Val + addition

		if addition != 0 {
			addition = 0
		}

		if sum >= 10 {
			addition = 1
			sum -= 10
		}

		newNode := ListNode{sum, nil}

		if prevNode != nil {
			prevNode.Next = &newNode
		} else {
			firstNode = &newNode
		}

		prevNode = &newNode

		if n1.Next != nil {
			n1 = *n1.Next
		} else {
			n1.Val = 0
		}

		if n2.Next != nil {
			n2 = *n2.Next
		} else {
			n2.Val = 0
		}
	}

	return firstNode
}

func (node ListNode) getSlice() []int {
	resList := []int{}

	for true {
		resList = append(resList, node.Val)

		if node.Next != nil {
			node = *node.Next
		} else {
			break
		}
	}

	return resList
}
