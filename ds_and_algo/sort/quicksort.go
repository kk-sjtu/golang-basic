package main

import "fmt"

func main() {
	array := []int{2, 4, 1, 0, 3, 5}

	quickSort(array, 0, len(array)-1)
	//fmt.Println(array)
}

func quickSort(array []int, left int, right int) {
	l, r := left, right
	if left >= right {
		return
	}
	fmt.Println(array)
	for left < right {
		for left < right && array[right] >= array[l] {
			right--
		}
		for left < right && array[left] <= array[l] {
			left++
		}
		if left < right {
			array[left], array[right] = array[right], array[left]
		}

	}
	temp := array[l]

	array[l] = array[left]
	array[left] = temp
	fmt.Println(array)
	quickSort(array, l, left-1)
	quickSort(array, left+1, r)

}
