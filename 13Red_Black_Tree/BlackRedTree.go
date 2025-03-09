package Red_Black_Tree

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Value  int
	Color  bool
	Left   *Node
	Right  *Node
	Parent *Node
}

type RedBlackTree struct {
	Head *Node
}

func (tree *RedBlackTree) LeftRotate(x *Node) {
	y := x.Right
	if y == nil {
		return
	}

	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.Head = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (tree *RedBlackTree) RightRotate(x *Node) {
	y := x.Left
	if y == nil {
		return
	}

	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.Head = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

func (tree *RedBlackTree) Insert(value int) {
	newNode := &Node{Value: value, Color: RED}
	if tree.Head == nil {
		tree.Head = newNode
		tree.Head.Color = BLACK
		return
	}

	insertNode(tree.Head, newNode)
	tree.fixInsert(newNode)
}

func insertNode(root, node *Node) {
	if node.Value < root.Value {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			insertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			insertNode(root.Right, node)
		}
	}
}

func (tree *RedBlackTree) fixInsert(node *Node) {
	for node != tree.Head && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RED {
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					tree.LeftRotate(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				tree.RightRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RED {
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					tree.RightRotate(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				tree.LeftRotate(node.Parent.Parent)
			}
		}
	}
	tree.Head.Color = BLACK
}

func (tree *RedBlackTree) Delete(value int) {
	node := searchNode(tree.Head, value)
	if node == nil {
		return
	}
	deleteColor := node.Color

	var replaceNode *Node
	if node.Left == nil {
		replaceNode = node.Right
		transplant(tree, node, node.Right)
	} else if node.Right == nil {
		replaceNode = node.Left
		transplant(tree, node, node.Left)
	} else {
		successor := treeMinimum(node.Right)
		deleteColor = successor.Color
		replaceNode = successor.Right
		if successor.Parent == node {
			if replaceNode != nil {
				replaceNode.Parent = successor
			}
		} else {
			transplant(tree, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		transplant(tree, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
		successor.Color = node.Color
	}
	if deleteColor == BLACK {
		tree.fixDelete(replaceNode)
	}
}

func searchNode(root *Node, value int) *Node {
	if root == nil || root.Value == value {
		return root
	}
	if value < root.Value {
		return searchNode(root.Left, value)
	}
	return searchNode(root.Right, value)
}

func transplant(tree *RedBlackTree, u, v *Node) {
	if u.Parent == nil {
		tree.Head = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func treeMinimum(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (tree *RedBlackTree) fixDelete(node *Node) {
	for node != tree.Head && (node == nil || node.Color == BLACK) {
		if node == node.Parent.Left {
			sibling := node.Parent.Right
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				tree.LeftRotate(node.Parent)
				sibling = node.Parent.Right
			}
			if (sibling.Left == nil || sibling.Left.Color == BLACK) && (sibling.Right == nil || sibling.Right.Color == BLACK) {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Right == nil || sibling.Right.Color == BLACK {
					if sibling.Left != nil {
						sibling.Left.Color = BLACK
					}
					sibling.Color = RED
					tree.RightRotate(sibling)
					sibling = node.Parent.Right
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				if sibling.Right != nil {
					sibling.Right.Color = BLACK
				}
				tree.LeftRotate(node.Parent)
				node = tree.Head
			}
		} else {
			sibling := node.Parent.Left
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				tree.RightRotate(node.Parent)
				sibling = node.Parent.Left
			}
			if (sibling.Left == nil || sibling.Left.Color == BLACK) && (sibling.Right == nil || sibling.Right.Color == BLACK) {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Left == nil || sibling.Left.Color == BLACK {
					if sibling.Right != nil {
						sibling.Right.Color = BLACK
					}
					sibling.Color = RED
					tree.LeftRotate(sibling)
					sibling = node.Parent.Left
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				if sibling.Left != nil {
					sibling.Left.Color = BLACK
				}
				tree.RightRotate(node.Parent)
				node = tree.Head
			}
		}
	}
	if node != nil {
		node.Color = BLACK
	}
}
