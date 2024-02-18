package piscinego

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {
	/*currentNode := l.Head

	for currentNode != nil {
		/*nextNode := currentNode.Next
		currentNode.Next = nil
		currentNode = nextNode
		currentNode.Data = nil
		currentNode = currentNode.Next
	}
	*/

	// After clearing all nodes, update the list's Head and Tail pointers to nil
	l.Head = nil
	// l.Tail = nil
}
