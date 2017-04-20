package main

import "fmt"

func main() {
	/*
	y := [7]float64{5, 6, 7, 8, 9, 10 , 13}
	x := make(y	, 5)
	fmt.Println(x, y)
	*/
	s := make([]string,1)
	//s := []string{"a", "b", "c", "d"}

	//s := make([]string, 3)
	/*fmt.Println("emp: ", s)

	s[0] = "s"
	s[1] = "t"*/

  s = append(s, "b")
  s = append(s, "v")
  s = append(s, "z")

	fmt.Println("emp: ", s)
}