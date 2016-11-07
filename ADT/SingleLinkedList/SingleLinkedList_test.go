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

	sll := NewADT()
	if sll == nil {
		t.Error("Failed to init new SSL")
		t.Fail()
	}

	sll.AddStart(1, dataOne)
	if sll.Size() != 1 {
		t.Errorf("did not add struct: %v", dataOne)
	}
	sll.AddStart(2, dataTwo)
	if sll.Size() != 2 {
		t.Errorf("did not add struct: %v", dataTwo)
	}
	sll.AddStart(3, dataThree)
	if sll.Size() != 3 {
		t.Errorf("did not add struct: %v", dataThree)
	}

	t.Logf("size of List: %d", sll.Size())

	printResult := sll.Print()
	if len(printResult) != 3 {
		t.Error("\nFailed to print slice is of size 0")
		t.Fail()
	} else {
		t.Log("list of nodes head to tail traversal")
		for _, result := range printResult {
			t.Logf("\n print result: \n %v", result)
		}
	}

	findResult := sll.Find(2)
	if findResult == nil {
		t.Error("Failed to find node with id 2")
		t.Fail()
	} else {
		t.Logf("\n node found result: \n %v", findResult)
	}

	sll.Reverse()
	reverseResult := sll.Print()
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

	if sll.Last.Id != 3 {
		t.Error("reverse did not keep sll.Last in sync")
		t.Fail()
	}

	removeResultTwo := sll.Remove(2)
	if !removeResultTwo {
		t.Error("Failed to remove node with id 2")
		t.Fail()
	}
	t.Logf("size of list prior to second reversal: %d", sll.Size())
	sll.Reverse()
	reverseResult = sll.Print()
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

	if sll.Last.Id != 1 {
		t.Error("reverse did not keep sll.Last in sync")
		t.Fail()
	}

	t.Logf("size of List: %d", sll.Size())

	printResult = sll.Print()
	for _, result := range printResult {
		t.Logf("\n print result: \n %v", result)
	}

	removeResultOne := sll.Remove(1)
	if !removeResultOne {
		t.Error("Failed to remove node with id 1")
		t.Fail()
	}

	t.Logf("size of List after Removal of id 1: %d", sll.Size())

	sll.AddEnd(4, dataOne)
	reverseResult = sll.Print()
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
	t.Logf("size of List: %d", sll.Size())
	sll.AddEnd(6, dataTwo)
	t.Logf("size of List: %d", sll.Size())
	sll.AddAfter(4, 5, dataThree)
	t.Logf("size of List: %d", sll.Size())
	reverseResult = sll.Print()
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
