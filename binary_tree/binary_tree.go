package binarytree

import (
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	value  T
	parent *node[T]
	left   *node[T]
	right  *node[T]
}

type Tree[T constraints.Ordered] struct {
	head  *node[T]
	count int
}

func NewTree[T constraints.Ordered]() *Tree[T] {
	return &Tree[T]{head: &node[T]{}, count: 0}
}

func (tree *Tree[T]) Insert(data T) {
	newNode := &node[T]{value: data}
	tree.count += 1

	if tree.isEmpty() {
		tree.head = newNode
		return
	}

	currentNode := tree.head

	for {
		if newNode.value <= currentNode.value {
			if currentNode.left == nil {
				newNode.parent = currentNode
				currentNode.left = newNode
				break

			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				newNode.parent = currentNode
				currentNode.right = newNode
				break

			} else {
				currentNode = currentNode.right
			}
		}
	}

	if tree.count > 1 && tree.count%2 == 1 {
		tree.Balance()
	}
}

func (tree *Tree[T]) Contains(value T) bool {
	currentNode := tree.head

	for {
		if currentNode == nil {
			break
		}
		if currentNode.value == value {
			return true
		}
		if value <= currentNode.value {
			if currentNode.left == nil {
				break
			}
			currentNode = currentNode.left
		}
		if value > currentNode.value {
			if currentNode.right != nil {
				break
			}
			currentNode = currentNode.right
		}
	}

	return false
}

// in-order traversal
func (tree *Tree[T]) GetValues() []T {
	return tree.inOrder(tree.head)
}

func (tree *Tree[T]) inOrder(currentNode *node[T]) []T {
	result := []T{}
	if currentNode == nil {
		return result
	}

	result = append(result, tree.inOrder(currentNode.left)...)
	result = append(result, currentNode.value)
	result = append(result, tree.inOrder(currentNode.right)...)

	return result
}

// naive balance - uses something like quicksort to add items to the tree
// not expecting this to be performant since it can't do local changes
// this rebuilds the whole tree when just moving some nodes could be all that is needed
// need to read up on node rotation and proper tree balancing
func (tree *Tree[T]) Balance() {
	values := tree.GetValues()
	tree.ClearTree()
	tree.balance(0, len(values)-1, values)

}

func (tree *Tree[T]) balance(low, high int, values []T) {
	if low <= high {
		mid := low + (high-low)/2
		tree.Insert(values[mid])
		tree.balance(low, mid-1, values)
		tree.balance(mid+1, high, values)

	}

}

func (tree *Tree[T]) Delete(value T) {
	if tree.Contains(value) {
		values := tree.GetValues()
		tree.ClearTree()

		for ind, val := range values {
			if val == value {
				values = append(values[:ind], values[ind+1:]...)
				break
			}
		}

		tree.balance(0, len(values)-1, values)
	}

}

func (tree *Tree[T]) ClearTree() {
	tree.head = nil
}

func (tree *Tree[T]) isEmpty() bool {
	return tree.head == nil
}
