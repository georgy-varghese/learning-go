package main

import (
  "fmt"
)

func main()  {

  even := make(chan int)
  odd := make(chan int)
  quit := make(chan int)

  //send
  go send(even, odd, quit)

  //select
  for {
    select {
      case v := <- even:
        fmt.Println("Received from even", v)
      case v := <- odd:
        fmt.Println("Received from odd", v)
      case v := <- quit:
        fmt.Println("Quitting ", v)
        return
    }
  }
  
}

func send(e, o, q chan<- int) {
  for i:=1; i<=10; i++ {
    if i % 2 == 0 {
      e <- i
    } else {
      o <- i
    }
  }
  q <- 0
  close(e)
  close(o)
}
