package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	ListPushBack(link, "three")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "1")

// 	fmt.Println(ListLast(link))
// 	fmt.Println(ListLast(link2))
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
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

func ListLast(l *List) interface{} {
	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		return nil
	}
	return l.Head.Data
}

/*panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x50a271]

goroutine 7 [running]:
testing.tRunner.func1(0xc0000b8100)
	/usr/local/go/src/testing/testing.go:874 +0x3a3
panic(0x530180, 0x653440)
	/usr/local/go/src/runtime/panic.go:679 +0x1b2
_/jail/student.ListLast(0xc000052b40, 0x523960, 0xc000018680)
	/jail/student/listlast.go:27 +0x11
_/jail_test.TestListLast(0xc0000b8100)
	/jail/listlast_test.go:74 +0x1d2
testing.tRunner(0xc0000b8100, 0x565758)
	/usr/local/go/src/testing/testing.go:909 +0xc9
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:960 +0x350
exit status 2
FAIL	_/jail	0.004s*/
