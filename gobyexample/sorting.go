package main

import (
	"fmt"
	"slices"
)

func main(){

	strs := []string{"c","a","b"}
	slices.Sort(strs)
	fmt.Println("Strings:",strs)

	ints := []int{3,5,2}
	slices.Sort(ints)
	fmt.Println("Ints:",ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ",s)

}