package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
  sum := Add(14, 7)
  if sum != 20 {
    t.Errorf("Failed: Expected: %d, Got: %d", 20, sum)
  }
}
