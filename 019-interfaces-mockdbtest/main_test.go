package main

import "testing"

func TestGetUser(t *testing.T) {
  mockdb := MockDatastore{
    Users: map[int]User{
      2: {Id: 1, First: "Van Dyk"},
      11: {Id: 11, First: "Salah"},
    },
  }
  player, err := mockdb.GetUser(2)

  if err != nil {
    t.Error(err)
  }

  if player.First != "Van Dyk" {
    t.Errorf("Got: %s; Wanted: %s", player.First, "Van Dyk")
  }
  
}
