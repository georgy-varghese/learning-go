package main

import "fmt"

func main()  {
  ci := make(chan int) //bidirectional
  cs := make(chan<- int) //send channel
  cr := make(<-chan int) //receive channel

  go func () {
    ci <- 43
  }()

  fmt.Printf("%v\n", <-ci)

  fmt.Println("----")
  fmt.Printf("ci\t%T\n", ci)
  fmt.Printf("cs\t%T\n", cs)
  fmt.Printf("cr\t%T\n", cr)
  
}

// declare a channel
// ci := make(chan int) //bidirectional
// ci := make(chan<- int) //send channel
// ci := make(<-chan int) //receive channel
// putting values on a channel
// ci <- 43
// taking values off channel
// <-ci
// set buffered channel (buffer size 1)
// ci := make(chan int, 1)
