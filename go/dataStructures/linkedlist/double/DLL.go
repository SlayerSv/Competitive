package double

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DLL struct {
	Head *Node
	Tail *Node
	size int
}

func NewDLL() *DLL {
	dll := &DLL{
		Head: &Node{},
		Tail: &Node{},
	}
	dll.Head.Next = dll.Tail
	dll.Tail.Prev = dll.Head
	return dll
}

func (dll *DLL) PushFront(node *Node) {
	dll.insert(node, dll.Head.Next)
}

func (dll *DLL) PushBack(node *Node) {
	dll.insert(node, dll.Tail)
}

func (dll *DLL) PopFront() *Node {
	if dll.size < 1 {
		return nil
	}
	return dll.delete(dll.Head.Next)
}

func (dll *DLL) PopBack() *Node {
	if dll.size < 1 {
		return nil
	}
	return dll.delete(dll.Tail.Prev)
}

func (dll *DLL) Insert(node *Node, position int) {
	next := dll.Get(position)
	dll.insert(node, next)
}

func (dll *DLL) insert(prev *Node, next *Node) {
	next.Prev.Next = prev
	prev.Prev = next.Prev
	next.Prev = prev
	prev.Next = next
	dll.size++
}

func (dll *DLL) Delete(position int) *Node {
	forDeletion := dll.Get(position)
	return dll.delete(forDeletion)
}

func (dll *DLL) delete(node *Node) *Node {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node.Next = nil
	node.Prev = nil
	dll.size--
	return node
}

func (dll *DLL) Get(position int) *Node {
	if position >= dll.size || position < 0 {
		panic("index out of range")
	}
	tmp := dll.Head.Next
	i := 0
	if position > dll.size/2 {
		tmp = dll.Tail.Prev
		i = dll.size - 1
		for i != position {
			i--
			tmp = tmp.Prev
		}
	} else {
		for i != position {
			i++
			tmp = tmp.Next
		}
	}
	return tmp
}

func (dll *DLL) Size() int {
	return dll.size
}
