//https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/160-hands-on-67-interface-mock-db-test/main.go

package main

import (
	"fmt"
	"log"
)

// User represents a User with id and first name
type User struct {
  Id int
  First string
}

// MockDatastore is a temporary datastore
type MockDatastore struct {
  Users map[int]User
}

// GetUser returns a User from MockDatastore
func (md *MockDatastore) GetUser(id int) (User, error) {
  u, ok := md.Users[id]
  if !ok {
    return User{}, fmt.Errorf("User %v not found", id)
  }
  return u, nil
}

// SaveUser saves test user to MockDatastore
func (md *MockDatastore) SaveUser(u User) error {
  _, ok := md.Users[u.Id]
  if ok {
    return fmt.Errorf("Error: User already exists - %v", ok)
  }
  md.Users[u.Id] = u
  return nil
}

type Datastore interface {
  GetUser(id int) (User, error)
  SaveUser(u User) error
}

type Service struct {
  ds Datastore
}

func (s *Service) GetUser(id int) (User, error) {
  return s.ds.GetUser(id)
}

func (s *Service) SaveUser(u User) error {
  return s.ds.SaveUser(u)
}
  
func main()  {
  mockdb := MockDatastore{
    Users: make(map[int]User),
  }

  svc := Service{
    ds: &mockdb,
  }

  user1 := User{
    Id: 1,
    First: "Alison",
  }

  error := svc.SaveUser(user1)

  if error != nil {
    log.Fatalf("Error: %s", error)
  }

  log.Println(user1)

  user, error := svc.GetUser(user1.Id)

  if error != nil {
    log.Fatalf("Error: %s", error)
  }

  fmt.Println(user)
  
}
