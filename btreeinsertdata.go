package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// If root is nil, create a new node and return it
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Traverse the tree to find the correct spot
	if data < root.Data {
		// Go to the left subtree
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		// Go to the right subtree
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}

	return root
}
