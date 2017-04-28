//https://www.hackerrank.com/challenges/plus-minus

package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "workout/input_string"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 var total_positive, total_negative, total_zero int64
 var fraction_negative, fraction_zero, fraction_positive float32

 input_numbers, _ := reader.ReadString('\n')
 str_numbers, _ := reader.ReadString('\n')

 total_n,_ := strconv.ParseInt(input_string.Clear_input(input_numbers), 10, 64)

 numbers := strings.Split(str_numbers," ")

 for _, value := range numbers {
   val, _ := strconv.ParseInt(input_string.Clear_input(value), 10, 64)
   if val == 0 {
     total_zero++
   }
   if val > 0 {
     total_positive++
   }
   if val < 0 {
     total_negative++
   }
 }

 fraction_negative = 0.0
 if total_negative != 0 {
   fraction_negative = float32(total_negative)/float32(total_n)
 }

 fraction_positive = 0.0
 if total_positive != 0 {
   fraction_positive = float32(total_positive)/float32(total_n)
 }

 fraction_zero = 0.0
 if total_zero != 0 {
   fraction_zero = float32(total_zero)/float32(total_n)
 }

 fmt.Printf("%f \n", fraction_positive)
 fmt.Printf("%f \n", fraction_negative)
 fmt.Printf("%f \n", fraction_zero)
}
