package Stack

import "reflect"

type Node struct {
	Data    interface{}
	NodePtr *Node
}

type Stack struct {
	Top *Node
}

func NewADT() *Stack {
	stack := &Stack{}
	return stack
}

func (stack *Stack) Push(data interface{}) {
	newNode := &Node{data, nil}
	if stack.Top == nil {
		stack.Top = newNode
	} else {
		newNode.NodePtr = stack.Top
		stack.Top = newNode
	}

}

func (stack *Stack) Pop() interface{} {
	if stack.Top == nil {
		return nil
	} else {
		curr := stack.Top
		stack.Top = curr.NodePtr
		dataType := reflect.Type(curr.Data)
		return curr.Data.(dataType)
	}
}

func (stack *Stack) Size() int {
	curr := stack.Top
	size := 0

	for curr != nil {
		size++
		curr = curr.NodePtr
	}
	return size
}
