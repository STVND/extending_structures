package doubly_linked_list

type Deque[T any] struct {
	dll *DllList[T]
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{dll: &DllList[T]{}}
}

// adds an item to the back of the Deque
func (q *Deque[T]) Append(data T) {
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

// adds an item to the front of the Deque
func (q *Deque[T]) Prepend(data T) {
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

// removes/returns items from the front
func (q *Deque[T]) PopFront() *T {
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

// removes/returns items from the back
func (q *Deque[T]) PopBack() *T {

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

func (q *Deque[T]) IsEmpty() bool {
	return q.dll.IsEmpty()
}

func (q *Deque[T]) AsSlice() []T {
	return q.dll.AsSlice()
}

func (q *Deque[T]) ClearQueue() {
	q.dll.ClearList()
}
