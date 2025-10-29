package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Visit left subtree first
	BTreeApplyPostorder(root.Left, f)

	// Then visit right subtree
	BTreeApplyPostorder(root.Right, f)

	// Finally, apply function to current node's data
	f(root.Data)
}
