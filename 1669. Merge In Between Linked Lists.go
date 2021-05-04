package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	start1 := list1
	start2 := list2

	var currentNode, nextNode *ListNode
	currentNode = list1
	for i := 0; i <= b; i++ {
		nextNode = currentNode.Next
		if i == a-1 {
			currentNode.Next = start2
		}
		currentNode = nextNode

	}

	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = currentNode
	return start1
}
