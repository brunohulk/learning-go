package main
import "fmt"

func main() {
 var input int
 var args []int
 var num int

 fmt.Println("Digite o número de elementos")
 fmt.Scanf("%d", &input)

 for x := 0; x <input; x++ {
   fmt.Println("Digite o número")
   fmt.Scanf("%d", &num)
   args = append(args, num)
 }

 fmt.Println(sum(args))
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}

func sum(args []int) int {
    var total int
    for _, x := range args {
        total += x
    }
   return total
}
