package main

import "fmt"

func main() {
  fmt.Println(half(1))
}

func half(val int) (int, bool) {
  if val % 2 == 1 {
    return 1, true
  }
  return 0, false
}
