package main

import "fmt"

type Person struct {
	name string
	sex  byte
	addr string
}

type Student struct {
	Person
	id  int
	age int
}

func main() {
	s1 := Student{Person{name: "mike", sex: 'm', addr: "bj"}, 1, 18}
	fmt.Println(s1.name, s1.sex)
}
