package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  go test(0)
  var input string;

  for i:= 0; i < 10; i++ {
    go test(i)
  }

  fmt.Scanln(&input)
}

func test(n int) {
  for i :=0; i <  10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}
