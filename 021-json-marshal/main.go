package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type footballer struct {
	First  string
	Last   string
	Jersey int
}

func main() {

	lvl1 := footballer{
		First:  "Diogo",
		Last:   "Jota",
		Jersey: 11,
	}
	lvl2 := footballer{
		First:  "Mc",
		Last:   "Alistair",
		Jersey: 10,
	}
	liverpool := []footballer{
		lvl1,
		lvl2,
	}

	b, err := json.Marshal(liverpool)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(liverpool)
	os.Stdout.Write(b)

}
