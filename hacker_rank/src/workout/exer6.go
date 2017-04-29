package main

import (
  "bufio"
  "fmt"
  "time"
  "os"
  inputs "workout/input_string"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  date, _ := reader.ReadString('\n')

  time,_ := time.Parse("03:04:05PM", inputs.Clear_input(date))
  fmt.Println(time.Format("15:04:05"))
}
