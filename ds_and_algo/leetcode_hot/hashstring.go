package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) (ans [][]string) {
	hashmap := make(map[string]([]string))
	for _, val := range strs {
		hashmap[sortLetters(val)] = append(hashmap[sortLetters(val)], val)
	}

	for _, val := range hashmap {
		ans = append(ans, val)
	}
	return ans
}

func sortLetters(strs string) string {
	s0 := []rune(strs)
	sort.Slice(s0, func(i, j int) bool {
		return s0[i] < s0[j]
	})
	return string(s0)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
