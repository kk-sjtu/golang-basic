package main

import (
	"fmt"
	"cmp"
	"slices"
)

func main(){

	fruits := []string{"abc","abcd","abcdef"}
	lenCmp := func(a,b string)int{
		return cmp.Compare(len(a),len(b))
	}

	slices.SortFunc(fruits,lenCmp)
	fmt.Println(fruits)

	type Person struct{
		name string
		age int
	}

	people := []Person{
		Person{"a",30},
		Person{"b",20},
		Person{"c",50},
	}
	slices.SortFunc(people,
		func(a,b Person)int{
			return cmp.Compare(a.age,b.age)
		})
	fmt.Println(people)
}
