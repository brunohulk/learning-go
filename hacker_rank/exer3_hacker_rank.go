//https://www.hackerrank.com/challenges/a-very-big-sum

package main
import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "input_string/clear_string"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 var total int64

 input_numbers, _ := reader.ReadString('\n')
 str_numbers, _ := reader.ReadString('\n')

 total_n,_ := strconv.Atoi(clear_input(input_numbers))

 numbers := strings.Split(str_numbers," ")

 for i, value := range numbers {
   val, _ := strconv.ParseInt(clear_input(value), 10, 64)
   total += val

   if i == total_n {
     break
   }
 }

 fmt.Println(total)
}

func clear_input(text string) string {
  return strings.Trim(text, "\n")
}




















