package main

import "fmt"

func main() {
  fmt.Println(sum(1,2,3,4,5))
}

func sum(args ...int) int {
  var total int
  total = 0
 
  for _, value := range args {
    total += value
  }

  return total
}
