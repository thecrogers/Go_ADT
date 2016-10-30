package GO_ADT

import "github.com/thecrogers/Go_ADT/ADT/SingleLinkedList"

type TestStruct struct {
	Item string
	Id   int
}

func main() {

	// adt := ADT{}
	SLLI := SingleLinkedList.NewADT()
	if err != nil {
		// do something
	}
	dataOne := TestStruct{"hello", 1}
	dataTwo := TestStruct{"World", 2}
	ListOne, err := SLLI.Add(dataOne)
	ListOne, err = SLLI.Add(dataTwo)
	SLL.Print()
	result := SSL.Find(2)

	//	fmt.Sprintf("Result of SSL fine is:\n ID: %d", result.data)
}
