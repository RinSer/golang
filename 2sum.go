package main

/**
 * Implementation of 2-sum algorithm 
 *
 * created by RinSer
 */

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "time"
)

/**
 * Function to execute 2-sum algorithm
 * that checks the existence of each sum value from the given range
 * for any of two distinct elements in the given array
 *
 * Returns the number of occuring sum values as an integer
 */
func TwoSums(array []int, minSum, maxSum int) int {
    // Initialiaze and populate the array hash set    
    arrayMap := make(map[int]int)
    for i := range(array) {
        if arrayMap[array[i]] == 1 {
            arrayMap[array[i]]++
        } else {
            arrayMap[array[i]] = 1
        }
    }
    // Count the number of occuring two elements sum
    sumsCounter := 0
    for t := minSum; t < maxSum+1; t++ {
        go func(t int, sumsCounter *int) {
            for x := range(array) {
                if (arrayMap[t-array[x]] == 1 && (t-array[x]) != array[x]) || arrayMap[t-array[x]] > 1 {
                    fmt.Println(array[x], t)
                    *sumsCounter++
                    return
                }
            }
        }(t, &sumsCounter)
    }
    time.Sleep(time.Second)
    return sumsCounter
}


func main() {
    // Assignment Case
    filePath := "2sum.txt"
    filePointer, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    buffer := bufio.NewReader(filePointer)
    assignmentArray := make([]int, 0)
    line, err := buffer.ReadString('\n')
    for err == nil {
        value, er := strconv.Atoi(line[:len(line)-2])
        if er != nil {
            panic(er)
        }
        assignmentArray = append(assignmentArray, value)
        line, err = buffer.ReadString('\n')
    }
    filePointer.Close()
    // Compute the 2-sums number
    //fmt.Println(assignmentArray[0])
    fmt.Println(TwoSums(assignmentArray, -10000, 10000))
    // Test Case
    //testArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
    //fmt.Println(TwoSums(testArray, -10, 10))
}
