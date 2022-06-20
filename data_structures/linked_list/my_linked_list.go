package linked_list

type MyLinkedListNode struct {
	val  int
	next *MyLinkedListNode
	prev *MyLinkedListNode
}

type MyLinkedList struct {
	head *MyLinkedListNode
}

func New() *MyLinkedList {
	return &MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	node, _ := this.getNodeAndPrev(index)
	if node == nil {
		return -1
	}

	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	afterHead := this.head
	this.head = &MyLinkedListNode{
		val:  val,
		next: afterHead,
	}
	if afterHead != nil {
		afterHead.prev = this.head
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}

	current := this.head

	for current.next != nil {
		current = current.next
	}

	current.next = &MyLinkedListNode{
		val:  val,
		prev: current,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	next, prev := this.getNodeAndPrev(index)
	if next == nil && prev == nil {
		return
	}
	current := &MyLinkedListNode{
		val:  val,
		next: next,
		prev: prev,
	}
	if next != nil {
		next.prev = current
	}
	if prev != nil {
		prev.next = current
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	current, prev := this.getNodeAndPrev(index)
	if current == nil {
		return
	}
	next := current.next
	if index == 0 {
		this.head = next
	}

	if next != nil {
		next.prev = prev
	}

	if prev != nil {
		prev.next = next
	}
}

func (this *MyLinkedList) getNodeAndPrev(index int) (*MyLinkedListNode, *MyLinkedListNode) {
	current := this.head
	var prev *MyLinkedListNode = nil

	for i := 0; i < index; i++ {
		prev = current
		if current != nil {
			current = current.next
		}
	}

	return current, prev
}
