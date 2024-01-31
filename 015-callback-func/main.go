package main

import (
  "fmt"
)

func main()  {

  fmt.Printf("add: %T\n", add)
  fmt.Printf("sub: %T\n", sub)
  fmt.Printf("callbk: %T\n", callbk)

  adding := callbk(27, 14, add)
  subtract := callbk(56, 17, sub)

  fmt.Println(adding)
  fmt.Println(subtract)
  
}

// add two integers
func add(a, b int) int {
  return a + b
}

// sub two integers
func sub(a, b int) int {
  return a - b
}

// Callbk func takes another function in its params
func callbk(a, b int, f func(int, int) int) int {
  return f(a, b)
}
