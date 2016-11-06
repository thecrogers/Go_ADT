package SingleLinkedList

//recursive Find
func (SSL *SingleLinkedList) findNode(id int, node *Node) *Node {
	if node.Id == id {
		return node
	} else if node.NodePtr == nil {
		return nil
	} else {
		return SSL.findNode(id, node.NodePtr)
	}
}

func (SSL *SingleLinkedList) findNodeForRemoval(id int, node *Node, last *Node) *Node {
	if node.Id == id {
		if last == nil { //only one node in list
			return node
		} else {
			return last
		}
	} else if node.NodePtr == nil { //make sure next node is not nil
		return nil
	} else {
		return SSL.findNodeForRemoval(id, node.NodePtr, node)
	}
}
