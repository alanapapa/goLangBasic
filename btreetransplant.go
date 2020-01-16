package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
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

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	BTreeApplyInorder(root.Left, f)
// 	f(root.Data)
// 	BTreeApplyInorder(root.Right, f)
// }

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if elem < root.Data {
// 		return BTreeSearchItem(root.Left, elem)
// 	} else if elem > root.Data {
// 		return BTreeSearchItem(root.Right, elem)
// 	} else {
// 		return root
// 	}
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	node := BTreeSearchItem(root, "1")
// 	replacement := &TreeNode{Data: "55"}
// 	root = BTreeTransplant(root, node, replacement)
// 	BTreeApplyInorder(root, fmt.Println)
// }

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}

// func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	replacement := node
// 	if node.Parent == nil {
// 		root = rplc

// 	} else if node == node.Parent.Left {
// 		replacement.Parent.Left = rplc
// 	} else if node < node.Parent.Left{
// 		root = nil
// 	} else {
// 		replacement.Parent.Right = rplc
// 	}
// 	replacement.Parent = node.Parent

// 	return root
// }

/*Tree:
 04
├── 55
│   ├── 60
│   └── 33
│       └── 12
│           └── 15
└── 01
    └── 02
        └── 03

--- FAIL: TestBTreeTransplant (0.00s)
    btreetransplant_test.go:108: Expected
        Node: 04 Parent: nil
        Node: 01 Parent: 04
        Node: 02 Parent: 01
        Node: 03 Parent: 02
        Node: 55 Parent: 04
        Node: 33 Parent: 55
        Node: 12 Parent: 33
        Node: 15 Parent: 12
        Node: 60 Parent: 55
         instead of
        Node: 04 Parent: nil
        Node: 01 Parent: 04
        Node: 02 Parent: 01
        Node: 03 Parent: 02
        Node: 55 Parent: nil
        Node: 33 Parent: 55
        Node: 12 Parent: 33
        Node: 15 Parent: 12
        Node: 60 Parent: 55

FAIL
exit status 1
FAIL	_/jail	0.002s*/
