//https://www.hackerrank.com/challenges/compare-the-triplets?h_r=next-challenge&h_v=zen

package main
import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 res1, _ := reader.ReadString('\n')
 res2, _ := reader.ReadString('\n')

 fmt.Println(compare(res1, res2))
}

func compare(res1 string, res2 string) (int, int){
  var res_alice int
  var res_bob int

  alice := strings.Split(res1," ")
  bob := strings.Split(res2, " ")



  for i := 0; i < len(alice); i++ {
    alice_val,_ := strconv.Atoi(strings.Trim(alice[i], "\n"))
    bob_val,_ := strconv.Atoi(strings.Trim(bob[i], "\n"))

    if alice_val > bob_val {
      res_alice++
    }
    if bob_val > alice_val {
      res_bob++
    }
  }

  return res_alice, res_bob
}
