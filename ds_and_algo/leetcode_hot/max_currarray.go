package main

import "fmt"

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	currentarr := 0

	hashSet := make(map[int]struct{})

	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	longestStreak := 0

	for num := range hashSet {
		if _, ok := hashSet[num-1]; !ok {
			currentNum := num
			currentarr = 1

			for {
				if _, ok := hashSet[currentNum+1]; ok {
					currentNum++
					currentarr++
				} else {
					break
				}

			}

			if currentarr > longestStreak {
				longestStreak = currentarr
			}

		}
	}
	return longestStreak

}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums1)) // 输出: 4

	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums2)) // 输出: 9

	nums3 := []int{1, 0, 1, 2}
	fmt.Println(longestConsecutive(nums3)) // 输出: 3

	nums4 := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums4)) // 输出: 3
}
