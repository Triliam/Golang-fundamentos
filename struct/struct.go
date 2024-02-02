package main

import "fmt"

type adresses struct {
	address string
	email string
	personalIdCPF int64
}

type user struct {
	name string
	age int8
	address adresses
}

func main() {

	var address1 adresses
	address1.address = "brooklyn 99"
	address1.email = "julius@email.com"
	address1.personalIdCPF = 33366699933

	var user1 user
	user1.name = "Julius"
	user1.age = 40
	user1.address = address1

	fmt.Println(user1)

	address2 := adresses{"Brooklyn", "chris@gmail.com", 22211122255}
	user2 := user{"Chris", 13, address2}

	fmt.Println(user2)
	fmt.Println(user2.address)
	fmt.Println(user2.address.address)
	fmt.Println(user2.address.email)
	//fmt.Println(user2.email) da erro - ver em pseudoHerança
	user3 := user{name: "Tais"}
	fmt.Println(user3)
}