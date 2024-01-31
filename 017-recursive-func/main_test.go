package main

import "testing"

func TestFactorial(t *testing.T)  {
  got := factorial(3)
  want := 6

  if got != want {
    t.Errorf("Failed\n: Want: %d, Got: %d", want, got)
  }
  
}

func TestLoop(t *testing.T)  {
  got := loop(3)
  want := 6

  if got != want {
    t.Errorf("Failed\n: Want: %d, Got: %d", want, got)
  }
  
}
