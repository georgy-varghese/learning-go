package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main()  {
  ctx, cancel := context.WithCancel(context.Background())

  fmt.Println("Error check 1: ", ctx.Err())
  fmt.Println("No of GoRoutines: ", runtime.NumGoroutine())

  go func ()  {
    n := 0
    for {
      select {
      case <- ctx.Done():
        return
      default:
        n++
        time.Sleep(time.Millisecond * 200)
        fmt.Println(n)
      }
    }
  } ()

  time.Sleep(time.Second * 2)
  fmt.Println("Error check 2: ", ctx.Err())
  fmt.Println("No of GoRoutines: ", runtime.NumGoroutine())

  fmt.Println("About to Cancel....")
  cancel()
  fmt.Println("Error check 3: ", ctx.Err())
  fmt.Println("No of GoRoutines: ", runtime.NumGoroutine())

  fmt.Println("Starting next goroutine/context....")
  gen := func (ctx context.Context) <-chan int {
    m := 100 
    dst := make(chan int)
    go func ()  {
      for {
        select {
        case <- ctx.Done():
          fmt.Println("Error check 4: ", ctx.Err())
          fmt.Println("No of GoRoutines: ", runtime.NumGoroutine())
        case dst <- m:
          m++
        }
      }
    } ()
    return dst
  }

  ctx2, cancel := context.WithCancel(context.Background())
  defer cancel()

  for i := range gen(ctx2) {
    fmt.Println(i)
    if i == 105 {
      break
    }
  }
  
}
