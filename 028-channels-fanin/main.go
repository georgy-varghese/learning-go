package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()  {

  even := make(chan int)
  odd := make(chan int)
  fanin := make(chan int)

  //send
  go send(even, odd)

  //receive
  go receive(even, odd, fanin)

  for v := range fanin {
    fmt.Println(v)
  }

  fmt.Println("About to exit")
  
}

func send(e, o chan<- int)  {
  for i:=1; i<=10; i++ {
    if i % 2 == 0 {
      e <- i
    } else {
      o <- i
    }
  }
  close(e)
  close(o)
}

func receive(e, o <-chan int, f chan<- int)  {
  //Adds two tasks to the waitgroup, as there are two goroutines
  wg.Add(2)

  go func ()  {
    for i := range e {
      f <- i
    }
    //Decrement waitgroup after completing the task
    defer wg.Done()
  } ()

  go func ()  {
    for j := range o {
      f <- j
    }
    //Decrement waitgroup after completing the task
    defer wg.Done()
  } ()

  //Wait until the waitgroup counter becomes zero
  wg.Wait()
  close(f)
}
