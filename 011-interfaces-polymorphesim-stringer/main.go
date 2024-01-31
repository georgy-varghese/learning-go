package main

import (
	"fmt"
	"log"
	"strconv"
)

func main()  {
  a := bat{
    name: "Eecha Bat",
  }

  var b count = 42

  fmt.Println(a)
  fmt.Println(b)

  log.Println(a)
  log.Println(b)

  logInfo(a)
  logInfo(b)
  
}

type bat struct {
  name string
}

type count int

// Below methods implicitly adds to Stringer interface
func (b bat) String() string {
  return fmt.Sprint("This is ", b.name)
}

func (c count) String() string {
  return fmt.Sprint("This is number ", strconv.Itoa(int(c)))
}

// Wrapper for logging
func logInfo(s fmt.Stringer) {
  log.Println("InterfaceLogs: ", s.String())
}
