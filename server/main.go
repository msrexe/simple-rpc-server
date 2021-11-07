package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Person struct {
	Name string
	Age  int
}

type API struct {
}

var database []Person

func (api *API) GetPerson(name string, reply *Person) error {
	for _, person := range database {
		if person.Name == name {
			*reply = person
			return nil
		}
	}
	return fmt.Errorf("person not found")
}

func (api *API) AddPerson(person Person, reply *Person) error {
	database = append(database, person)
	*reply = person
	return nil
}

func (api *API) EditPersonAge(person Person, reply *Person) error {
	for index, p := range database {
		if p.Name == person.Name {
			database[index].Age = person.Age
			*reply = database[index]
			return nil
		}
	}
	return fmt.Errorf("update failed")
}

func (api *API) DeletePerson(name string, reply *Person) error {
	for index, p := range database {
		if p.Name == name {
			database = append(database[:index], database[index+1:]...)
			*reply = p
			return nil
		}
	}
	return fmt.Errorf("delete failed")
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatalf("api register failed, err : %s", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listener error, %s", err)
	}

	log.Printf("serving rpc on %d", 8080)
	log.Fatal(http.Serve(listener, nil))
}
