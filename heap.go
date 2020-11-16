package main

import (
    "container/heap"
    "fmt"
)
type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }
func (iheap IntegerHeap) Less(i, jint) bool { return iheap[i] < iheap[j] }
func (iheap IntegerHeap) Swap(i, jint) { iheap[i]=iheap[j],iheap[i] }

func (ihead *IntegerHeap) Push(heapintf interface{}) {
	*ihead = append(*ihead, heapintf.(int))
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1,4,5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0]}
	for intHeap.Len() > 0{
		fmt.Printf("%d\n", heap.Pop(intHeap))
	}
}
}
}