package main

import (
  "fmt"
)

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }(argument(s))

func main() {
  foo()

  //anonymus func
  func ()  {
    fmt.Println("This is anonymus func without params")
  } ()

  //anonymus func with params
  func (s string)  {
    fmt.Println("My name is", s)
  } ("Georgy")

}

func foo() {
  fmt.Println("This is name func")
}
