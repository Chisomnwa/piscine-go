package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Visit root first
	f(root.Data)

	// Then left subtree
	BTreeApplyPreorder(root.Left, f)

	// Then right subtree
	BTreeApplyPreorder(root.Right, f)
}
