package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
