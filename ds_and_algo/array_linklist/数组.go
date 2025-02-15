// 1. 初始化

var arr [5]int

nums := [] int{1,3,2,5,4}

func randomAccess(nums[]int)(randomNum int){
	randomIndex := rand.Intn(len(nums))
	randomNum = nums[randomIndex]
	return randomNum
}

func insert(nums[]int,num int,index int){
	for i:=len(nums)-1;i>index;i--{
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

func remove(nums[]int,index int){
	for i:= index;i<len(nums)-1;i++{
		nums[i]=nums[i+1]
	}
}

func traverse(nums[] int){
	count := 0
	for i:=0;i<len(nums);i++{
		count += nums[i]
	}
	count = 0
	for _,num := range nums{
		count += nim
	}

	for i,num := range nums{
		count += nums[i]
		count += num
	}
}

func find(nums []int,target int)(index int){
	index = -1
	for i:= 0;i<len(nums);i++{
		if nums[i] == target{
			index = i
			break
		}
	}
}

func extend(nums[] int,enlarge int)[]int{
	res := make([]int,len(nums)+enlarge)
	for i,num: range nums{
		res[i] = num
	}
	return res
}