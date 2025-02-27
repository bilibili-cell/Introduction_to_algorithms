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
func (bst *Bst) InorderTreeWalk(node *TreeNode, resultList *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		bst.InorderTreeWalk(node.Left, resultList)
	}
	*resultList = append(*resultList, node.Val)
	if node.Right != nil {
		bst.InorderTreeWalk(node.Right, resultList)
	}
}
func (bst *Bst) PreorderTreeWalk(node *TreeNode, resultList *[]int) {
	if node == nil {
		return
	}
	*resultList = append(*resultList, node.Val)
	if node.Left != nil {
		bst.PreorderTreeWalk(node.Left, resultList)
	}
	if node.Right != nil {
		bst.PreorderTreeWalk(node.Right, resultList)
	}
}

func (bst *Bst) PostorderTreeWalk(node *TreeNode, resultList *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		bst.PostorderTreeWalk(node.Left, resultList)
	}
	if node.Right != nil {
		bst.PostorderTreeWalk(node.Right, resultList)
	}
	*resultList = append(*resultList, node.Val)
}
