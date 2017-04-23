package main

import "fmt"

func main() {
  fmt.Println(max_value(1,4,2,7,6,4,9))
}

func max_value(args ...int) int {
  max := 0

  for  _, value := range args {
    if (value > max) {
      max = value
    }
  }
  return max
}
