package doubly_linked_list

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type DllList[T any] struct {
	head, tail *node[T]
	length     int
}

func (dll *DllList[T]) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DllList[T]) ClearList() {
	dll.head = nil
	dll.tail = nil
}

func (dll *DllList[T]) AsSlice() []T {
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
