package singly_linked_list

type Queue[T any] struct {
	sll *SllList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{sll: &SllList[T]{}}
}

func (q *Queue[T]) Enqueue(data T) {
	newNode := &node[T]{value: data}
	sll := q.sll

	if sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	sll.tail.next = newNode
	sll.tail = newNode
}

func (q *Queue[T]) Dequeue() *T {
	sll := q.sll

	if sll.head == nil {
		return nil
	}

	retValue := sll.head.value
	sll.head = sll.head.next

	return &retValue
}

func (q *Queue[T]) IsEmpty() bool {
	return q.sll.IsEmpty()
}

func (q *Queue[T]) AsSlice() []T {
	return q.sll.AsSlice()
}

func (q *Queue[T]) ClearQueue() {
	q.sll.ClearList()
}
