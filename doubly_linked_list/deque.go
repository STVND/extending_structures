package doubly_linked_list

type Queue[T any] struct {
	dll *List[T]
}

func NewDeque[T any]() *Queue[T] {
	return &Queue[T]{dll: &List[T]{}}
}

func (q *Queue[T]) Append(data T) {
	newNode := &node[T]{value: data}
	dll := q.dll

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
	dll.length += 1
}

func (q *Queue[T]) Prepend(data T) {
	newNode := &node[T]{value: data}
	dll := q.dll

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.head.prev = newNode
	newNode.next = dll.head
	dll.head = newNode
	dll.length += 1

}

func (q *Queue[T]) PopFront() *T {
	dll := q.dll

	if dll.head == nil {
		return nil
	}

	retValue := dll.head.value
	dll.head = dll.head.next
	dll.head.prev = nil
	dll.length -= 1

	//this is the easiest implementation without changing logic higher in the function
	if dll.length == 0 {
		dll.ClearList()
	}

	return &retValue
}

func (q *Queue[T]) PopBack() *T {

	dll := q.dll

	if dll.tail == nil {
		return nil
	}

	retValue := dll.tail.value
	dll.tail = dll.tail.prev
	dll.length -= 1

	//this is the easiest implementation without changing logic higher in the function
	if dll.length == 0 {
		dll.ClearList()
	}

	return &retValue
}

func (q *Queue[T]) IsEmpty() bool {
	return q.dll.IsEmpty()
}

func (q *Queue[T]) AsSlice() []T {
	return q.dll.AsSlice()
}

func (q *Queue[T]) ClearQueue() {
	q.dll.ClearList()
}
