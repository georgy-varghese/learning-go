package main

import "fmt"

//AddI accepts only Int args
func addI(a, b int) int {
  return a + b
}

//AddF will accept only float64 args
func addF(a, b float64) float64 {
  return a + b
}

//Generic Function
//AddT function can take int or float64
//This method is called type constraint
//as we are constraining type to int and float64
func addT[T int | float64](a, b T) T {
  return a + b
}

//Another way of using generic function is using type set interface
type myNumbers interface {
  int | float64
}

func addTI[T myNumbers](a, b T) T {
  return a + b
}

func main()  {
  fmt.Println(addI(23, 9))
  fmt.Println(addF(23.5, 19.3))
  fmt.Println(addT(22, 23.5))
  fmt.Println(addTI(22, 23.5))
}
