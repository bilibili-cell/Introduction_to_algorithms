package Binary_Search_Tree

import "math"

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

func (bst *Bst) Search(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}
	if val < node.Val {
		return bst.Search(node.Left, val)
	} else {
		return bst.Search(node.Right, val)
	}
}

func (bst *Bst) Minimum(node *TreeNode) int {
	if node == nil {
		return math.MinInt
	}
	if node.Left != nil {
		return bst.Minimum(node.Left)
	} else {
		return node.Val
	}
}

func (bst *Bst) Maximum(node *TreeNode) int {
	if node == nil {
		return math.MaxInt
	}
	if node.Right != nil {
		return bst.Maximum(node.Right)
	} else {
		return node.Val
	}
}

func (bst *Bst) TreeSuccessor(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return bst.TreeNodeMinimum(node.Right)
	}

	var successor *TreeNode
	current := bst.Head
	for current != nil {
		if node.Val < current.Val {
			successor = current
			current = current.Left
		} else if node.Val > current.Val {
			current = current.Right
		} else {
			break
		}
	}
	return successor
}

func (bst *Bst) TreeNodeMinimum(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *Bst) Delete(val int) {
	bst.Head = bst.deleteNode(bst.Head, val)
}

func (bst *Bst) deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = bst.deleteNode(node.Left, val)
	} else if val > node.Val {
		node.Right = bst.deleteNode(node.Right, val)
	} else {
		// 找到要删除的节点
		if node.Left == nil && node.Right == nil {
			// 节点是叶子节点
			return nil
		} else if node.Left == nil {
			// 节点只有右子节点
			return node.Right
		} else if node.Right == nil {
			// 节点只有左子节点
			return node.Left
		} else {
			// 节点有两个子节点
			minRight := bst.Minimum(node.Right)
			node.Val = minRight
			node.Right = bst.deleteNode(node.Right, minRight)
		}
	}
	return node
}
