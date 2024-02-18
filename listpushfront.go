package piscinego

type NodeL2 struct {
	Data interface{}
	Next *NodeL
}

type List2 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head // Update the Next pointer of the new node to the current head
		l.Head = newNode
	}
}
