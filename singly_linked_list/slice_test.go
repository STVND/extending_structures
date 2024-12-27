package singly_linked_list

import (
	"strconv"
	"testing"
)

func BenchmarkAppendAndPop(b *testing.B) {

	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		slice := []int{}
		b.Run("Append and Remove[Pop]: "+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				slice = append(slice, i)
			}

			for i := 0; i < size; i++ {
				slice = slice[:len(slice)-1]
			}

		})

	}

}

func BenchmarkAppendAndDeque(b *testing.B) {

	sizes := []int{1, 10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		slice := []int{}
		b.Run("Append and Remove[Deque]: "+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < size; i++ {
				slice = append([]int{i}, slice...)
			}

			for i := 0; i < size; i++ {
				slice = slice[1:]
			}

		})

	}

}
