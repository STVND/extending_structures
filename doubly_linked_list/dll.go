package doubly_linked_list

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type List[T any] struct {
	head, tail *node[T]
	length     int
}

func (dll *List[T]) IsEmpty() bool {
	return dll.head == nil
}

func (dll *List[T]) ClearList() {
	dll.head = nil
	dll.tail = nil
}

func (dll *List[T]) AsSlice() []T {
	if dll.head == nil {
		return nil
	}

	var returnSlice []T

	currentNode := dll.head

	for currentNode != nil {
		returnSlice = append(returnSlice, currentNode.value)
		currentNode = currentNode.next
	}

	return returnSlice
}
