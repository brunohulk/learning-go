package main

import "fmt"

func plus(a int, b int) int {
  return a + b
}

func vals() (int, int, int ) {
  return 3, 7, 1
}

func main() {
  res := plus(1,2)
  res2, res1, res3 := vals()

  fmt.Println(res2, res1, res3)
  fmt.Println("1+2 =", res)
}