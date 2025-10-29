package piscine

// BTreeDeleteNode deletes a node from the binary search tree
// and returns the new root of the tree.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1: Node has no left child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)

		// Case 2: Node has no right child
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)

		// Case 3: Node has two children
	} else {
		// Find the minimum node in the right subtree (the inorder successor)
		successor := node.Right
		for successor.Left != nil {
			successor = successor.Left
		}

		// If the successor is not the node's immediate right child
		if successor.Parent != node {
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}

		// Replace node with successor
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
	}

	return root
}
