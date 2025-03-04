package Red_Black_Tree

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node // 添加 Parent 指针用于维护父节点
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
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	}
	x.Right = y
	y.Parent = x

}
