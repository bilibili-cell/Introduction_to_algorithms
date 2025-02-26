package Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Bst struct {
	Head *TreeNode
}

func NewBst() *Bst {
	return &Bst{}
}

func (bst *Bst) Insert(val int) {
	if bst.Head == nil {
		bst.Head = &TreeNode{Val: val}
	} else {
		insertNode(bst.Head, val)
	}
}

func insertNode(node *TreeNode, val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			insertNode(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			insertNode(node.Right, val)
		}
	}
}
