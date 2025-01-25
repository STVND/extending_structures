package singly_linked_list

type node[T any] struct {
	value T
	next  *node[T]
}

type SllList[T any] struct {
	head, tail *node[T]
}

func (sll *SllList[T]) IsEmpty() bool {
	return sll.head == nil
}

func (sll *SllList[T]) ClearList() {
	sll.head = nil
	sll.tail = nil
}

func (sll *SllList[T]) AsSlice() []T {
	if sll.head == nil {
		return nil
	}

	var returnSlice []T

	currentNode := sll.head

	for currentNode != nil {
		returnSlice = append(returnSlice, currentNode.value)
		currentNode = currentNode.next
	}

	return returnSlice
}
