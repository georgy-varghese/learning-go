package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {

  //Chose log package generally for logging errors
  //log.Println
  //log.Fatalln
  //log.Panicln

  //Create a log file
  logfile, err := os.Create("log.txt")
  //Use the above file as log destination
  log.SetOutput(logfile)

  //a sample defer func
  defer foo()

  f, err := os.Open("no-file.go") // For read access.
  if err != nil {
    //Example of Fatal error
    //func exits, defer is not called
    //log.Fatal(err)
    //Example of log.Println
    //If log output (SetOuput) is set, logs are written to file
    //log.Println("Error: ", err)
    //Example of log.Panicln, defered func gets called
    log.Panicln("Panic", err)
    //Example of panic function, defered func gets called
    // panic(err)
  }
  defer f.Close()
}

func foo()  {
  fmt.Println("Deferred func called")
}
