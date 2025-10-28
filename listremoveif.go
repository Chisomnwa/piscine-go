package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove matching nodes from the start of the list
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list becomes empty, reset Tail too
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Now remove matching nodes from the rest of the list
	prev := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Data == data_ref {
			prev.Next = current.Next
			if current == l.Tail { // if we remove the tail
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
