package main

import "fmt"

func moveZeroes(nums []int) {
	leng := len(nums)
	counter := 0
	total_zero := 0
	for idx, val := range nums {
		if val == 0 {
			counter++
			total_zero++
		} else {
			//fmt.Println(val)
			fmt.Println(idx, counter)
			//fmt.Println(idx-counter)
			nums[idx-total_zero] = val
			//fmt.Println(nums[idx-counter])
			counter = 0
		}
	}
	for i := leng - 1; i > leng-total_zero-1; i-- {
		nums[i] = 0
	}
}
