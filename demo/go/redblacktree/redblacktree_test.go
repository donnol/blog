package redblacktree

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	gp := &Node{
		color: Black,
		value: 3,
	}
	uncle := &Node{
		color:  Black,
		value:  5,
		parent: gp,
	}
	pa := &Node{
		color:  Black,
		value:  2,
		parent: gp,
	}
	node := &Node{
		color:  Black,
		value:  1,
		parent: pa,
	}
	sn := &Node{
		color:  Black,
		value:  4,
		parent: pa,
	}
	gp.leftTree = pa
	gp.rightTree = uncle
	pa.leftTree = node
	pa.rightTree = sn

	if !reflect.DeepEqual(node.Grandparent(), gp) {
		t.Fatalf("Bad result: %+v != %+v\n", node.Grandparent(), gp)
	}
	if !reflect.DeepEqual(node.Uncle(), uncle) {
		t.Fatalf("Bad result: %+v != %+v\n", node.Uncle(), uncle)
	}
	if !reflect.DeepEqual(node.Sibling(), sn) {
		t.Fatalf("Bad result: %+v != %+v\n", node.Sibling(), sn)
	}
}

func TestRBTInsert(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{
			name:   "",
			values: []int{3, 4, 2, 10, 9, 20, 30, 12, 12, 11},
		},
		{
			name:   "",
			values: []int{4, 3, 2, 10, 9, 20, 30, 12, 12, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewRBT()
			for _, value := range tt.values {
				tree.Insert(value)
			}
			tree.Inorder()
		})
	}
}

func TestRBTSearch(t *testing.T) {
}

func TestRBTDelete(t *testing.T) {

}
