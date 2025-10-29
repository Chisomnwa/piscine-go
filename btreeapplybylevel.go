package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// create a queue (slide of *TreeNode)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// take the first node from queue
		current := queue[0]
		queue = queue[1:] // dequeue

		// apply the function to the curent node
		f(current.Data)

		// enqueue children (left first, then right)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
