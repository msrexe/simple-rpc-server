package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var reply Person

	client, err := rpc.DialHTTP("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// Add person
	client.Call("API.AddPerson", Person{Name: "Melih", Age: 20}, &reply)
	fmt.Println("Added : ", reply)
	reply = Person{}

	client.Call("API.AddPerson", Person{Name: "Burak", Age: 19}, &reply)
	fmt.Println("Added : ", reply)
	reply = Person{}

	client.Call("API.AddPerson", Person{Name: "Ahmet", Age: 30}, &reply)
	fmt.Println("Added : ", reply)
	reply = Person{}

	// Get Person
	client.Call("API.GetPerson", "Melih", &reply)
	fmt.Println(reply)
	reply = Person{}

	client.Call("API.DeletePerson", "Burak", nil)
	client.Call("API.GetPerson", "Burak", &reply)
	fmt.Println("Deleted : ", reply)
	reply = Person{}

	client.Call("API.EditPersonAge", Person{Name: "Melih", Age: 21}, &reply)
	client.Call("API.GetPerson", "Melih", &reply)
	fmt.Println("Edited : ", reply)
}
