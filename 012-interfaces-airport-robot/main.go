package main

import (
  "fmt"
)
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
  LanguageName() string
  Greet(name string) string
}

type Visitor struct {
  hello string
  language string
}

func (v Visitor) LanguageName() string {
  return(v.language)
}

func (v Visitor) Greet(name string) string {
  return(fmt.Sprintf("%s, %s", v.hello, name))
}

type Italian struct {}

func (i Italian) LanguageName() string {
  return("Italian")
}

func (i Italian) Greet(name string) string {
  return(fmt.Sprintf("Ciao, %s", name))
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
  return("Portuguese")
}

func (p Portuguese) Greet(name string) string {
  return(fmt.Sprintf("Olá, %s", name))
}

func SayHello(name string, g Greeter) string {
  message := fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
  return message
}

func main() {
    germanGreeter := Visitor{
      hello: "Hallo",
      language: "German",
    }
    fmt.Println(SayHello("Boca", germanGreeter))

    italianGreeter := Italian{}
    fmt.Println(SayHello("Piomonte", italianGreeter))

    portugeseGreeter := Portuguese{}
    fmt.Println(SayHello("Jota", portugeseGreeter))
}
