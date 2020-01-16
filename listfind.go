package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "hello1")
// 	ListPushBack(link, "hello2")
// 	ListPushBack(link, "hello3")

// 	found := ListFind(link, interface{}("hello2"), CompStr)

// 	fmt.Println(found)
// 	fmt.Println(*found)
// }

// func ListPushBack(l *List, data interface{}) {
// 				res := &NodeL{
// 					Data: data,
// 				}
// 				if l.Head == nil {
// 					l.Head = res
// 				} else {
// 					l.Tail.Next = res
// 				}
// 				l.Tail = res
// 			}

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := l.Head
	for iterator != nil {
		if comp(iterator.Data, ref) {
			return &iterator.Data
		}

		iterator = iterator.Next
	}
	return nil
}
