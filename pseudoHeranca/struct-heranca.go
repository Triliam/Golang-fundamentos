package main

import "fmt"

type person struct {
	name string
	age int8
}

// Go n é orientado a objetos, o mais proximo de objetos e herança são as structs. É dessa forma que traz os dados(atributos) contidos na struct person para a struct student
type student struct {
	person
	course string
	ra int64
}

func main() {

	person1 := person{"Tais", 39}
	student1 := student{person1, "Programacao", 3334445}
	fmt.Println(student1)
	fmt.Println(student1.name) // ao inves de student1.person.name
}