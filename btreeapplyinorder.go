package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Visit left subtree first
	BTreeApplyInorder(root.Left, f)

	// Apply function to the current node's data
	f(root.Data)

	// Visit right subtree next
	BTreeApplyInorder(root.Right, f)
}
