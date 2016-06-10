/**************
Find the sum of the even Fibonacci numbers < 4*10^6
***************/

package main
import "fmt"
func main() {
  total := 0
  prevTerm :=0
  thisTerm := 1
  for ; thisTerm <= 4000000; {
    if thisTerm %2 == 0 {
      total += thisTerm
      // fmt.Printf("__ thisTerm: %v, total: %v \n", thisTerm, total)
    }
    temp := prevTerm
    prevTerm = thisTerm
    thisTerm = temp + prevTerm
  }
  fmt.Printf("Total sum of even Fibonacci numbers under 4G is %v\n", total)
}
