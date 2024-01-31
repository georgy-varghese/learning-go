package main

import (
  "fmt"
)

func main()  {

  closure := incrementer()
  fmt.Println(closure()) //1
  fmt.Println(closure()) //2
  fmt.Println(closure()) //3
  fmt.Println(closure()) //4
  
}

//closure func returns a function
func incrementer() func() int {
    x := 0
  return func() int {
    x++
    return x
  }
}
