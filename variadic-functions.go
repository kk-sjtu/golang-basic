package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums," ")
	total := 0 
	// range returns both the index and the value of each element

	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main(){

	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)
}
