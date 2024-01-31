package main

import (
  "fmt"
  "runtime"
  "sync"
)

var wg sync.WaitGroup


func foo() {
  fmt.Printf("--- Single ---\n")
  for i:=0;i<10;i++{
    fmt.Println(i)
  }
}

func bar() {
  fmt.Printf("--- GoRoutine ---\n")
  for i:=0;i<10;i++{
    fmt.Println(i)
  }
  wg.Done()
}

func main()  {
  fmt.Printf("GOARCH:\t\t %s\n", runtime.GOARCH)
  fmt.Printf("OS:\t\t %s\n", runtime.GOOS)
  fmt.Printf("GORoutines:\t %d\n", runtime.NumGoroutine())
  fmt.Printf("CPUs:\t\t %d\n", runtime.NumCPU())


  wg.Add(1)
  foo()
  go bar()

  fmt.Printf("GORoutines:\t %d\n", runtime.NumGoroutine())
  wg.Wait()
}
