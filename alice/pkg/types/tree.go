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

func (n *Node) SetLeftChild(child *Node) {
	n.lhc = child
}

func (n *Node) SetRightChild(child *Node) {
	n.rhc = child
}

func (n *Node) toString() string {
	if n == nil {
		return "nil"
	}
	return "(" + n.lhc.toString() + " " + strconv.FormatInt(int64(n.value), 16) + " " + n.rhc.toString() + ")"
}
