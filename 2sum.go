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
    "strings"
    //"sync"
    //"time"
)

/**
 * Function to execute 2-sum algorithm
 * that checks the existence of each sum value from the given range
 * for any of two distinct elements in the given array
 *
 * Returns the number of occuring sum values as an integer
 */
func TwoSums(array []int64, minSum, maxSum int) int {
    // Initialiaze and populate the array hash set    
    arrayMap := make(map[int64]int)
    for i := range(array) {
        if arrayMap[array[i]] == 1 {
            arrayMap[array[i]]++
        } else {
            arrayMap[array[i]] = 1
        }
    }
    // Count the number of occuring two elements sum
    sumsCounter := 0
    for j := minSum; j < maxSum+1; j++ {
        t := int64(j)
        for x := range(array) {
            if (arrayMap[t-array[x]] == 1 && (t-array[x]) != array[x]) || arrayMap[t-array[x]] > 1 {
                //fmt.Println(array[x], t)
                sumsCounter++
                break
            }
        }
    }
    
    return sumsCounter
}


func main() {
    // go routine controller
    //var wg sync.WaitGroup
    //wg.Add(2)
    // Assignment Case
    filePath := "2sum.txt"
    filePointer, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    buffer := bufio.NewReader(filePointer)
    assignmentArray := make([]int64, 0)
    line, err := buffer.ReadString('\n')
    for err == nil {
        number := strings.TrimSpace(line)
        value, er := strconv.ParseInt(number, 10, 64)
        if er != nil {
            panic(er)
        }
        assignmentArray = append(assignmentArray, value)
        line, err = buffer.ReadString('\n')
    }
    filePointer.Close()
    // Compute the 2-sums number
    fmt.Println(assignmentArray[0])
    //fmt.Println(len(assignmentArray))
    fmt.Println(TwoSums(assignmentArray, -10000, 10000))
    /*/ Concurrent execution
    var negSums, posSums int
    go func(ns *int) {
        defer wg.Done()
        *ns = TwoSums(assignmentArray, -10000, -1)
    } (&negSums)
    go func(ps *int) {
        defer wg.Done()
        *ps = TwoSums(assignmentArray, 0, 10000)
    } (&posSums)
    go func() {
        wg.Wait()
        fmt.Println(negSums+posSums)
    } ()
    time.Sleep(time.Second)
    // Test Case
    /*testArray := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
    sums := 0
    for i := -10; i < 11; i++ {
        go func(sumsp *int, j int) {
            defer wg.Done()
            s := TwoSums(testArray, j, j)
            *sumsp += s
        } (&sums, i)
    }
    go func() {
        wg.Wait()        
        fmt.Println(sums)
    } ()*/
    //wg.Wait()
}
