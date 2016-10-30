package GO_ADT

type Node struct {
	data    struct{}
	nodePtr *Node
}

type SingleLinkedList struct {
	Head *Node
	Last *Node
}

func NewADT() *SingleLinkedList {
	SSL := &SingleLinkedList{nil, nil}
	return SSL
}

func (SSL *SingleLinkedList) ADD(data struct{}) (*Node, error) {
	if SSL.Head == nil { //new list
		newNode := &Node{data, nil}
		SSL.Head = newNode
		SSL.Last = newNode
	} else { //existing list
		nodePtr := SSL.Head
		newNode := Node{data, nodePtr}
		SSL.Head = &newNode
	}
	return SSL.Head, nil
}

func (SSL *SingleLinkedList) Remove(data struct{}) (*Node, error) {
	if SSL.Head == nil { //new list
		return nil, nil
	} else {
		return nil, nil
	}

}

func (SSL *SingleLinkedList) Find(data struct{}) error {
	if SSL.Head == nil { //new list
		return nil
	} else {
		return nil
	}
}

func (SSL *SingleLinkedList) Reverse() (*Node, error) {
	if SSL.Head == nil { //new list
		return nil, nil
	} else {
		return nil, nil
	}
}

func (SSL *SingleLinkedList) Print() error {
	if SSL.Head == nil { //new list
		return nil
	} else {
		return nil
	}
}
