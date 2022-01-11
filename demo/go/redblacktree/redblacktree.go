package redblacktree

import "fmt"

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

func (color Color) output() string {
	switch color {
	case Red:
		return "RED"
	case Black:
		return "BLACK"
	default:
		return ""
	}
}

type Node struct {
	value int
	color Color

	leftTree, rightTree, parent *Node
}

// Grandparent 爷
func (node *Node) Grandparent() *Node {
	if node.parent == nil {
		return nil
	}
	return node.parent.parent
}

// Uncle 叔伯
func (node *Node) Uncle() *Node {
	if node.Grandparent() == nil {
		return nil
	}
	if node.parent == node.Grandparent().rightTree {
		return node.Grandparent().leftTree
	} else {
		return node.Grandparent().rightTree
	}
}

// Sibling 兄弟
func (node *Node) Sibling() *Node {
	if node.parent == nil {
		return nil
	}
	if node.parent.leftTree == node {
		return node.parent.rightTree
	} else {
		return node.parent.leftTree
	}
}

type RBT struct {
	root, NIL *Node
}

// 右翻转
// 涉及节点：本节点、父节点的左子节点、右节点、根节点、爷节点
// 移动方向：
func (t *RBT) rotateRight(node *Node) {
	gp := node.Grandparent()
	fa := node.parent
	rt := node.rightTree

	fa.leftTree = rt

	if rt != t.NIL {
		rt.parent = fa
	}
	node.rightTree = fa
	fa.parent = node

	if t.root == fa {
		t.root = rt
	}
	node.parent = gp

	if gp != nil {
		if gp.leftTree == fa {
			gp.leftTree = node
		} else {
			gp.rightTree = node
		}
	}
}

func (t *RBT) rotateLeft(node *Node) {
	if node.parent == nil {
		t.root = node
		return
	}
	gp := node.Grandparent()
	fa := node.parent
	lt := node.leftTree

	fa.rightTree = lt

	if lt != t.NIL {
		lt.parent = fa
	}
	node.leftTree = fa
	fa.parent = node

	if t.root == fa {
		t.root = node
	}
	node.parent = gp

	if gp != nil {
		if gp.leftTree == fa {
			gp.leftTree = node
		} else {
			gp.rightTree = node
		}
	}
}

func (t *RBT) inorder(node *Node, level int) {
	level++
	var prefix string
	for i := 1; i < level; i++ {
		prefix += "  "
	}

	if node == t.NIL {
		fmt.Printf("%s-NIL(%s)\n", prefix, node.color.output())
		return
	}
	if node.leftTree != nil {
		t.inorder(node.leftTree, level)
	}

	fmt.Printf("%s-%d(%s)\n", prefix, node.value, node.color.output())

	if node.rightTree != nil {
		t.inorder(node.rightTree, level)
	}
}

func (t *RBT) getSmallestChild(node *Node) *Node {
	if node.leftTree == t.NIL {
		return node
	}
	return t.getSmallestChild(node.leftTree)
}

func (t *RBT) GetMaxChild(node *Node) *Node {
	if node.rightTree == nil {
		return node
	}
	return t.GetMaxChild(node.rightTree)
}

func (t *RBT) deleteChild(node *Node, data int) bool {
	if node.value > data {
		if node.leftTree == t.NIL {
			return false
		}
		return t.deleteChild(node.leftTree, data)
	} else if node.value < data {
		if node.rightTree == t.NIL {
			return false
		}
		return t.deleteChild(node.rightTree, data)
	} else if node.value == data {
		if node.rightTree == t.NIL {
			t.deleteOneChild(node)
			return true
		}

		smallest := t.getSmallestChild(node.rightTree)
		// swap(p->value, smallest->value)
		node.value, smallest.value = smallest.value, node.value
		t.deleteOneChild(smallest)

		return true
	}
	return false
}

func (t *RBT) deleteOneChild(node *Node) {
	var child *Node
	if node.leftTree == t.NIL {
		child = node.rightTree
	} else {
		child = node.leftTree
	}

	if node.parent == nil && node.leftTree == t.NIL && node.rightTree == t.NIL {
		node = nil
		t.root = node
		return
	}

	if node.parent == nil {
		child.parent = nil
		t.root = child
		t.root.color = Black
		return
	}

	if node.parent.leftTree == node {
		node.parent.leftTree = child
	} else {
		node.parent.rightTree = child
	}
	child.parent = node.parent

	if node.color == Black {
		if child.color == Red {
			child.color = Black
		} else {
			t.deleteCase(child)
		}
	}
}

