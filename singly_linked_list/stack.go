package singly_linked_list

type Stack[T any] struct {
	sll *List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{sll: &List[T]{}}
}

func (s *Stack[T]) Push(data T) {
	newNode := &node[T]{value: data}
	sll := s.sll

	if s.sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	newNode.next = sll.head
	sll.head = newNode
}

func (s *Stack[T]) Pop() *T {
	sll := s.sll

	if sll.head == nil {
		return nil
	}

	retValue := sll.head.value
	sll.head = sll.head.next

	return &retValue
}

func (s *Stack[T]) IsEmpty() bool {
	return s.sll.IsEmpty()
}

func (s *Stack[T]) AsSlice() []T {
	return s.sll.AsSlice()
}

func (s *Stack[T]) ClearStack() {
	s.sll.ClearList()
}
