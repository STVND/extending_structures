package singly_linked_list

import (
	"strconv"
	"testing"
)

func TestNewQueue(t *testing.T) {
	t.Run("New Queue Properties", func(t *testing.T) {
		queue := NewQueue[int]()

		if !queue.IsEmpty() {
			t.Error("New Queue should be empty")
		}

		queue.Enqueue(1)

		if queue.IsEmpty() {
			t.Error("Queue has entry added - should not be empty")
		}

		if queue.sll.head.value != 1 {
			t.Error("Queue not pointing to correct value in head")
		}

		if queue.sll.tail.value != 1 {
			t.Error("Queue not pointing to correct value in tail")
		}

		queue.Enqueue(2)
		queue.Enqueue(3)

		if queue.sll.head.value != 1 {
			t.Error("Queue head value changed after new entries added")
		}

		queue.Dequeue()

		if queue.sll.head.value == 1 {
			t.Error("Queue Dequeue function failed")
		}

		if queue.sll.head.value == queue.sll.tail.value {
			t.Error("Queue Head and Tail matching after entries added and removed")
		}

	})

}

func TestEnqueueAndDequeue(t *testing.T) {
	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		testSlice := []int{0, 1, 2, 3}
		queue := NewQueue[int]()

		for _, testInt := range testSlice {
			queue.Enqueue(testInt)
		}

		if queue.sll.head.value != testSlice[0] {
			t.Error("Returned wrong head after Enqueue")
		}

		if queue.sll.tail.value != testSlice[len(testSlice)-1] {
			t.Error("Return wrong tail after Enqueue")
		}

		queue.Dequeue()
		queue.Dequeue()

		if queue.sll.head.value != testSlice[2] {
			t.Error("Returned wrong head after Dequeue")
		}

		if queue.sll.tail.value != testSlice[len(testSlice)-1] {
			t.Error("Return wrong tail after Dequeue")
		}

	})

	t.Run("Dequeue empty Queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Dequeue()
	})
}

func TestQueueAsArray(t *testing.T) {
	t.Run("Conversion to array", func(t *testing.T) {
		testSlice := []int{10, 20, 25, 15}
		queue := NewQueue[int]()

		for _, testInt := range testSlice {
			queue.Enqueue(testInt)
		}

		queueSlice := queue.AsSlice()

		for i := 0; i < len(testSlice); i++ {
			if queueSlice[i] != testSlice[i] {
				t.Error("Queue slice did not return expected value")
			}
		}

	})
}

func TestClearQueue(t *testing.T) {
	t.Run("ClearingQueue", func(t *testing.T) {
		queue := NewQueue[int]()
		n := 10

		for i := 0; i < n; i++ {
			queue.Enqueue(i)
		}

		if queue.sll.head.value != 0 {
			t.Error("Did not add entries properly")
		}

		queue.ClearQueue()

		if queue.sll.head != nil || queue.sll.tail != nil {
			t.Error("clearQueue did not properly unassign values")
		}

	})
}

func BenchmarkQueueAddAndRemove(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		queue := NewQueue[int]()

		b.Run("Queue Capacity: "+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				queue.Enqueue(i)
			}

			for i := 0; i < size; i++ {
				queue.Dequeue()
			}

		})
	}
}

func BenchmarkQueueClear(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		queue := NewQueue[int]()

		b.Run("Queue Capacity: "+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				queue.Enqueue(i)
			}

			queue.ClearQueue()

		})

	}
}
