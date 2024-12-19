package singly_linked_list

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	head, tail *node[T]
}

func (sll *List[T]) IsEmpty() bool {
	return sll.head == nil
}

func (sll *List[T]) ClearList() {
	sll.head = nil
	sll.tail = nil
}

func (sll *List[T]) AsSlice() []T {
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
