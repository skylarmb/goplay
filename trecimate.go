package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  trecimate(53)
  // expected output of trecimate(33)
  // Have 33, adding 0
  // Have 11, adding 1
  // Have 4, adding -1
  // Have 1, stopping

}

func trecimate(n int) {
  num := n
  for num > 0 {
    fmt.Println(num % 3)
    if num % 3  == 0 {
      fmt.Printf("Have %v, adding 0\n", num)
      num = num / 3
    } else {
      rand.Seed(time.Now().UTC().UnixNano())
      rand := rand.Intn(2) // 0 or 1
      if rand == 0 {
        rand = -1 // make 0s into -1s
      }
      fmt.Printf("Have %v, adding %v\n", num, rand)
      num = num + rand
    }
  }
  fmt.Printf("Have %v, stopping\n", num)
}
