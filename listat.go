package piscine

// import (
// 	"fmt"
// )

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
// 	res := &NodeL{
// 		Data: data,
// 	}
// 	if l.Head == nil {
// 		l.Head = res
// 	} else {
// 		l.Tail.Next = res
// 	}
// 	l.Tail = res
// }

func ListAt(l *NodeL, nbr int) *NodeL {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}
	if nbr <= index {
		for l != nil {
			if count == nbr {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
