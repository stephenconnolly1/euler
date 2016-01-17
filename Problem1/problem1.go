/***********************************
PROBLEM 1 - https://projecteuler.net/problem=1
Sum of the Multiples of 5 and 3
***********************************/
package main

import "fmt"

func main() {
   total := 0
   for i := 1; i < 1000; i++ {
      if i % 3 == 0 || i % 5 == 0 {
          total += i
      }
   }
   fmt.Printf("Total of multiples of 3 or 5 is : %v\n", total)
}