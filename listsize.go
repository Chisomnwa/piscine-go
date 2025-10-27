package piscine

type NOdeL struct {
	Data interface{}
	Next *NodeL
}

type LIst struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0      // start a counter
	it := l.Head    // start from the first node
	for it != nil { // loop until we reach the end (nil)
		count++      // increase count for each node
		it = it.Next // move to the next node
	}
	return count // return the total number of nodes
}
