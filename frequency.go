package main

import "fmt"
import "reflect"
import "sort"

func analyseFrequency(nums []int) (map[int]int, []int) {
  // make the frequency map
  frequency := make(map[int]int)
  // for each number
  for _,num := range nums {
    _,exists := frequency[num]
    // if we havent come across it before
    if !exists {
      // lets put it into the map
      frequency[num] = 0
    }
    // increase the number of occurences
    frequency[num]++
  }

  // the uniques are just the keys of the frequency map, so...
  uniques := make([]int, 0, len(frequency))
  for key := range frequency {
    uniques = append(uniques, key)
  }

  return frequency, uniques
}

func main() {
  nums := []int{10, 5, 2, 10, 3, 5, 7, 1, 2, 4, 10, 9, 9, 1, 4, 6, 1, 7, 10, 9, 6, 5, 1, 4, 10, 1, 7, 7, 5, 6, 7, 6, 7, 9, 2, 6, 4, 5, 1, 8, 8, 3, 3, 7, 2, 4, 1, 5, 9, 1, 1, 5, 9, 3, 6, 7, 1, 4, 1, 3, 2, 2, 7, 4, 7, 7, 7, 7, 10, 2, 2, 10, 6, 6, 5, 9, 1, 10, 3, 4, 8, 8, 7, 9, 8, 3, 10, 4, 8, 9, 1, 9, 6, 9, 2, 6, 1, 8, 2, 7}

  // get our calculated frequency and uniques
  frequency, uniques := analyseFrequency(nums)

  // establish the expected frequencies
  expectedFrequency := map[int]int{1:14, 4:9, 9:11, 5:8, 2:10, 3:7, 7:15, 6:10, 8:7, 10:9}

  // compare to see if we calculated correctly
  fmt.Println("Expected:\t", expectedFrequency)
  fmt.Println("Actual:\t\t", frequency)
  // we have to use DeepEqual to determine if all keys and values are identical
  fmt.Println("Equal:\t\t", reflect.DeepEqual(expectedFrequency, frequency))

  // establish the expected uniques
  expectedUniques := []int{10, 5, 2, 3, 7, 1, 4, 9, 6, 8}

  // easiest way to tell if arrays of ints are equal is by sorting them first
  sort.Ints(expectedUniques)
  sort.Ints(uniques)
  // compare to see if we calculated correctly
  fmt.Println("Expected:\t", expectedUniques)
  fmt.Println("Actual:\t\t", uniques)
  fmt.Println("Equal:\t\t", reflect.DeepEqual(expectedUniques, uniques))

}
