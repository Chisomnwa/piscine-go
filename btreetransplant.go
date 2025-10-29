package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// Case 1: if node is nil, nothing to replace
	if node == nil {
		return root
	}

	// Case 2: if node is the root  of the tree
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// Case 3: node is the left child of its parent
		node.Parent.Left = rplc
	} else {
		// Case 4: node is the right child of its parent
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
