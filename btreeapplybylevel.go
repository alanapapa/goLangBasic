package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyByLevel(root, fmt.Println)
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		root.Left.Parent = root
// 	} else if data > root.Data {
// 		root.Right = BTreeInsertData(root.Right, data)
// 		root.Right.Parent = root
// 	}
// 	return root
// }

// func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	f(root.Data)
// 	BTreeApplyByLevel(root.Left, f)
// 	BTreeApplyByLevel(root.Right, f)
// }

func LevelMove(root *TreeNode, f func(...interface{}) (int, error), lv int) {

	if root == nil {
		return
	}
	if lv == 0 {
		f(root.Data)
	} else if lv > 0 {
		LevelMove(root.Left, f, lv-1)
		LevelMove(root.Right, f, lv-1)
	}
}
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	level := BTreeLevelCount(root)
	for i := 0; i < level; i++ {
		LevelMove(root, f, i)
	}
}

// func BTreeLevelCount(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left := BTreeLevelCount(root.Left)
// 	right := BTreeLevelCount(root.Right)

// 	if left > right {
// 		return left + 1
// 	}
// 	return right + 1
// }

/*--- FAIL: TestBTreeApplyByLevel (0.00s)
    challenge.go:103: BTreeApplyByLevel(
        04
        ├── 07
        │   ├── 12
        │   │   └── 10
        │   └── 05
        └── 01
            └── 02
                └── 03
        )
         prints "0401020307051210" instead of "0401070205120310"
FAIL
exit status 1
FAIL	_/jail	0.002s*/
