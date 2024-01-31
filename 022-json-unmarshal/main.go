package main

import (
	"encoding/json"
	"fmt"
)

type footballer struct {
	First  string `json:"First"`
	Last   string `json:"Last"`
	Jersey int    `json:"Jersey"`
}

func main()  {
  s := `[{"First":"Diogo","Last":"Jota","Jersey":11},{"First":"Mc","Last":"Alistair","Jersey":10}]`
  sb := []byte(s)

  var liverpool []footballer

  err := json.Unmarshal(sb, &liverpool)

  if err != nil {
    fmt.Println("error: ", err)
  }

  fmt.Printf("All data\n%+v", liverpool)

  for i, v := range liverpool {
    fmt.Printf("\n--- Player %d ---\n", i)
    fmt.Printf("%d %s %s", v.Jersey, v.First, v.Last)
  }
  
}

