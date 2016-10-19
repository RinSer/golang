package main

/**
 * My implementation of the QuickSort algorithm 
 * with different pivot choice.
 *
 * created by RinSer
 */

import (
  "os"
  "bufio"
  "fmt"
  "strconv"
)

var count_first int64
var count_last int64
var count_median int64

func swap(array []int, i int, j int) {
  temp := array[i];
  array[i] = array[j]
  array[j] = temp
}

func partition(array []int, p int, r int) int {
  q := p+1
  for j := p+1; j < r+1; j++ {
    if (array[j] < array[p]) {
      swap(array, q, j)
      q++
    }
  }
  swap(array, p, q-1)
  return q-1
}  

func quickSort(array []int, p int, r int, pivot func(array []int, p int, r int)) {
  if (p < r) {
    pivot(array, p, r)
    q := partition(array, p, r)
    quickSort(array, p, q-1, pivot)
    quickSort(array, q+1, r, pivot)
  }
}

func main() {
  file, err := os.Open("QuickSort.txt")
  if err != nil {
    panic(err)
  }

  buffer := bufio.NewReader(file)

  integers1 := make([]int, 0)
  integers2 := make([]int, 0)
  integers3 := make([]int, 0)

  line, err := buffer.ReadString('\n')
  for err == nil {
    integer, er := strconv.Atoi(line[:len(line)-2])
    if er != nil {
      panic(er)
    }
    integers1 = append(integers1, integer)
    integers2 = append(integers2, integer)
    integers3 = append(integers3, integer)
    line, err = buffer.ReadString('\n')
  }

  // Pivot choice
  First := func(array []int, p int, r int) {
    count_first += int64(r-p)
  }

  Last := func(array []int, p int, r int) {
    count_last += int64(r-p)
    swap(array, p, r)
  }

  Median := func(array []int, p int, r int) {
    count_median += int64(r-p)
    empty := func(array []int, p int, r int) {}
    first := array[p]
    middle := array[p+(r-p)/2]
    last := array[r]
    triple := []int{first, middle, last}
    quickSort(triple, 0, 2, empty)
    pivot := triple[1]
    switch pivot {
      case middle:
        swap(array, p, p+(r-p)/2)
      case last:
        swap(array, p, r)
    }
  }
   
  quickSort(integers1, 0, len(integers1)-1, First)
  //fmt.Println(integers1)
  fmt.Println("First pivot count: ", count_first)
  quickSort(integers2, 0, len(integers2)-1, Last)
  //fmt.Println(integers2)
  fmt.Println("Last pivot count: ", count_last)
  quickSort(integers3, 0, len(integers3)-1, Median)
  //fmt.Println(integers3)
  fmt.Println("Median pivot count: ", count_median)
}
