package main

import "fmt"

func main(){

	nums := []int {2,3,4}
	sum := 0
	for n,sum := range nums{
		sum += num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums{
		if num == 3{
			fmt.Println("index:",i)
		}
	}

	


}