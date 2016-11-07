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
	sll := &SingleLinkedList{nil, nil}
	return sll
}

func (sll *SingleLinkedList) AddStart(id int, data interface{}) {
	if sll.Head == nil { //new list
		newNode := &Node{id, data, nil}
		sll.Head = newNode
		sll.Last = newNode
	} else { //existing list
		nodePtr := sll.Head
		newNode := Node{id, data, nodePtr}
		sll.Head = &newNode
	}
}

func (sll *SingleLinkedList) AddEnd(id int, data interface{}) {
	if sll.Last == nil { //new list
		newNode := &Node{id, data, nil}
		sll.Head = newNode
		sll.Last = newNode
	} else { //existing list
		nodePtr := sll.Last
		newNode := &Node{id, data, nil}
		nodePtr.NodePtr = newNode
		sll.Last = newNode
	}
}

func (sll *SingleLinkedList) AddAfter(findId, id int, data interface{}) *Node {
	if sll.Head == nil { //new list
		newNode := &Node{id, data, nil}
		sll.Head = newNode
		sll.Last = newNode
	} else { //existing list
		result := sll.findNode(findId, sll.Head)
		if result.NodePtr == nil {
			newNode := &Node{id, data, nil}
			result.NodePtr = newNode
			sll.Last = newNode
		} else {
			next := result.NodePtr
			newNode := &Node{id, data, next}
			result.NodePtr = newNode
		}
	}
	return sll.Head
}

func (sll *SingleLinkedList) Remove(id int) bool {
	if sll.Head == nil { //new list
		return false
	} else {
		result := sll.findNodeForRemoval(id, sll.Head, nil) //returns the node previout to the one needed to be removed.
		if result == nil {
			return false
		} else {
			toRemove := result.NodePtr
			if toRemove.NodePtr == nil { //ie end of list
				result.NodePtr = nil
				sll.Last = result
				return true
			} else {
				if result.Id == id { //returned node = id to be removed ie HEAD
					sll.Head = sll.Head.NodePtr // move up Head
					return true
				} else {
					result.NodePtr = toRemove.NodePtr
					return true
				}
			}
		}
	}
}

func (sll *SingleLinkedList) Find(id int) *Node {
	if sll.Head == nil { //new list
		return nil
	} else {
		result := sll.findNode(id, sll.Head)
		return result
	}
}

func (sll *SingleLinkedList) Size() int {
	size := 0
	curr := sll.Head
	for curr != nil {
		size += 1
		curr = curr.NodePtr
	}
	return size
}

func (sll *SingleLinkedList) Reverse() {
	if sll.Size() == 2 { //deal with list size of two
		start := sll.Head
		last := sll.Last
		start.NodePtr = nil
		last.NodePtr = start
		sll.Head = last
		sll.Last = start
	} else {
		curr := sll.Head
		sll.Last = curr
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
		sll.Head = next
	}
}

func (sll *SingleLinkedList) Print() []*Node { //for testing only
	if sll.Head == nil { //new list
		return nil
	} else {
		var list []*Node
		curr := sll.Head
		for curr != nil {
			list = append(list, curr)
			curr = curr.NodePtr
		}
		return list
	}
}
