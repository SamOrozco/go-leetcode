package main

import (
	"errors"
	"math/rand"
)

func main() {
	hp := NewHeap(10)
	for i := 0; i < 100; i++ {
		hp.add(rand.Int() % 25)
	}

	for i := 0; i < 100; i++ {
		print(hp.poll())
	}
}

type Heap struct {
	capacity int
	size     int
	values   []int
}

func NewHeap(initCapacity int) *Heap {
	return &Heap{
		capacity: initCapacity,
		values:   make([]int, initCapacity),
	}
}

func (hp *Heap) min() int {
	if hp.empty() {
		hp.emptyHeap()
	}
	return hp.values[0]
}

func (hp *Heap) heapifyUp() {
	idx := hp.size - 1
	for hp.hasParent(idx) && hp.parent(idx) > hp.values[idx] {
		hp.swap(idx, hp.getParentIndex(idx))
		idx = hp.getParentIndex(idx)
	}
}

func (hp *Heap) heapifyDown() {
	idx := 0
	for hp.hasLeftIndex(idx) {
		smallerIndex := hp.getLeftIndex(idx)
		if hp.hasRightIndex(idx) && hp.right(idx) < hp.leftChild(idx) {
			smallerIndex = hp.getRightIndex(idx)
		}

		if hp.values[idx] < hp.values[smallerIndex] {
			break
		} else {
			hp.swap(idx, smallerIndex)
		}
		idx = smallerIndex
	}
}

func (hp *Heap) getLeftIndex(idx int) int {
	return (2 * idx) + 1
}

func (hp *Heap) getRightIndex(idx int) int {
	return (2 * idx) + 2
}

func (hp *Heap) getParentIndex(idx int) int {
	return (idx - 1) / 2
}

func (hp *Heap) hasLeftIndex(idx int) bool {
	return hp.getLeftIndex(idx) < hp.size
}

func (hp *Heap) hasRightIndex(idx int) bool {
	return hp.getRightIndex(idx) < hp.size
}

func (hp *Heap) hasParent(idx int) bool {
	return hp.getParentIndex(idx) >= 0
}

func (hp *Heap) leftChild(idx int) int {
	return hp.values[hp.getLeftIndex(idx)]
}

func (hp *Heap) right(idx int) int {
	return hp.values[hp.getRightIndex(idx)]
}

func (hp *Heap) parent(idx int) int {
	return hp.values[hp.getParentIndex(idx)]
}

func (hp *Heap) swap(idx1, idx2 int) {
	temp := hp.values[idx1]
	hp.values[idx1] = hp.values[idx2]
	hp.values[idx2] = temp
}

func (hp *Heap) add(val int) {
	hp.ensureCapacity()
	hp.values[hp.size] = val
	hp.size++
	hp.heapifyUp()
}

func (hp *Heap) peek() int {
	if hp.empty() {
		hp.emptyHeap()
	}
	return hp.min()
}

func (hp *Heap) poll() int {
	if hp.empty() {
		hp.emptyHeap()
	}
	val := hp.min()
	hp.swap(0, hp.size-1)
	hp.size--
	hp.heapifyDown()
	return val
}

func (hp *Heap) ensureCapacity() {
	if hp.size == hp.capacity-1 {
		len := len(hp.values) * 2
		swapArray := make([]int, len)
		copy(swapArray, hp.values)
		hp.values = swapArray
		hp.capacity = len
	}
}

func (hp *Heap) empty() bool {
	return hp.size == 0
}

func (hp *Heap) emptyHeap() {
	panic(errors.New("empty heap"))
}
