package main

/**
 * My implementation of median count based on heap data structure.
 *
 * created by RinSer
 */


import (
    "fmt"
    "os"
    "bufio"
    "strconv"
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
 * Returns an empty Heap
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
    parent := (child-1)/2
    for heap.body[parent] > heap.body[child] {
        Swap(heap.body, child, parent)
        child = parent
        parent = (child-1)/2
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
    parent := (child-1)/2
    for heap.body[parent] < heap.body[child] {
        Swap(heap.body, child, parent)
        child = parent
        parent = (child-1)/2
    }
}

/**
 * Helper function to extract the mininmum value from a heap
 * Returns the minimum value deleting it from a heap
 */
func ExtractMin(heap *Heap) int {
    // Swap the first and last heap body elements
    Swap(heap.body, 0, heap.size-1)
    // Extract the min value
    minimum := heap.body[heap.size-1]
    heap.body = heap.body[:heap.size-1]
    // Decrement the heap size
    heap.size--
    // Bubble-down the new root
    if heap.size > 0 {
        root := 0
        left_child := 0
        right_child := 0
        if heap.size > 1 {
            left_child = 1
        } 
        if heap.size > 2 {
            right_child = 2
        }
        for heap.body[root] > heap.body[left_child] || heap.body[root] > heap.body[right_child] {
            // swap the values
            if heap.body[left_child] < heap.body[right_child] {
                Swap(heap.body, root, left_child)
                root = left_child
            } else if heap.body[right_child] < heap.body[left_child] {
                Swap(heap.body, root, right_child)
                root = right_child
            }
            // adjust the indices
            if heap.size > root*2+1 {
                left_child = root*2+1
            } else {
                left_child = root
            }
            if heap.size > root*2+2 {
                right_child = root*2+2
            } else {
                right_child = root
            }
        }
    }
    
    return minimum
}

/**
 * Helper function to extract the maximum value from a heap
 * Returns the maximum value deleting it from a heap
 */
func ExtractMax(heap *Heap) int {
    // Swap the first and last heap body elements
    Swap(heap.body, 0, heap.size-1)
    // Extract the max value
    maximum := heap.body[heap.size-1]
    heap.body = heap.body[:heap.size-1]
    // Decrement the heap size
    heap.size--
    // Bubble-down the new root
    if heap.size > 0 {
        root := 0
        left_child := 0
        right_child := 0
        if heap.size > 1 {
            left_child = 1
        } 
        if heap.size > 2 {
            right_child = 2
        }
        for heap.body[root] < heap.body[left_child] || heap.body[root] < heap.body[right_child] {
            // swap the values
            if heap.body[left_child] > heap.body[right_child] {
                Swap(heap.body, root, left_child)
                root = left_child
            } else if heap.body[right_child] > heap.body[left_child] {
                Swap(heap.body, root, right_child)
                root = right_child
            }
            // adjust the indices
            if heap.size > root*2+1 {
                left_child = root*2+1
            } else {
                left_child = root
            }
            if heap.size > root*2+2 {
                right_child = root*2+2
            } else {
                right_child = root
            }
        }
    }
    
    return maximum
}

/**
 * Median Maintanence algorithm
 * Returns the sum of medians
 */
func MedianMaintanence(array []int) int {
    if len(array) == 0 {
        return 0
    }
    if len(array) == 1 {
        return array[0]
    }
    // Initialize two heaps and medians' sum counter
    low_heap := NewHeap() // Max heap
    hi_heap := NewHeap()  // Min heap
    summ := 0
    first_value := array[0]
    second_value := array[1]
    if first_value < second_value {
        InsertMax(first_value, &low_heap)
        InsertMin(second_value, &hi_heap)
        summ += first_value*2
    } else {
        InsertMax(second_value, &low_heap)
        InsertMin(first_value, &hi_heap)
        summ += first_value+second_value
    }
    // The algorithm
    var next_value int
    for i := 2; i < len(array); i++ {
        next_value = array[i]
        // Insert the next value into one of the heaps
        if next_value < hi_heap.body[0] {
            InsertMax(next_value, &low_heap)
        } else {
            InsertMin(next_value, &hi_heap)
        }
        // Check if the heaps sizes do not differ more than by one
        if low_heap.size-hi_heap.size > 1 {
            extra_value := ExtractMax(&low_heap)
            InsertMin(extra_value, &hi_heap)
        } else if hi_heap.size-low_heap.size > 1 {
            extra_value := ExtractMin(&hi_heap)
            InsertMax(extra_value, &low_heap)
        }
        // Find the median index
        var median_index int
        if (i+1) % 2 == 0 {
            median_index = (i+1)/2
        } else {
            median_index = (i+2)/2
        }
        // Find the median value and add it to the medians' sum
        var median int
        if median_index == low_heap.size {
            median = low_heap.body[0]
        } else {
            median = hi_heap.body[0]
        }
        summ += median
        //fmt.Println("M", i+1)
        //fmt.Println(median, summ)
    }

    return summ
}


func main() {
    // Open the array file
    file_path := "Median.txt"
    file_pointer, err := os.Open(file_path)
    if err != nil {
        panic(err)
    }
    // Initialize the array and populate it
    medianArray := make([]int, 0)
    buffer := bufio.NewReader(file_pointer)
    line, err := buffer.ReadString('\n')
    for err == nil {
        value, er := strconv.Atoi(line[:len(line)-2])
        if er != nil {
            panic(er)
        }
        medianArray = append(medianArray, value)
        line, err = buffer.ReadString('\n')
    }
    file_pointer.Close()
    fmt.Println(MedianMaintanence(medianArray)%10000)
    // Test cases
    /*testArray := []int{5, 9, 4, 8, 2, 10, 1, 3, 6, 7, 12, 11}
    minHeap := NewHeap()
    maxHeap := NewHeap()
    for i := range(testArray) {
        InsertMin(testArray[i], &minHeap)
    }
    fmt.Println(minHeap.body)
    fmt.Println(minHeap.size)
    for minHeap.size > 0 {
        fmt.Println("Value: ", ExtractMin(&minHeap))
    }
    for j := range(testArray) {
        InsertMax(testArray[j], &maxHeap)
    }
    fmt.Println(maxHeap.body)
    fmt.Println(maxHeap.size)
    for maxHeap.size > 0 {
        fmt.Println("Value: ", ExtractMax(&maxHeap))
    }
    fmt.Println(MedianMaintanence(testArray))*/
}
