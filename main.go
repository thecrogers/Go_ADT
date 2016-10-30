package GO_ADT

type Node struct {
	data    struct{}
	nodePtr *Node
}

type SingleLinkedListInterface interface {
	Add(struct{}) *Node
	Remove(struct{}) *Node
	Find(struct{}) struct{}
	Reverse() *Node
	Print()
}

type SingleLinkedList struct {
	Head *Node
	Last *Node
}

type ADT struct {
	Head *Node
	Last *Node
}

func (Adt *ADT) SingleLinkedList() (SingleLinkedListInterface, error) {
	var SingleLinkedList SingleLinkedListInterface
	SingleLinkedList = &SingleLinkedList{nil, nil}
	return SingleLinkedList, nil
}
