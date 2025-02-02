package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Eason",20})

	fmt.Println(person{name:"alice",age:30})

	fmt.Println(person{name:"fred"})

	fmt.Println(person{age:11})

	fmt.Println(&person{name:"ann",age:40})

	fmt.Println(newPerson("jon"))

	s := person{name:"sean",age:50}
	fmt.Println(s.name)

	sp:= &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct{
		name string
		isGood bool
	}{
		"kk",
		true,
	}
	fmt.Println(dog)
}