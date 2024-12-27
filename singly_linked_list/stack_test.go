package singly_linked_list

import (
	"slices"
	"strconv"
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("New Stack Properties", func(t *testing.T) {
		stack := NewStack[int]()

		if !stack.IsEmpty() {
			t.Error("New Stack should be empty")
		}

		stack.Push(1)

		if stack.IsEmpty() {
			t.Error("Stack has entry added - should not be empty")
		}

		if stack.sll.head.value != 1 {
			t.Error("Stack not pointing to correct value in head")
		}

		if stack.sll.tail.value != 1 {
			t.Error("Stack not pointing to correct value in tail")
		}

		stack.Push(2)
		stack.Push(3)

		if stack.sll.head.value == 1 {
			t.Error("Stack head value changed after new entries added")
		}

		stack.Pop()

		if stack.sll.head.value == 1 {
			t.Error("Stack Destack function failed")
		}

		if stack.sll.head.value == stack.sll.tail.value {
			t.Error("Stack Head and Tail matching after entries added and removed")
		}

	})

}

func TestPushPop(t *testing.T) {
	t.Run("Push and Pop operations", func(t *testing.T) {
		testSlice := []int{0, 10, 20, 100, 95, 13}
		stack := NewStack[int]()

		for _, testInt := range testSlice {
			stack.Push(testInt)
		}

		if stack.sll.head.value != 13 {
			t.Error("Stack created wrong head after Push")
		}

		if stack.sll.tail.value != 0 {
			t.Error("Stack created wrong tail after Pop")
		}

		stack.Pop()

		if stack.sll.head.value != 95 {
			t.Error("Stack created wrong head after Pop")
		}

		stack.Pop()
		stack.Pop()
		stack.Pop()

		if stack.sll.head.value != 10 {
			t.Error("Stack created wrong head after multiple Pop")
		}

		if stack.sll.tail.value != 0 {
			t.Error("Stack created wrong tail after multiple Pop")
		}

	})

	t.Run("Pop empty Stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Pop()
	})
}

func TestStackAsArray(t *testing.T) {
	t.Run("Conversion to array", func(t *testing.T) {
		testSlice := []int{10, 20, 25, 15}
		stack := NewStack[int]()

		for _, testInt := range testSlice {
			stack.Push(testInt)
		}

		stackSlice := stack.AsSlice()
		slices.Reverse(testSlice)

		for i := 0; i < len(testSlice); i++ {
			if stackSlice[i] != testSlice[i] {
				t.Error("Stack slice did not return expected value")
			}
		}

	})
}

func TestClearStack(t *testing.T) {
	t.Run("ClearingStack", func(t *testing.T) {
		stack := NewStack[int]()
		n := 10

		for i := 0; i < n; i++ {
			stack.Push(i)
		}

		if stack.sll.tail.value != 0 {
			t.Error("Did not add entries properly")
		}

		stack.ClearStack()

		if stack.sll.head != nil {
			t.Error("clearStack did not properly unassign values")
		}

	})
}

func BenchmarkStackAddAndRemove(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		stack := NewStack[int]()

		b.Run("Stack Capacity:"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				stack.Push(i)
			}

			for i := 0; i < size; i++ {
				stack.Pop()
			}

		})
	}
}

func BenchmarkStackClear(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		stack := NewStack[int]()

		b.Run("Stack Capacity: "+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				stack.Push(i)
			}

			stack.ClearStack()

		})

	}
}
