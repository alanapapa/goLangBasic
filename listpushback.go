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
	res := &NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = res
	} else {
		l.Tail.Next = res
	}
	l.Tail = res
}
