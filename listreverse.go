package piscinego

type NodeL8 struct {
	Data interface{}
	Next *NodeL
}

type List8 struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	temp := l.Head
	current := l.Head
	prev := l.Head
	prev = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = temp
	l.Tail.Next = nil
}

// l,kmijuhygvtrf
