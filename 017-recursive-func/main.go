package main

import (
  "fmt"
)

func main() {
  fmt.Println(factorial(4))
  fmt.Println(loop(4))
}

//recursive function calls itself as a return, resulting in a mirrored loop
func factorial(n int) int {
  fmt.Println("This is", n)
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}


// anything that uses recursive function can be implemented using function loop
func loop(n int) int {
  x := 1
  for n > 0 {
    fmt.Println("n is", n)
    x *= n
    n--
  }
  return x
}
