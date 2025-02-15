package main

import "fmt"

type myList struct {
	data []int
	len  int
	cap  int
}
type IntSlice []int

func ini(nums []int) {
	nums = append(nums, 1)
	nums = append(nums, 2)
}

func visit(nums []int, i int) (r int) {
	return nums[i]
}

func (nums *IntSlice) del(i int) {
	*nums = append((*nums)[:i], (*nums)[i+1:]...)
}

func (nums *IntSlice) travel() {
	for _, v := range *nums {
		fmt.Printf("%d ", v) // 怎么不换行？告诉我怎么不换行
		// fmt.Println(v)
	}
}

func (nums *IntSlice) merge(nums1 IntSlice) {
	*nums = append(*nums, nums1...)
}

func main() {
	nums := IntSlice{1, 2, 3, 4, 5}
	fmt.Println(nums)
	nums.del(1)
	fmt.Println(nums)
	nums.travel()
	fmt.Println(nums)
	nums.merge(IntSlice{6, 7, 8})
	fmt.Println(nums)
	nums.travel()
	fmt.Println(nums)
}
