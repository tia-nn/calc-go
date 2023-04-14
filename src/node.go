package main

type Node interface {
	calc() int
	getPosition() Position
}

// AddNode

type AddNode struct {
	lhs      Node
	rhs      Node
	position Position
	addType  int
}

func (node AddNode) calc() int {
	if node.addType == int('+') {
		return node.lhs.calc() + node.rhs.calc()
	} else if node.addType == int('-') {
		return node.lhs.calc() - node.rhs.calc()
	} else {
		panic("unreachable")
	}
}

func (node AddNode) getPosition() Position {
	return node.position
}

// MulNode

type MulNode struct {
	lhs      Node
	rhs      Node
	position Position
	mulType  int
}

func (node MulNode) calc() int {
	if node.mulType == int('*') {
		return node.lhs.calc() * node.rhs.calc()
	} else if node.mulType == int('/') {
		return node.lhs.calc() / node.rhs.calc()
	} else {
		panic("unreachable")
	}
}

func (node MulNode) getPosition() Position {
	return node.position
}

// NumNode

type NumNode struct {
	value    int
	position Position
}

func (node NumNode) calc() int {
	return node.value
}

func (node NumNode) getPosition() Position {
	return node.position
}

var nilNode Node = NumNode{
	value:    0,
	position: nilPosition,
}
