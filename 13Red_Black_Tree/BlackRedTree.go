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
