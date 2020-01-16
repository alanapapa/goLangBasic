package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "1")
// 	ListPushBack(link, "2")
// 	ListPushBack(link, "3")
// 	ListPushBack(link, "5")

// 	ListForEach(link, Add2_node)

// 	it := link.Head
// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}
// }

// func ListPushBack(l *List, data interface{}) {
// 		res := &NodeL{
// 			Data: data,
// 		}
// 		if l.Head == nil {
// 			l.Head = res
// 		} else {
// 			l.Tail.Next = res
// 		}
// 		l.Tail = res
// 	}

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func ListForEach(l *List, f func(*NodeL)) {
	res := l.Head
	for res != nil {
		f(res)
		res = res.Next
	}
}
