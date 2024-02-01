package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
  c1 := make(chan int)
  c2 := make(chan int)

  //populate
  go populate(c1)

  //fanout and in
  go fanoutin(c1, c2)

  for v := range c2 {
    fmt.Println(v)
  }

  fmt.Println("About to exit")
  
}

func populate(c chan<- int)  {
  for i:=0; i<=10; i++ {
    c <- i
  }
  close(c)
}

// Non throttled edition
// func fanoutin(c1, c2 chan int)  {
//   var wg sync.WaitGroup
//   for i := range c1 {
//     wg.Add(1)
//     //Fanout into go routines
//     go func (j int)  {
//       //Fanin the output into a single channel
//       c2 <- timeconsumingwork(j)
//       wg.Done()
//     } (i)
//   }
//   wg.Wait()
//   close(c2)
// }

// Throttled edition
func fanoutin(c1, c2 chan int)  {
  var wg sync.WaitGroup
  const goroutines = 10
  wg.Add(goroutines)

  for i:=1; i<=goroutines; i++{
    //Fanout into go routines
    go func ()  {
      for i := range c1 {
        func (j int)  {
          //Fanin the output into a single channel
          c2 <- timeconsumingwork(j)
        }(i)
      }
      wg.Done()
    }()
  }
  wg.Wait()
  close(c2)
}
func timeconsumingwork(x int) int {
  time.Sleep(time.Second * 2)
  return x
}
