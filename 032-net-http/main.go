package main

import (
	"fmt"
	"net/http"
)

type pumpkin int

func (p pumpkin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "This is pumpkin")
}

func main()  {
  var d pumpkin
  http.ListenAndServe(":8080", d)
}
