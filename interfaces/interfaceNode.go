package main

import (
	"errors"
	"fmt"
)

// ErrInvalidNode : Node is not valid
var ErrInvalidNode = errors.New("Node is not valid")

// Node : interface
type Node interface {
	SetValue(v int) error
	GetValue() int
}

// SLLNode : struct
type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

// SetValue : func(v int) error
func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

// GetValue : func() int
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// NewSLLNode : func() *SLLNode
func NewSLLNode() *SLLNode {
	return &SLLNode{
		SNodeMessage: "This is a message from the normal node",
	}
}

// PowerNode : struct
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

// SetValue : func(v int) error
func (pNode *PowerNode) SetValue(v int) error {
	if pNode == nil {
		return ErrInvalidNode
	}
	pNode.value = v * 10
	return nil
}

// GetValue : func() int
func (pNode *PowerNode) GetValue() int {
	return pNode.value
}

// NewPowerNode : func() *PowerNode
func NewPowerNode() *PowerNode {
	return &PowerNode{
		PNodeMessage: "This is a message from the power node",
	}
}

func main() {
	n := createNode(5)
	switch concreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message: ", concreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message: ", concreten.PNodeMessage)
	}

}

func createNode(v int) Node {
	pn := NewPowerNode()
	pn.SetValue(v)
	return pn
}
