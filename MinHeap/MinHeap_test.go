package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func isMinHeapPropertySatisfied(heap MinHeap) bool {
	for i := 1; i < len(heap); i++ {
		parentIdx := (i - 1) / 2
		if parentIdx < 0 {
			return true
		}

		if heap[parentIdx] > heap[i] {
			return false
		}
	}
	return true
}

func TestMinHeap(t *testing.T) {
	var minHeap = NewMinHeap([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	minHeap.Insert(76)
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	require.Equal(t, -5, minHeap.Peek())
	require.Equal(t, -5, minHeap.Remove())
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	require.Equal(t, 2, minHeap.Peek())
	require.Equal(t, 2, minHeap.Remove())
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	require.Equal(t, 6, minHeap.Peek())
	minHeap.Insert(87)
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
}
