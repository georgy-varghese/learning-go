package main

import "fmt"

func main()  {
  c := make(chan int)

  //send
  go func ()  {
    for i:=1; i<=5; i++ {
      c <- i
    }
    close(c)
  } ()

  //receive
  for i := range c {
    fmt.Println(i)
  }
  
  fmt.Println("about to exit")
  
}
