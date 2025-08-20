package tree

import (
	"strconv"
)

type Node struct {
	lhc *Node
	rhc *Node
	value int
}

func NewNode(value int) *Node {
	return &Node{
		lhc:   nil,
		rhc:   nil,
		value: value,
	}
}

func T(left *Node, value int, right *Node) *Node {
	return &Node{
		lhc:   left,
		rhc:   right,
		value: value,
	}
}

func L(value int) *Node {
	return &Node{
		lhc:   nil,
		rhc:   nil,
		value: value,
	}
}

func (n *Node) SetLeftChild(child *Node) {
	n.lhc = child
}

func (n *Node) SetRightChild(child *Node) {
	n.rhc = child
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	var leftSubtree string
	if n.lhc != nil {
		leftSubtree = n.lhc.String()
	}

	var rightSubtree string
	if n.rhc != nil {
		rightSubtree = n.rhc.String()
	}

	return "{" + leftSubtree + "/" + strconv.FormatInt(int64(n.value), 16) + "\\" + rightSubtree + "}"
}
