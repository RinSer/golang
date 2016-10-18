package main

/**
 * Counts the number of inversions
 * in a given integer array.
 *
 * created by RinSer
**/

import (
  "os"
  "bufio"
  "fmt"
  "strconv"
)

var inversions int64

func merge_sort(array []int) []int {
  if len(array) == 1 {
    return array
  } else {
    l := len(array)/2
    first := array[:l]
    second := array[l:]
    first = merge_sort(first)
    second = merge_sort(second)
    sarray := make([]int, 0)
    j := 0
    for i := 0; i < len(first); i++ {
      if j < len(second) {
        if first[i] < second[j] {
          sarray = append(sarray, first[i])
        } else {
          sarray = append(sarray, second[j])
          inversions += int64(len(first)-i)
          j++
          i--
        }
      } else {
        sarray = append(sarray, first[i])
      }
    }
    for k := j; k < len(second); k++ {
      sarray = append(sarray, second[k])
    }
    return sarray
  }
}

func main() {
  file, err := os.Open("IntegerArray.txt")
  if err != nil {
    panic(err)
  }

  buffer := bufio.NewReader(file)

  integers := make([]int, 0)

  line, err := buffer.ReadString('\n')
  for err == nil {
    integer, er := strconv.Atoi(line[:len(line)-2])
    if er != nil {
      panic(er)
    }
    integers = append(integers, integer)
    line, err = buffer.ReadString('\n')
  }

  //test1 := []int{6, 5, 4, 3, 2, 1}
  //test2 := []int{1, 4, 2, 5, 3, 6}
  //test3 := []int{1, 2, 3, 4, 5, 6}

  sorted := merge_sort(integers)
  fmt.Println(len(sorted))
  fmt.Println(inversions)
}