func (t *RBT) deleteCase(node *Node) {
	if node.parent == nil {
		node.color = Black
		return
	}

	if node.Sibling().color == Red {
		node.parent.color = Red
		node.Sibling().color = Black
		if node == node.parent.leftTree {
			t.rotateLeft(node.parent)
		} else {
			t.rotateRight(node.parent)
		}
	}

	if node.parent.color == Black && node.Sibling().color == Black &&
		node.Sibling().leftTree.color == Black &&
		node.Sibling().rightTree.color == Black {
		node.Sibling().color = Red
		t.deleteCase(node.parent)
	} else if node.parent.color == Red &&
		node.Sibling().color == Black &&
		node.Sibling().leftTree.color == Black &&
		node.Sibling().rightTree.color == Black {
		node.Sibling().color = Red
		node.parent.color = Black
	} else {
		if node.Sibling().color == Black {
			if node == node.parent.leftTree &&
				node.Sibling().leftTree.color == Red &&
				node.Sibling().rightTree.color == Black {
				node.Sibling().color = Red
				node.Sibling().leftTree.color = Black
				t.rotateRight(node.Sibling().leftTree)
			} else if node == node.parent.rightTree &&
				node.Sibling().leftTree.color == Black &&
				node.Sibling().rightTree.color == Red {
				node.Sibling().color = Red
				node.Sibling().rightTree.color = Black
				t.rotateLeft(node.Sibling().rightTree)
			}
		}
		node.Sibling().color = node.parent.color
		node.parent.color = Black
		if node == node.parent.leftTree {
			node.Sibling().rightTree.color = Black
			t.rotateLeft(node.Sibling())
		} else {
			node.Sibling().leftTree.color = Black
			t.rotateRight(node.Sibling())
		}
	}
}

func (t *RBT) insert(node *Node, data int) {
	if node.value >= data {
		if node.leftTree != nil && node.leftTree != t.NIL {
			t.insert(node.leftTree, data)
		} else {
			tmp := new(Node)
			tmp.value = data
			tmp.leftTree = t.NIL
			tmp.rightTree = t.NIL
			tmp.parent = node
			node.leftTree = tmp
			t.insertCase(tmp)
		}
	} else {
		if node.rightTree != nil && node.rightTree != t.NIL {
			t.insert(node.rightTree, data)
		} else {
			tmp := new(Node)
			tmp.value = data
			tmp.leftTree = t.NIL
			tmp.rightTree = t.NIL
			tmp.parent = node
			node.rightTree = tmp
			t.insertCase(tmp)
		}
	}
}

func (t *RBT) insertCase(node *Node) {
	if node.parent == nil {
		t.root = node
		node.color = Black
		return
	}
	fmt.Printf("insertCase: node: %+v, parent: %+v\n", node, node.parent)
	// t.inorder(t.root)
	if node.parent.color == Red {
		if node.Uncle() != nil && node.Uncle().color == Red {
			node.parent.color = Black
			node.Uncle().color = Black
			node.Grandparent().color = Red
			t.insertCase(node.Grandparent())
		} else {
			if node.parent.rightTree == node &&
				node.Grandparent().leftTree == node.parent {
				t.rotateLeft(node)
				node.color = Black
				node.parent.color = Red
				t.rotateRight(node)
			} else if node.parent.leftTree == node &&
				node.Grandparent().rightTree == node.parent {
				t.rotateRight(node)
				node.color = Black
				node.parent.color = Red
				t.rotateLeft(node)
			} else if node.parent.leftTree == node &&
				node.Grandparent().leftTree == node.parent {
				node.parent.color = Black
				node.Grandparent().color = Red
				t.rotateRight(node.parent)
			} else if node.parent.rightTree == node &&
				node.Grandparent().rightTree == node.parent {
				node.parent.color = Black
				node.Grandparent().color = Red
				t.rotateLeft(node.parent)
			}
		}
	}
}

func (t *RBT) DeleteTree(node *Node) {
	if node == nil || node == t.NIL {
		return
	}
	t.DeleteTree(node.leftTree)
	t.DeleteTree(node.rightTree)
}

func (t *RBT) search(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if node.value > data {
		if node.leftTree == nil {
			return nil
		}
		return t.search(node.leftTree, data)
	} else if node.value < data {
		if node.rightTree == nil {
			return nil
		}
		return t.search(node.rightTree, data)
	} else {
		return node
	}
}

func (t *RBT) Inorder() {
	if t.root == nil {
		return
	}
	t.inorder(t.root, 0)
	fmt.Println("")
}

func (t *RBT) Search(data int) (*Node, bool) {
	if t.root == nil {
		return nil, false
	}
	node := t.search(t.root, data)
	return node, node != nil
}

func (t *RBT) Insert(data int) {
	if t.root == nil {
		t.root = new(Node)
		t.root.color = Black
		t.root.leftTree = t.NIL
		t.root.rightTree = t.NIL
		t.root.value = data
	} else {
		t.insert(t.root, data)
	}
}

func (t *RBT) Delete(data int) bool {
	return t.deleteChild(t.root, data)
}

func NewRBT() *RBT {
	return &RBT{
		root: nil,
		NIL:  &Node{color: Black},
	}
}
