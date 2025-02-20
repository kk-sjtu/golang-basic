// 超过时间限制

//package main
//
//func maxArea(height []int) int {
//	Max :=0
//	for p,_ := range height{
//		for q:=p+1;q<len(height){
//			Max = max_(Max,(q-p) * min_(height[p],height[q]))
//
//		}
//	}
//
//
//	return Max
//}
//
//func min_(a,b int)int{
//	if(a<b) return a
//	return b
//}
//func max_(a,b int)int{
//	if(a>b) return a
//	return b
//}

func maxArea(height []int) int {
	Max := 0
	l, r := 0, len(height)-1
	width := r - l

	for l < r {
		Max = max_(Max, width*min_(height[l], height[r]))
	if height[l] < height[r] {
		l++
	} else {
		r--
	}
		width = r - l
	}
	return Max

}

func min_(a, b int) int {
	if a < b {
	return a
	}
	return b
}

func max_(a, b int) int {
	if a > b {
	return a
	}
	return b
}

