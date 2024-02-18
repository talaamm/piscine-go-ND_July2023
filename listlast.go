package piscinego

type NodeL7 struct {
	Data interface{}
	Next *NodeL
}

type List7 struct {
	Head *NodeL
	Tail *NodeL
}

/*
func ListLast(l *List) interface{} {
	currentNode := l.Head
	var lastData interface{}

	for currentNode != nil {
		lastData = currentNode.Data
		currentNode = currentNode.Next
	}
	return lastData
}
*/

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
