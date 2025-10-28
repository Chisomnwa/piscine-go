package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	l.Tail = l.Head // The current head will become the new tail

	for current != nil {
		next := current.Next // Save the next node
		current.Next = prev  // REverse the link
		prev = current       // Move prev forward
		current = next       // Move current forward
	}

	l.Head = prev // The last processed node becomes the new head
}
