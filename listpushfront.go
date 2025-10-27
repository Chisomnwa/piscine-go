package piscine

func ListPushFront(l *List, data interface{}) {
	newMode := &NodeL{Data: data} // create a new node

	if l.Head == nil { // if the list is empty
		l.Head = newMode
		l.Tail = newMode
	} else {
		newMode.Next = l.Head // link the new node to the current head
		l.Head = newMode      // update head
	}
}
