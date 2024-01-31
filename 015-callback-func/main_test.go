package main

import "testing"

func TestAdd(t *testing.T)  {
  sum := add(19, 53)
  if sum != 72 {
    t.Errorf("Failed:\n Expected: %d, Got: %d", 72, sum)
  }
}

func TestSub(t *testing.T)  {
  minus := sub(53, 19)
  if minus != 34 {
    t.Errorf("Failed:\n Expected: %d, Got: %d", 34, minus)
  }
}

func TestCallbk(t *testing.T)  {
  domath := callbk(53, 20, add)
  if domath != 72 {
    t.Errorf("Failed:\n Expected: %d, Got: %d", 34, domath)
  }
}
