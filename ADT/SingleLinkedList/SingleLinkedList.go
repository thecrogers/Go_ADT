package SingleLinkedList

type Node struct {
	Id      int
	Data    interface{}
	NodePtr *Node
}

type SingleLinkedList struct {
	Head *Node
	Last *Node
}

func NewADT() *SingleLinkedList {
	SLL := &SingleLinkedList{nil, nil}
	return SLL
}

func (SLL *SingleLinkedList) AddStart(id int, data interface{}) {
	if SLL.Head == nil { //new list
		newNode := &Node{id, data, nil}
		SLL.Head = newNode
		SLL.Last = newNode
	} else { //existing list
		nodePtr := SLL.Head
		newNode := Node{id, data, nodePtr}
		SLL.Head = &newNode
	}
}

func (SLL *SingleLinkedList) AddEnd(id int, data interface{}) {
	if SLL.Last == nil { //new list
		newNode := &Node{id, data, nil}
		SLL.Head = newNode
		SLL.Last = newNode
	} else { //existing list
		nodePtr := SLL.Last
		newNode := &Node{id, data, nil}
		nodePtr.NodePtr = newNode
		SLL.Last = newNode
	}
}

func (SLL *SingleLinkedList) AddAfter(findId, id int, data interface{}) *Node {
	if SLL.Head == nil { //new list
		newNode := &Node{id, data, nil}
		SLL.Head = newNode
		SLL.Last = newNode
	} else { //existing list
		result := SLL.findNode(findId, SLL.Head)
		if result.NodePtr == nil {
			newNode := &Node{id, data, nil}
			result.NodePtr = newNode
			SLL.Last = newNode
		} else {
			next := result.NodePtr
			newNode := &Node{id, data, next}
			result.NodePtr = newNode
		}
	}
	return SLL.Head
}

func (SLL *SingleLinkedList) Remove(id int) bool {
	if SLL.Head == nil { //new list
		return false
	} else {
		result := SLL.findNodeForRemoval(id, SLL.Head, nil) //returns the node previout to the one needed to be removed.
		if result == nil {
			return false
		} else {
			toRemove := result.NodePtr
			if toRemove.NodePtr == nil { //ie end of list
				result.NodePtr = nil
				SLL.Last = result
				return true
			} else {
				if result.Id == id { //returned node = id to be removed ie HEAD
					SLL.Head = SLL.Head.NodePtr // move up Head
					return true
				} else {
					result.NodePtr = toRemove.NodePtr
					return true
				}
			}
		}
	}
}

func (SLL *SingleLinkedList) Find(id int) *Node {
	if SLL.Head == nil { //new list
		return nil
	} else {
		result := SLL.findNode(id, SLL.Head)
		return result
	}
}

func (SLL *SingleLinkedList) Size() int {
	size := 0
	curr := SLL.Head
	for curr != nil {
		size += 1
		curr = curr.NodePtr
	}
	return size
}

func (SLL *SingleLinkedList) Reverse() {
	if SLL.Size() == 2 { //deal with list size of two
		start := SLL.Head
		last := SLL.Last
		start.NodePtr = nil
		last.NodePtr = start
		SLL.Head = last
		SLL.Last = start
	} else {
		curr := SLL.Head
		SLL.Last = curr
		next := curr.NodePtr
		rList := next.NodePtr
		curr.NodePtr = nil
		for rList != nil {
			next.NodePtr = curr
			curr = next
			next = rList
			rList = rList.NodePtr
		}
		next.NodePtr = curr
		SLL.Head = next
	}
}

func (SLL *SingleLinkedList) Print() []*Node { //for testing only
	if SLL.Head == nil { //new list
		return nil
	} else {
		var list []*Node
		curr := SLL.Head
		for curr != nil {
			list = append(list, curr)
			curr = curr.NodePtr
		}
		return list
	}
}
