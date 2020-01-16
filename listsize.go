package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushFront(l *List, data interface{}) {
// 	res := &NodeL{
// 		Data: data,
// 	}
// 	if l.Head == nil {
// 		l.Head = res
// 	} else {
// 		res.Next = l.Head
// 	}
// 	l.Head = res
// }

func ListSize(l *List) int {
	count := 0

	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
