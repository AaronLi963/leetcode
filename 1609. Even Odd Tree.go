package main

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	next := []*TreeNode{}
	current := []*TreeNode{root}
	isOdd := false
	for {
		for i := range current {
			if isOdd {
				if current[i].Val%2 != 0 {
					return false
				}
			} else {
				if current[i].Val%2 != 1 {
					return false
				}
			}

			if i > 0 {
				if isOdd {
					if !(current[i].Val < current[i-1].Val) {
						return false
					}
				} else {
					if !(current[i].Val > current[i-1].Val) {
						return false
					}
				}
			}
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}
		}
		if len(next) == 0 {
			break
		}

		isOdd = !isOdd
		current = next
		next = []*TreeNode{}
	}
	return true
}
