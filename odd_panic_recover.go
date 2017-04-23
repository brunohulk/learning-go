package main

import "fmt"

func main() {
  nextOdd := makeOddGenerator();
  defer func() {
   str := recover()
   fmt.Println(str)
  }()

  panic("TESTESTE")
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
}

func makeOddGenerator() func() uint {
  number := uint(1)
  return func() (ret uint) {
    ret = number
    number += 2
    return
  }
}
