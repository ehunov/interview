package linked_list

import (
	"testing"
)

func TestMyLinkedList1(t *testing.T) {
	myLinkedList := New()
	assert(t, -1, myLinkedList.Get(0))
	myLinkedList.AddAtHead(1) // 1
	assert(t, 1, myLinkedList.Get(0))
	myLinkedList.AddAtHead(12) // 12 -> 1
	assert(t, 12, myLinkedList.Get(0))
	assert(t, 1, myLinkedList.Get(1))
	myLinkedList.AddAtTail(3) // 12 -> 1 -> 3
	assert(t, 3, myLinkedList.Get(2))
	myLinkedList.AddAtIndex(1, 2) // 12 -> 2 -> 1 -> 3
	assert(t, 2, myLinkedList.Get(1))
	myLinkedList.DeleteAtIndex(1) // 12 -> 1 -> 3
	assertLinkedList(t, []int{12, 1, 3}, myLinkedList)
}

func TestMyLinkedList2(t *testing.T) {
	myLinkedList := New()
	myLinkedList.AddAtHead(7)     // 7
	myLinkedList.AddAtHead(2)     // 2 -> 7
	myLinkedList.AddAtHead(1)     // 1 -> 2 -> 7
	myLinkedList.AddAtIndex(3, 0) // 1 -> 2 -> 7 -> 0
	myLinkedList.DeleteAtIndex(2) // 1 -> 2 -> 0
	myLinkedList.AddAtHead(6)     // 6 -> 1 -> 2 -> 0
	myLinkedList.AddAtTail(4)     // 6 -> 1 -> 2 -> 0 -> 4
	assert(t, 4, myLinkedList.Get(4))
	myLinkedList.AddAtHead(4)     // 4 -> 6 -> 1 -> 2 -> 0 -> 4
	myLinkedList.AddAtIndex(5, 0) // 4 -> 6 -> 1 -> 2 -> 0 -> 0 -> 4
	myLinkedList.AddAtHead(6)     // 6 -> 4 -> 6 -> 1 -> 2 -> 0 -> 0 -> 4

	assertLinkedList(t, []int{6, 4, 6, 1, 2, 0, 0, 4}, myLinkedList)
}

func TestMyLinkedList3(t *testing.T) {
	myLinkedList := New()
	myLinkedList.AddAtHead(1)     // 1
	myLinkedList.AddAtTail(3)     // 1 -> 3
	myLinkedList.AddAtIndex(1, 2) // 1 -> 2 -> 3
	assert(t, 2, myLinkedList.Get(1))
	myLinkedList.DeleteAtIndex(0) // 2 -> 3
	assertLinkedList(t, []int{2, 3}, myLinkedList)
}

func TestMyLinkedList4(t *testing.T) {
	myLinkedList := New()
	myLinkedList.AddAtIndex(0, 10) // 10
	myLinkedList.AddAtIndex(0, 20) // 20 -> 10
	myLinkedList.AddAtIndex(1, 30) // 20 -> 30 -> 10
	assertLinkedList(t, []int{20, 30, 10}, myLinkedList)
}

func TestMyLinkedList5(t *testing.T) {
	myLinkedList := New()
	myLinkedList.AddAtHead(1)     // 1
	myLinkedList.AddAtTail(3)     // 1 -> 3
	myLinkedList.AddAtIndex(1, 2) // 1 -> 2 -> 3
	assert(t, 2, myLinkedList.Get(1))
	myLinkedList.DeleteAtIndex(1) // 1 -> 3
	assert(t, 3, myLinkedList.Get(1))
	assert(t, -1, myLinkedList.Get(3))
	myLinkedList.DeleteAtIndex(3) // 1 -> 3
	myLinkedList.DeleteAtIndex(0) // 3
	assert(t, 3, myLinkedList.Get(0))
	myLinkedList.DeleteAtIndex(0) //
	assert(t, -1, myLinkedList.Get(0))
}

func assertLinkedList(t *testing.T, expected []int, actual *MyLinkedList) {
	for i, element := range expected {
		assert(t, element, actual.Get(i))
	}
}

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
