package main

import "fmt"

func main() {
  var wait string

  go oi()
  go hello()
  fmt.Println("Não estou preso")
  fmt.Scanln(&wait)
}

func oi() {
  for i := 0; i < 300; i++ {
   fmt.Println("Oi")
  }
}

func hello() {
  for i:= 0; i < 300; i++ {
    fmt.Println("Hello")
  }
}
