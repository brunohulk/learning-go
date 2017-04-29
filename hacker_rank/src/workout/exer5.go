//www.hackerrank.com/challenges/staircase

package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {

 var stair_int int

 reader := bufio.NewReader(os.Stdin)

 stair_length, _ := reader.ReadString('\n')

 int64_stair_length,_ := strconv.ParseInt(clear_input(stair_length), 10, 64)

 stair_int = int(int64_stair_length)
 for i := 1; i <= stair_int; i++  {
   str_line := strings.Repeat(" ", stair_int - i)   + strings.Repeat("#", i)
   fmt.Println(str_line)
 }
}

func clear_input(text string) string {
  return strings.Trim(text, "\n")
}
