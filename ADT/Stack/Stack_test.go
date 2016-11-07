package Stack

import "testing"

type TestStruct struct {
	Item string
	Id   int
}

func TestStack(t *testing.T) {
	dataOne := TestStruct{"hello", 1}
	dataTwo := TestStruct{"World", 2}
	dataThree := TestStruct{"Boom", 3}

	stack := NewADT()
	if stack == nil {
		t.Error("Failed to init new stack")
		t.Fail()
	}

	stack.Push(dataOne)
	if stack.Size() != 1 {
		t.Errorf("did not add struct: %v", dataOne)
	}
	stack.Push(dataTwo)
	if stack.Size() != 2 {
		t.Errorf("did not add struct: %v", dataTwo)
	}
	stack.Push(dataThree)
	if stack.Size() != 3 {
		t.Errorf("did not add struct: %v", dataThree)
	}

	var popOne TestStruct
	popOne = stack.Pop().(TestStruct)
	if stack.Size() != 2 {
		t.Errorf("did not pop stack")
	}
	if popOne.Id != 3 {
		t.Errorf("poped the wrong node: %d", popOne.Id)
	}
}
