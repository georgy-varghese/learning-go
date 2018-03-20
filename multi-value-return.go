package main

import "fmt"

func PlusMinus(a, b int) (plus, minus int) {
    plus = a + b
    minus = a - b
    return
}

func main() {
  x := 6
  y := 4
  xPLUSy, xMINUSy := PlusMinus(x,y)
  fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
  fmt.Printf("%d - %d = %d\n", x, y, xMINUSy)
}
