package main

/**
 * My implementation of median count based on heap data structure.
 *
 * created by RinSer
 */


import (
    "fmt"
)


/**
 * Heap data structure
 * consists of integer slice that stores values
 * and size value
 */
type Heap struct {
    body []int
    size int
}

/**
 * Helper function to create a new Heap
 *
 * returns an empty Heap
 */
func NewHeap() Heap {
    var new_heap Heap
    new_heap.body = make([]int, 0)
    new_heap.size = 0

    return new_heap
}

/**
 * Helper function to swap two values in a given 
 * array by their indices
 */
func Swap(array []int, i int, j int) {
    tmp := array[i]
    array[i] = array[j]
    array[j] = tmp
}

/**
 * Helper function to insert a new value in a minimum heap
 */
func InsertMin(value int, heap *Heap) {
    // Append the new value to the heap end
    heap.body = append(heap.body, value)
    // Increase the heap size by one
    heap.size++
    // Bubble-up the new value if necessary
    child := heap.size-1
    parent := child/2
    for heap.body[parent] > heap.body[child] {
        Swap(heap.body, child, parent)
        child = parent
        parent = child/2
    }
}

/**
 * Helper function to insert a new value in a maximum heap
 */
func InsertMax(value int, heap *Heap) {
    // Append the new value to the heap end
    heap.body = append(heap.body, value)
    // Increase the heap size by one
    heap.size++
    // Bubble-up the new value if necessary
    child := heap.size-1
    parent := child/2
    for heap.body[parent] < heap.body[child] {
        Swap(heap.body, child, parent)
        child = parent
        parent = child/2
    }
}


func main() {
    minHeap := NewHeap()
    maxHeap := NewHeap()
    testArray := [12]int{5, 9, 4, 8, 2, 10, 1, 3, 6, 7, 12, 11}
    for i := range(testArray) {
        InsertMin(testArray[i], &minHeap)
    }
    fmt.Println(minHeap.body)
    fmt.Println(minHeap.size)
    for j := range(testArray) {
        InsertMax(testArray[j], &maxHeap)
    }
    fmt.Println(maxHeap.body)
    fmt.Println(maxHeap.size)
}
