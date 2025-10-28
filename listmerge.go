package piscine

func ListMerge(l1 *List, l2 *List) {
	// If the first list is empty, point it to the second list
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// If the second list is empty, nothing to merge
	if l2.Head == nil {
		return
	}

	// Connect l1’s tail to l2’s head
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
