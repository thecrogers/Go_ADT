package SingleLinkedList

import "testing"

type TestStruct struct {
	Item string
	Id   int
}

func TestSingeLinkedList(t *testing.T) {
	dataOne := TestStruct{"hello", 1}
	dataTwo := TestStruct{"World", 2}
	dataThree := TestStruct{"Boom", 3}

	SLL := NewADT()
	if SLL == nil {
		t.Error("Failed to init new SSL")
		t.Fail()
	}

	SLL.AddStart(1, dataOne)
	if SLL.Size() != 1 {
		t.Errorf("did not add struct: %v", dataOne)
	}
	SLL.AddStart(2, dataTwo)
	if SLL.Size() != 2 {
		t.Errorf("did not add struct: %v", dataTwo)
	}
	SLL.AddStart(3, dataThree)
	if SLL.Size() != 3 {
		t.Errorf("did not add struct: %v", dataThree)
	}

	t.Logf("size of List: %d", SLL.Size())

	printResult := SLL.Print()
	if len(printResult) != 3 {
		t.Error("\nFailed to print slice is of size 0")
		t.Fail()
	} else {
		t.Log("list of nodes head to tail traversal")
		for _, result := range printResult {
			t.Logf("\n print result: \n %v", result)
		}
	}

	findResult := SLL.Find(2)
	if findResult == nil {
		t.Error("Failed to find node with id 2")
		t.Fail()
	} else {
		t.Logf("\n node found result: \n %v", findResult)
	}

	SLL.Reverse()
	reverseResult := SLL.Print()
	if len(reverseResult) != 3 {
		t.Error("A node was lost while reversing")
		t.Fail()
	} else {
		count := 1
		for _, node := range reverseResult {
			if node.Id != count {
				t.Error("reverse is out of order")
				t.Fail()
			}
			t.Logf("\n print reverse result: \n %v", node)
			count++
		}
	}

	if SLL.Last.Id != 3 {
		t.Error("reverse did not keep SLL.Last in sync")
		t.Fail()
	}

	removeResultTwo := SLL.Remove(2)
	if !removeResultTwo {
		t.Error("Failed to remove node with id 2")
		t.Fail()
	}
	t.Logf("size of list prior to second reversal: %d", SLL.Size())
	SLL.Reverse()
	reverseResult = SLL.Print()
	if len(reverseResult) != 2 {
		t.Error("A node was lost while reversing")
		t.Fail()
	} else {
		count := 3
		for _, node := range reverseResult {
			if node.Id != count {
				t.Error("reverse is out of order")
				t.Fail()
			}
			t.Logf("\n print reverse result: \n %v", node)
			count -= 2
		}
	}

	if SLL.Last.Id != 1 {
		t.Error("reverse did not keep SLL.Last in sync")
		t.Fail()
	}

	t.Logf("size of List: %d", SLL.Size())

	printResult = SLL.Print()
	for _, result := range printResult {
		t.Logf("\n print result: \n %v", result)
	}

	removeResultOne := SLL.Remove(1)
	if !removeResultOne {
		t.Error("Failed to remove node with id 1")
		t.Fail()
	}

	t.Logf("size of List after Removal of id 1: %d", SLL.Size())

	SLL.AddEnd(4, dataOne)
	reverseResult = SLL.Print()
	if len(reverseResult) != 2 {
		t.Error("A node was lost while reversing")
		t.Fail()
	} else {
		count := 3
		for _, node := range reverseResult {
			if node.Id != count {
				t.Error("reverse is out of order")
				t.Fail()
			}
			t.Logf("\n print reverse result: \n %v", node)
			count++
		}
	}
	t.Logf("size of List: %d", SLL.Size())
	SLL.AddEnd(6, dataTwo)
	t.Logf("size of List: %d", SLL.Size())
	SLL.AddAfter(4, 5, dataThree)
	t.Logf("size of List: %d", SLL.Size())
	reverseResult = SLL.Print()
	if len(reverseResult) != 4 {
		t.Error("A node was lost")
		t.Fail()
	} else {
		count := 3
		for _, node := range reverseResult {
			if node.Id != count {
				t.Error("reverse is out of order")
				t.Fail()
			}
			t.Logf("\n print reverse result: \n %v", node)
			count++
		}
	}

}
