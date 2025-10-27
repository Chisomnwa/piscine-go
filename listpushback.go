package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data} // Create a new node with the given data

	if l.Head == nil { // if the list is empty
		l.Head = newNode // the new node becomes both head
		l.Tail = newNode // and tail
	} else {
		l.Tail.Next = newNode // link the current tail to the new node
		l.Tail = newNode      // update the tail pointer to the new node
	}
}
