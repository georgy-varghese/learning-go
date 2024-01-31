package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
  first string
}

func (p person) writeOut(w io.Writer) {
  w.Write([]byte(p.first))
}

func main() {

  geo := person{
    first: "Georgy",
  }

  // Write to file
  f, err := os.Create("output.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  geo.writeOut(f)

  // Write to buffer
  var s bytes.Buffer 
  geo.writeOut(&s)
  fmt.Println(s.String())

}
