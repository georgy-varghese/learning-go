package main

import "fmt"

func main()  {
  cody := Footballer{
    name: "Gakpo",
    club: "Liverpool",
    country: "Netherlands",
  }
  diego := Striker{
    Footballer: Footballer{
      name: "Jota",
      club: "Liverpool",
      country: "Portugal",
    },
    lfoot: false,
  }

  //Without Interface
  cody.Speak()
  diego.Speak()

  //Using Interface
  //This is also example of polymorphesim (multiple type) as cody is Footballer as well as Player type
  saySomething(cody)
  saySomething(diego)
}

// Structs
type Footballer struct {
  name string
  club string
  country string
}

type Striker struct {
  Footballer
  lfoot bool
}

// Receiver methods
func (c Footballer) Speak() {
  fmt.Printf("%s says I love %s\n", c.name, c.club)
}

func (c Striker) Speak() {
  fmt.Printf("%s says I love to attack\n", c.name)
}

// Interface
type Player interface {
  Speak()
}

func saySomething(p Player) {
  p.Speak()
}
