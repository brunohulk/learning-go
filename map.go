package main

import "fmt"

func main() {
  m := make(map[string]string)
  //var m map[string]string

  m["k1"] = "Teste"
  m["k2"] = "Teste"

  fmt.Println("Map ", m)
}