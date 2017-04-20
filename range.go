package main

import "fmt"

func main() {
  nums := []int{2,3,4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("Sum :", sum)

  kvc := map[string]string{"a":"apple", "b":"banana"}

  for k, v := range kvc {
    fmt.Printf("%s -> %s\n", k, v)
  }
}