package main

import "testing"

func TestIncrementer(t *testing.T)  {
  closure := incrementer()
  got := closure()
  want := 1

  if got != want {
    t.Errorf("Failed\n: Want: %d, Got: %d", want, got)
  }
  
}
